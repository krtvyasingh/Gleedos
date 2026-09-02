package ebpf

type XDPAction int

const (
	XDPPass XDPAction = 0
	XDPDrop XDPAction = 1
)

type XDPFilter struct {
	TargetPort uint16
}

func NewXDPFilter(port uint16) *XDPFilter {
	return &XDPFilter{TargetPort: port}
}

func (x *XDPFilter) ProcessPacket(port uint16) XDPAction {
	if port == x.TargetPort {
		return XDPPass
	}
	return XDPPass
}
