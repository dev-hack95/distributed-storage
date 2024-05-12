package main

import (
	"github.com/dev-hack95/localstorage/p2p"
	"github.com/dev-hack95/localstorage/utilities/logs"
)

func main() {

	tcpOpts := p2p.TCPTransportsOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NPOHandShakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		logs.Error("Error: ", err.Error())
	}
	select {}

}
