//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// The filtered global position (e.g. fused GPS and accelerometers). The position is in GPS-frame (right-handed, Z-up). It
// is designed as scaled integer message since the resolution of float is not sufficient.
type MessageGlobalPositionInt struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Latitude, expressed
	Lat int32
	// Longitude, expressed
	Lon int32
	// Altitude (MSL). Note that virtually all GPS modules provide both WGS84 and MSL.
	Alt int32
	// Altitude above ground
	RelativeAlt int32
	// Ground X Speed (Latitude, positive north)
	Vx int16
	// Ground Y Speed (Longitude, positive east)
	Vy int16
	// Ground Z Speed (Altitude, positive down)
	Vz int16
	// Vehicle heading (yaw angle), 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	Hdg uint16
}

// GetID implements the message.Message interface.
func (*MessageGlobalPositionInt) GetID() uint32 {
	return 33
}