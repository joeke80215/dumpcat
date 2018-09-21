package networklayer

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/joeke80215/dumpcat/pkg"
)

// tcpLayer tcp packet layer
type tcpLayer struct {
	layer *layer
}

// NewtcpLayer new tcp layer
func NewtcpLayer(packet gopacket.Packet) NetworkLayer {
	return &tcpLayer{
		layer: &layer{
			packet: packet,
		},
	}
}

// CheckLayer check tcp layer is exsist,if exsist return true or false
func (tcp *tcpLayer) CheckLayer() bool {
	if tcp.layer.packet.Layer(layers.LayerTypeTCP) == nil {
		return false
	}

	return true
}

// CheckPackageLayer check tcp package layer is exsist,if exsist return true or false
func (tcp *tcpLayer) CheckPackageLayer() bool {
	if err := tcp.layer.packet.Layer(layers.LayerTypeTCP).(*layers.TCP); err == nil {
		return false
	}

	return true
}

// Tran translate tcp gopacket.Packet to pkg.Packet
//
// refernece: https://en.wikipedia.org/wiki/Transmission_Control_Protocol
// ---PackTCP tcp package information---
// *Proto(string): portocal
// *IPv4(*PackIPv4): ipv4 package info
// *Src(string): source port
// *Dst(string): destination port
// *NS(bool): ECN-nonce - concealment protection
// *CWR(bool): Congestion Window Reduced (CWR) flag is set by the sending host to indicate that
//  it received a TCP segment with the ECE flag set and had responded in congestion control
//  mechanism (added to header by RFC 3168).
// *ECE(bool): ECN-Echo has a dual role, depending on the value of the SYN flag
// *URG(bool):  indicates that the Urgent pointer field is significant
// *Urgent: if the URG flag is set, then this 16-bit field is an offset from the sequence number
//  indicating the last urgent data byte
// *ACK(bool): indicates that the Acknowledgment field is significant. All packets after the
//  initial SYN packet sent by the client should have this flag set.
// *Ack(uint32): if the ACK flag is set then the value of this field is the next sequence number
//  that the sender of the ACK is expecting. This acknowledges receipt of all prior bytes (if any).
//  The first ACK sent by each end acknowledges the other end's initial sequence number itself, but no data.
// *PSH(bool): Push function. Asks to push the buffered data to the receiving application.
// *RST(bool): Reset the connection
// *SYN(bool): Synchronize sequence numbers. Only the first packet sent from each end should have this flag set.
//  Some other flags and fields change meaning based on this flag, and some are only valid when it is set,
//  and others when it is clear.
// *SEQ(uint32):
//  Has a dual role:
//  If the SYN flag is set (1), then this is the initial sequence number.
//  The sequence number of the actual first data byte and the acknowledged number in the corresponding
//  ACK are then this sequence number plus 1.
//  If the SYN flag is clear (0), then this is the accumulated sequence number of the first data byte of this segment
//  for the current session.
// *FIN(bool): Last packet from sender.
// *Window(uint16): The size of the receive window, which specifies the number of window size units (by default, bytes)
//  (beyond the segment identified by the sequence number in the acknowledgment field) that the sender of this segment is
//  currently willing to receive
// *Checksum(uint16): The 16-bit checksum field is used for error-checking of the header, the Payload and a Pseudo-Header.
//  The Pseudo-Header consists of the Source IP Address, the Destination IP Address, the protocol number for the TCP-Protocol (0x0006)
//  and the length of the TCP-Headers including Payload (in Bytes).
//
func (tcp *tcpLayer) Tran(p pkg.Packet) pkg.Packet {
	tcpipLayer := tcp.layer.packet.Layer(layers.LayerTypeTCP)
	tcps, _ := tcpipLayer.(*layers.TCP)
	p.Set("Src", fmt.Sprintf("%s:%s", p.Get("Src"), tcps.SrcPort.String()))
	p.Set("Dst", fmt.Sprintf("%s:%s", p.Get("Dst"), tcps.DstPort.String()))
	p.Set("Proto", "tcp")
	p.Set("NS", tcps.NS)
	p.Set("NECES", tcps.ECE)
	p.Set("URG", tcps.URG)
	p.Set("Urgent", tcps.Urgent)
	p.Set("ACK", tcps.ACK)
	p.Set("Ack", tcps.Ack)
	p.Set("PSH", tcps.PSH)
	p.Set("RST", tcps.RST)
	p.Set("SYN", tcps.SYN)
	p.Set("SEQ", tcps.Seq)
	p.Set("FIN", tcps.FIN)
	p.Set("Window", tcps.Window)
	p.Set("Checksum", tcps.Checksum)

	return p
}
