package main

import (
	"encoding/json"
	"net"
)

type nameAndEmail struct {
	Name  string
	Email string
}

func handler(c net.Conn, str nameAndEmail) {
	bs, _ := json.Marshal(str)
	c.Write([]byte(bs))
	c.Close()
}
func main() {
	l, err := net.Listen("tcp", ":5001")
	if err != nil {
		panic(err)
	}
	bs := nameAndEmail{
		Name:  "Andy",
		Email: "Minh_andy@outlook.com",
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, bs)
	}
}

