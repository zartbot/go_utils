package vxlangpe

import (
	"encoding/binary"
	"fmt"
)

//  0                   1                   2                   3
//  0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |R|R|Ver|I|P|B|O|       Reserved                |  Next_Prot    |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                VXLAN Network Identifier (VNI) |   Reserved    |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//NextProto is a enum for next_proto
type NextProto uint8

const (
	_                         NextProto = iota
	NextProtoIPV4                       //0x01
	NextProtoIPV6                       //0x02
	NextProtoEthernet                   //0x03
	NextProtoMPLS                       //0x04
	NextProtoNSH                        //0x05
	NextProtoGroupBasedPolicy           //0x06
	NextProtoVBNG                       //0x07
	NextProtoINT                        //0x08,inband network telemetry

)

// VXLANGPE is a VXLAN packet header
type VXLANGPE struct {
	ValidIDFlag        bool      // 'I' bit
	ValidNextProtoFlag bool      // 'P' bit
	ValidBUMFlag       bool      //'B' bit
	NextProto          NextProto // 'NextProto
	VNI                uint32    // 'VXLAN Network Identifier' 24 bits
	Payload            []byte
}

//Decoder is used to decode vxlan gpe header
func Decoder(data []byte) (*VXLANGPE, error) {
	vx := &VXLANGPE{}
	if len(data) < 8 {
		return vx, fmt.Errorf("INVALID_VXLAN_DATA_LENGTH")
	}

	vx.ValidIDFlag = data[0]&0x08 > 0        // 'I' bit
	vx.ValidNextProtoFlag = data[0]&0x04 > 0 // 'P' bit
	vx.ValidBUMFlag = data[0]&0x02 > 0       // 'B' bit

	// VNI is a 24bit number, Uint32 requires 32 bits
	var buf [4]byte
	copy(buf[1:], data[4:7])
	vx.VNI = binary.BigEndian.Uint32(buf[:])
	if vx.ValidNextProtoFlag {
		vx.NextProto = NextProto(uint8(data[3]))
	}

	const vxlanLength = 8
	vx.Payload = data[vxlanLength:]
	return vx, nil
}
