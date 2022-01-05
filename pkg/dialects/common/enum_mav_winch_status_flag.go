//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package common

import (
	"errors"
)

// Winch status flags used in WINCH_STATUS
type MAV_WINCH_STATUS_FLAG int

const (
	// Winch is healthy
	MAV_WINCH_STATUS_HEALTHY MAV_WINCH_STATUS_FLAG = 1
	// Winch thread is fully retracted
	MAV_WINCH_STATUS_FULLY_RETRACTED MAV_WINCH_STATUS_FLAG = 2
	// Winch motor is moving
	MAV_WINCH_STATUS_MOVING MAV_WINCH_STATUS_FLAG = 4
	// Winch clutch is engaged allowing motor to move freely
	MAV_WINCH_STATUS_CLUTCH_ENGAGED MAV_WINCH_STATUS_FLAG = 8
)

var labels_MAV_WINCH_STATUS_FLAG = map[MAV_WINCH_STATUS_FLAG]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_WINCH_STATUS_FLAG) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_WINCH_STATUS_FLAG[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_WINCH_STATUS_FLAG = map[string]MAV_WINCH_STATUS_FLAG{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_WINCH_STATUS_FLAG) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_WINCH_STATUS_FLAG[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_WINCH_STATUS_FLAG) String() string {
	if l, ok := labels_MAV_WINCH_STATUS_FLAG[e]; ok {
		return l
	}
	return "invalid value"
}