// +build ignore

package main

import (
	"fmt"
	"github.com/gswly/gomavlib"
	"github.com/gswly/gomavlib/dialects/ardupilotmega"
)

func main() {
	// create a node which
	// - communicates through an UDP endpoint in broadcast mode
	// - understands ardupilotmega dialect
	// - writes messages with given system id and component id
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointUdpBroadcast{BroadcastAddress: "192.168.7.255:5600"},
		},
		Dialect:     ardupilotmega.Dialect,
		SystemId:    10,
		ComponentId: 1,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	for {
		// wait until a message is received.
		res, ok := node.Read()
		if ok == false {
			break
		}

		// print message details
		fmt.Printf("received: id=%d, %+v\n", res.Message().GetId(), res.Message())
	}
}
