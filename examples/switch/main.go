package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bluenviron/gomavlib/v2"
	"github.com/spaolacci/murmur3"
)

const packetLifetime = 1 * time.Second // adjust as necessary
const AvgPacketsPerSecond = 124
const ConnectionInactiveDuration = 1 * time.Second
const channelAddrA = "127.0.0.1:14557"
const channelAddrB = "127.0.0.1:14558"
const downLinkAddr = "127.0.0.1:14560"

func main() {

	endpointA := gomavlib.EndpointUDPServer{Address: channelAddrA}
	endpointB := gomavlib.EndpointUDPServer{Address: channelAddrB}
	endpointDown := gomavlib.EndpointUDPClient{Address: downLinkAddr}

	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			endpointA,
			endpointB,
			endpointDown,
		},
		Dialect:     nil,
		OutVersion:  gomavlib.V2,
		OutSystemID: 10,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	var channelA *gomavlib.Channel
	var channelB *gomavlib.Channel
	var downChannel *gomavlib.Channel

	// wait till number of channels is 3
	for {
		if len(node.Channels) == 3 {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	// set channels
	for channel := range node.Channels {
		// address without protocol tag (assumed UDP)
		splitTmpAddr := strings.Split(channel.String(), ":")
		fmtTmpAddr := splitTmpAddr[1] + ":" + splitTmpAddr[2]

		if fmtTmpAddr != downLinkAddr {
			if channelA == nil {
				channelA = channel
			} else if channelB == nil {
				channelB = channel
			}
		} else {
			downChannel = channel
		}
	}

	fmt.Printf("channels set!")

	// Initialize the active channel
	activeChannel := channelA
	//lastActiveChannelMessageTime := time.Now()
	fmt.Printf("Active channel is %+v\n", activeChannel)

	sentFrames := make(map[uint64]time.Time)
	duplicates := 0
	notDuplicates := 0

	receivedPacketsA := 0
	receivedPacketsB := 0
	//packetsPerSecond := 0
	ticker := time.NewTicker(1 * time.Second)
	packetsMu := sync.Mutex{}

	// Create and open a csv file
	file, err := os.Create("results.csv")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	err = writer.Write([]string{"Time", "Channel A PPS", "Channel B PPS", "Channel A Loss (%)", "Channel B Loss (%)", "Filtered Loss (%)"})
	if err != nil {
		panic(err)
		return
	}

	go func() {
		for range ticker.C {
			packetsMu.Lock()
			tmpPacketsA := receivedPacketsA
			tmpPacketsB := receivedPacketsB
			tmpTotal := notDuplicates

			receivedPacketsA = 0
			receivedPacketsB = 0
			notDuplicates = 0
			packetsMu.Unlock()

			// Write to CSV file instead of printing
			err := writer.Write([]string{
				time.Now().Format(time.RFC3339),
				fmt.Sprintf("%v", tmpPacketsA),
				fmt.Sprintf("%v", tmpPacketsB),
				fmt.Sprintf("%v", (1-float32(tmpPacketsA)/AvgPacketsPerSecond)*100),
				fmt.Sprintf("%v", (1-float32(tmpPacketsB)/AvgPacketsPerSecond)*100),
				fmt.Sprintf("%v", (1-float32(tmpTotal)/AvgPacketsPerSecond)*100),
			})
			if err != nil {
				log.Fatalln(err)
				return
			}
			writer.Flush()
			fmt.Printf("Channel A Packets per second: %v / %v\n", tmpPacketsA, AvgPacketsPerSecond)
			fmt.Printf("Channel B Packets per second: %v / %v\n", tmpPacketsB, AvgPacketsPerSecond)
			fmt.Printf("Channel A loss = : %v percent\n", (1-float32(tmpPacketsA)/AvgPacketsPerSecond)*100)
			fmt.Printf("Channel B loss = : %v percent\n", (1-float32(tmpPacketsB)/AvgPacketsPerSecond)*100)
			fmt.Printf("Filtered loss  = %v percent\n\n", (1-float32(tmpTotal)/AvgPacketsPerSecond)*100)
		}
	}()

	go func() {
		for range time.Tick(packetLifetime) {
			packetsMu.Lock()
			for hash, info := range sentFrames {
				if time.Since(info) > packetLifetime {
					delete(sentFrames, hash)
				}
			}
			packetsMu.Unlock()
		}
	}()

	for evt := range node.Events() {
		if frm, ok := evt.(*gomavlib.EventFrame); ok {

			switch frm.Channel {

			case downChannel:
				node.WriteFrameExcept(downChannel, frm.Frame)

			case channelA, channelB:
				if frm.SystemID() != 1 || frm.ComponentID() != 1 {
					continue
				}
				packetsMu.Lock()
				if frm.Channel == channelA {
					receivedPacketsA++
				} else if frm.Channel == channelB {
					receivedPacketsB++
				}
				packetsMu.Unlock()

				// hash frame
				rawFrame := frm.Frame.GetRawFrame()
				frm.Frame.GetChecksum()
				hash := murmur3.Sum64(rawFrame)
				//node.WriteFrameTo(downChannel, frm.Frame)
				//lastActiveChannelMessageTime = time.Now()
				packetsMu.Lock()
				if _, ok := sentFrames[hash]; !ok {
					//fmt.Printf("new message!\n")
					node.WriteFrameTo(downChannel, frm.Frame)
					sentFrames[hash] = time.Now()
					notDuplicates++
				} else {
					duplicates++
					//fmt.Printf("old message!\n")
				}
				packetsMu.Unlock()
			}
		}
	}
}
