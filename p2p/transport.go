package p2p

// Peer is an interface that represents the remote node
// When some one create a node
type Peer interface{}

// Transport is an interface that handles the communication
// between the nodes in the nwtwork. This can be of the
// from (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
}
