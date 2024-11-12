package constants

type OrderStatus uint

const (
	New         OrderStatus = iota + 1 //1
	OnKitchen                          //2
	WaitDeliver                        //3
	OnTheWay                           //4
	Delivered                          //5
	Cancelled                          //6
)

func (s OrderStatus) String() string {
	switch s {
	case New:
		return "New"
	case OnKitchen:
		return "On kitchen"
	case WaitDeliver:
		return "Wait deliver"
	case OnTheWay:
		return "On the-way"
	case Delivered:
		return "Delivered"
	case Cancelled:
		return "Closed"
	default:
		return "Unknown"
	}
}

func (s OrderStatus) ToUInt() uint {
	return uint(s)
}
