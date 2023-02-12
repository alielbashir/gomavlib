//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"errors"
)

// RADIO_LINK_STATS flags (bitmask).
type RADIO_LINK_STATS_FLAGS uint32

const (
	// Rssi are in negative dBm. Values 0..254 corresponds to 0..-254 dBm.
	RADIO_LINK_STATS_FLAGS_RSSI_DBM RADIO_LINK_STATS_FLAGS = 1
)

var labels_RADIO_LINK_STATS_FLAGS = map[RADIO_LINK_STATS_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e RADIO_LINK_STATS_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_RADIO_LINK_STATS_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_RADIO_LINK_STATS_FLAGS = map[string]RADIO_LINK_STATS_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *RADIO_LINK_STATS_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_RADIO_LINK_STATS_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e RADIO_LINK_STATS_FLAGS) String() string {
	if l, ok := labels_RADIO_LINK_STATS_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}