package webtest

import (
	"fastping"
	"fmt"
	"net"
	"time"
)

var judge bool = false

func PingTest(ping_url string) bool {
	judge = false
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ping_url) //ra为IP类型（也即[]byte）类型，记录IP地址
	if err != nil {
		fmt.Println(err)
		return false
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		judge = true
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return judge
}
