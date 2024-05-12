package p2p

// Message is a struct of any data that is send over each transport
// betweeen two nodes
type Message struct {
	Payload []byte
}
