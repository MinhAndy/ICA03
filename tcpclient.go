package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":5001")
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	connbuf := bufio.NewReader(conn)
	for {
		str, err := connbuf.ReadString('\n')
		if len(str) > 0 {
			fmt.Println(str)
		}
		if err != nil {
			break
		}
	}
}
