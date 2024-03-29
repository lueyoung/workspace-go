package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func isping(ip string) bool {
	recvBuf1 := make([]byte, 2048)
	payload := []byte{0x08, 0x00, 0x4d, 0x4b, 0x00, 0x01, 0x00, 0x10, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69}
	Time, _ := time.ParseDuration("3s")
	conn, err := net.DialTimeout("ip4:icmp", ip, Time)
	if err != nil {
		fmt.Println("bibi")
		return false
	}
	_, err = conn.Write(payload)
	if err != nil {
		return false
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	num, err := conn.Read(recvBuf1[0:])
	if err != nil {
		//check 80 3389 443 22 port
		Timetcp, _ := time.ParseDuration("1s")
		conn1, err := net.DialTimeout("tcp", ip+":80", Timetcp)
		if err == nil {
			defer conn1.Close()
			return true
		}

		conn2, err := net.DialTimeout("tcp", ip+":443", Timetcp)
		if err == nil {
			defer conn2.Close()
			return true
		}

		conn3, err := net.DialTimeout("tcp", ip+":3389", Timetcp)
		if err == nil {
			defer conn3.Close()
			return true
		}

		conn4, err := net.DialTimeout("tcp", ip+":22", Timetcp)
		if err == nil {
			defer conn4.Close()
			return true
		}

		return false
	}
	conn.SetReadDeadline(time.Time{})
	if string(recvBuf1[0:num]) != "" {
		return true
	}
	return false

}

func makeVIP(ip string) string {
	mostWanted := 240
	nums := strings.Split(ip, ".")
	for i := mostWanted; i < 254; i++ {
		ip := fmt.Sprintf("%v.%v.%v,%v", nums[0], nums[1], nums[2], i)
		fmt.Println(isping(ip))
		if !isping(ip) {
			return ip
		}
	}
	for i := 2; i < mostWanted-1; i++ {
		ip := fmt.Sprintf("%v.%v.%v,%v", nums[0], nums[1], nums[2], i)
		fmt.Println(isping(ip))
		if !isping(ip) {
			return ip
		}
	}
	return ""
}

func main() {
	ip := "192.168.100.81"
	//fmt.Println(makeVIP(ip))
		fmt.Println(isping(ip))
}
