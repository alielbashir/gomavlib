//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Emit the value of a onboard parameter. The inclusion of param_count and param_index in the message allows the recipient to keep track of received parameters and allows him to re-request missing parameters after a loss or timeout. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html
type MessageParamValue struct {
	// Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Onboard parameter value
	ParamValue float32
	// Onboard parameter type.
	ParamType MAV_PARAM_TYPE `mavenum:"uint8"`
	// Total number of onboard parameters
	ParamCount uint16
	// Index of this onboard parameter
	ParamIndex uint16
}

// GetID implements the message.Message interface.
func (*MessageParamValue) GetID() uint32 {
	return 22
}