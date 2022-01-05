//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

type SCRIPTING_CMD int

const (
	// Start a REPL session.
	SCRIPTING_CMD_REPL_START SCRIPTING_CMD = 0
	// End a REPL session.
	SCRIPTING_CMD_REPL_STOP SCRIPTING_CMD = 1
)

var labels_SCRIPTING_CMD = map[SCRIPTING_CMD]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e SCRIPTING_CMD) MarshalText() ([]byte, error) {
	if l, ok := labels_SCRIPTING_CMD[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_SCRIPTING_CMD = map[string]SCRIPTING_CMD{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *SCRIPTING_CMD) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_SCRIPTING_CMD[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e SCRIPTING_CMD) String() string {
	if l, ok := labels_SCRIPTING_CMD[e]; ok {
		return l
	}
	return "invalid value"
}