package gorav1e

// #cgo CFLAGS: -I/usr/local/include
//#include "rav1e.h"
import "C"

type RaPacket struct {
	packet *C.RaPacket
}

func (rap *RaPacket) Unref() {
	if rap.packet == nil {
		return
	}
	C.rav1e_packet_unref(rap.packet)
}

func (rap *RaPacket) InputFrameNo() int {
	if rap.packet == nil {
		return -1
	}
	return (int)(rap.packet.input_frameno)
}
