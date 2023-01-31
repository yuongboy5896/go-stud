package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

var lock sync.Mutex

type IpAlive struct {
	id     int
	ip     string
	status int
}

func main() {
	ScanIP()
}

func ScanIP() []IpAlive {
	start := time.Now()
	ip := "192.168.48.0"
	//
	ips := make(chan string, 255)
	ipslist := make([]IpAlive, 0)
	wg.Add(254)
	for i := 1; i <= 254; i++ {
		//fmt.Println(ip + strconv.Itoa(i))
		true_ip := ip + strconv.Itoa(i)
		//go ping(true_ip, ips)
		go pingg(true_ip, &ipslist)
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Println("执行时间:", cost)
	println(len(ips), cap(ips))
	fmt.Print(ips)
	return ipslist
}

func pingg(ip string, ips *[]IpAlive) {
	var beaf = "false"
	Command := fmt.Sprintf("ping -c 1 %s  > /dev/null && echo true || echo false", ip)
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	real_ip := strings.TrimSpace(string(output))

	if real_ip == beaf {
		fmt.Printf("IP: %s  失败\n", ip)
		ipAlive := IpAlive{}
		ipAlive.ip = ip
		ipAlive.status = 0
		(*ips) = append((*ips), ipAlive)

	} else {
		lock.Lock()
		ipAlive := IpAlive{}
		ipAlive.ip = ip
		ipAlive.status = 1
		(*ips) = append((*ips), ipAlive)
		lock.Unlock()
		fmt.Printf("IP: %s  成功 ping通\n", ip)
	}
	wg.Done()

}

func ping(ip string, ips chan string) {
	var beaf = "false"
	Command := fmt.Sprintf("ping -c 1 %s  > /dev/null && echo true || echo false", ip)
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	real_ip := strings.TrimSpace(string(output))

	if real_ip == beaf {
		fmt.Printf("IP: %s  失败\n", ip)
	} else {
		ips <- ip
		fmt.Printf("IP: %s  成功 ping通\n", ip)
	}
	wg.Done()

}
