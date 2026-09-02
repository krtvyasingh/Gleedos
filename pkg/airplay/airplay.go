package airplay

type CastDevice struct {
	Name    string
	Address string
	Type    string
}

func DiscoverDevices() []CastDevice {
	return []CastDevice{
		{Name: "Living Room Apple TV", Address: "192.168.1.100:7000", Type: "airplay"},
	}
}
