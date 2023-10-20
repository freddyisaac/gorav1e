package main

// #cgo CFLAGS: -I/usr/local/include/rav1e
//#include "rav1e.h"
//#include "glue.h"
import "C"
import "unsafe"

func (rax *RaContext) NewFrame() *RaFrame {
	return &RaFrame{
		frame: C.rav1e_frame_new(rax.ctx),
	}
}

type RaFrame struct {
	frame *C.RaFrame
}

func (raf *RaFrame) Unref() {
	C.rav1e_frame_unref(raf.frame)
	raf.frame = nil
}

func (raf *RaFrame) FillPlane(plane int, data []byte, stride int64, byteWidth int) {
	C.rav1e_frame_fill_plane(raf.frame, (C.int)(plane), (*C.uchar)(unsafe.Pointer(&data[0])), (C.ulong)(len(data)), (C.long)(stride), (C.int)(byteWidth))
}

func (raf *RaFrame) ExtractPlane(plane int, data []byte, stride int64, byteWidth int) {
	C.rav1e_frame_extract_plane(raf.frame, (C.int)(plane), (*C.uchar)(unsafe.Pointer(&data[0])), (C.ulong)(len(data)), (C.long)(stride), (C.int)(byteWidth))
}
