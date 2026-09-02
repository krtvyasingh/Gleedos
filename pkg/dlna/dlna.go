package dlna

const DeviceDescriptionXML = `<?xml version="1.0"?>
<root xmlns="urn:schemas-upnp-org:device-1-0">
  <device>
    <deviceType>urn:schemas-upnp-org:device:MediaServer:1</deviceType>
    <friendlyName>Gleedos DLNA Server</friendlyName>
  </device>
</root>`

func GetDeviceXML() string {
	return DeviceDescriptionXML
}
