package exercise

import "fmt"

type IPAddr struct {
	Addr1 uint8
	Addr2 uint8
	Addr3 uint8
	Addr4 uint8
}

func (ipaddr *IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipaddr.Addr1, ipaddr.Addr2, ipaddr.Addr3, ipaddr.Addr4)
}
