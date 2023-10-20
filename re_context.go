package main

// #cgo CFLAGS: -I/usr/local/include/rav1e
//#include "rav1e.h"
//#include "glue.h"
import "C"

type RaContext struct {
	ctx *C.RaContext
}

func (rax *RaContext) Ctx() *C.RaContext {
	return rax.ctx
}

func (rax *RaContext) Unref() {
	if rax.ctx == nil {
		return
	}
	C.rav1e_context_unref(rax.ctx)
}

func (rax *RaContext) SendFrame(raf *RaFrame) int {
	cret := C.rav1e_send_frame(rax.ctx, raf.frame)
	return (int)(cret)
}

func (rax *RaContext) FlushFrames() {
	var nullFrame *C.RaFrame
	C.rav1e_send_frame(rax.ctx, nullFrame)
}

func (rax RaContext) ReceivePacket() (*RaPacket, int) {
	packet := &RaPacket{}
	cret := C.rav1e_receive_packet(rax.ctx, &packet.packet)
	return packet, (int)(cret)
}
