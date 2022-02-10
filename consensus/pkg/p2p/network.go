package p2p

import (
	"fmt"
	"log"
	"net"

	"pocket/consensus/pkg/p2p/p2p_types"
	"pocket/consensus/pkg/types"
)

type network struct {
	p2p_types.Network

	AddrBook []*p2p_types.NetworkPeer
}

func ConnectToNetwork(validators types.ValMap) (n p2p_types.Network) {
	n = &network{}
	for nodeId, v := range validators {
		err := n.ConnectToValidator(nodeId, v)
		if err != nil {
			log.Println("Error connecting to validator: ", err)
			continue
		}
	}
	return
}

func (n *network) ConnectToValidator(nodeId types.NodeId, v *types.Validator) error {
	var tcpAddr, tcpAddrDebug *net.TCPAddr
	var err, errDebug error
	tcpAddr, err = net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", v.Host, v.Port))
	tcpAddrDebug, errDebug = net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", v.Host, v.DebugPort))
	if err != nil || errDebug != nil {
		return fmt.Errorf("error resolving addr: %v %v", err, errDebug)
	}
	peer := &p2p_types.NetworkPeer{
		ConsensusAddr: tcpAddr,
		DebugAddr:     tcpAddrDebug,
		NodeId:        nodeId,
		PublicKey:     v.PublicKey,
	}
	n.AddrBook = append(n.AddrBook, peer)
	return nil
}

func (n *network) NetworkBroadcast(data []byte, self types.NodeId) error {
	for _, peer := range n.AddrBook {
		// TODO: Discuss how self-broadcasts should be handled. Currently done internally in consensus.
		if self == peer.NodeId {
			continue
		}
		client, err := net.DialTCP("tcp", nil, peer.ConsensusAddr)
		if err != nil {
			log.Println("Error connecting to peer: ", err)
			continue
		}
		client.Write(data)
		client.Close()
	}

	return nil
}

func (n *network) NetworkSend(data []byte, node types.NodeId) error {
	for _, peer := range n.AddrBook {
		if node != peer.NodeId {
			continue
		}
		client, err := net.DialTCP("tcp", nil, peer.ConsensusAddr)
		if err != nil {
			log.Println("Error connecting to peer: ", err)
			continue
		}
		client.Write(data)
		client.Close()
		break
	}

	return nil
}

func (n *network) GetAddrBook() []*p2p_types.NetworkPeer {
	return n.AddrBook
}