package p2p

import "net"

// Message is a struct of any data that is send over each transport
// betweeen two nodes
type Message struct {
	Form    net.Addr
	Payload []byte
}
