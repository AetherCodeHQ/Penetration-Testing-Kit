
package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	host := "127.0.0.1"
	ports := []int{22, 80, 443, 8080}
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	if len(os.Args) > 2 {
		ports = ports[:0]
		for _, q := range strings.Split(os.Args[2], ",") {
			if n, err := strconv.Atoi(q); err == nil {
				ports = append(ports, n)
			}
		}
	}
	for _, p := range ports {
		addr := fmt.Sprintf("%s:%d", host, p)
		c, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if err != nil {
			fmt.Printf("%-24s closed\n", addr)
			continue
		}
		fmt.Printf("%-24s open\n", addr)
		c.Close()
	}
}
