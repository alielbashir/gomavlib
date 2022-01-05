//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// A mapping of rover flight modes for custom_mode field of heartbeat.
type ROVER_MODE int

const (
	ROVER_MODE_MANUAL       ROVER_MODE = 0
	ROVER_MODE_ACRO         ROVER_MODE = 1
	ROVER_MODE_STEERING     ROVER_MODE = 3
	ROVER_MODE_HOLD         ROVER_MODE = 4
	ROVER_MODE_LOITER       ROVER_MODE = 5
	ROVER_MODE_FOLLOW       ROVER_MODE = 6
	ROVER_MODE_SIMPLE       ROVER_MODE = 7
	ROVER_MODE_AUTO         ROVER_MODE = 10
	ROVER_MODE_RTL          ROVER_MODE = 11
	ROVER_MODE_SMART_RTL    ROVER_MODE = 12
	ROVER_MODE_GUIDED       ROVER_MODE = 15
	ROVER_MODE_INITIALIZING ROVER_MODE = 16
)

var labels_ROVER_MODE = map[ROVER_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ROVER_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_ROVER_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ROVER_MODE = map[string]ROVER_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ROVER_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ROVER_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ROVER_MODE) String() string {
	if l, ok := labels_ROVER_MODE[e]; ok {
		return l
	}
	return "invalid value"
}