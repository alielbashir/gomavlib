//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Request a list of available logs. On some systems calling this may stop on-board logging until LOG_REQUEST_END is called. If there are no log files available this request shall be answered with one LOG_ENTRY message with id = 0 and num_logs = 0.
type MessageLogRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// First log id (0 for first available)
	Start uint16
	// Last log id (0xffff for last available)
	End uint16
}

// GetID implements the message.Message interface.
func (*MessageLogRequestList) GetID() uint32 {
	return 117
}