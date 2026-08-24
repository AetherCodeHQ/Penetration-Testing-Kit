package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: penetration-testing-kit <mode> [args]")
		fmt.Println("modes: wordlist <file> | brute <host> | recon <host>")
		os.Exit(1)
	}
	mode := os.Args[1]
	switch mode {
	case "wordlist":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "need wordlist file")
			os.Exit(1)
		}
		data, _ := os.ReadFile(os.Args[2])
		lines := strings.Split(string(data), "\n")
		words := 0
		unique := map[string]bool{}
		for _, line := range lines {
			t := strings.TrimSpace(line)
			if t != "" && !unique[t] {
				unique[t] = true
				words++
			}
		}
		fmt.Printf("wordlist: %d unique entries\n", words)
	case "brute":
		host := "127.0.0.1"
		if len(os.Args) > 2 {
			host = os.Args[2]
		}
		fmt.Printf("simulating brute-force on %s (demo mode)\n", host)
		for i := 0; i < 3; i++ {
			fmt.Printf("  attempt %d...\n", i+1)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("demo complete - no actual brute force performed")
	case "recon":
		host := "127.0.0.1"
		if len(os.Args) > 2 {
			host = os.Args[2]
		}
		fmt.Printf("recon target: %s\n", host)
		fmt.Printf("standard ports: 22, 80, 443, 3306, 5432\n")
		fmt.Println("use network-inspector for actual port scanning")
	default:
		fmt.Println("unknown mode:", mode)
		os.Exit(1)
	}
}
