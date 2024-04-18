package main

import (
	"github.com/dev-hack95/localstorage/p2p"
	"github.com/dev-hack95/localstorage/utilities/logs"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		logs.Error("Error: ", err.Error())
	}
	select {}

}
