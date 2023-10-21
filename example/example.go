package main

// #cgo CFLAGS: -I/usr/local/include/rav1e
// #cgo LDFLAGS: -L/usr/local/lib -lrav1e
//#include "rav1e.h"
import "C"
import "gorav1e"

import (
	"fmt"
	"log"
)

const (
	width  = 64
	height = 96
	speed  = 9
)

func scaleU16(f float32, n uint16) uint16 {
	u := 1 << n
	sf := f * float32(u)
	return uint16(sf)
}

func scaleU32(f float32, n uint16) uint32 {
	u := 1 << n
	sf := f * float32(u)
	return uint32(sf)
}

func test_setup() {
	raConfig := gorav1e.NewRaConfig()

	raConfig.SetConfigValue("width", width)
	raConfig.SetConfigValue("height", height)
	raConfig.SetConfigValue("speed", speed)

	raConfig.SetColorDescription(C.RA_MATRIX_COEFFICIENTS_UNSPECIFIED, C.RA_COLOR_PRIMARIES_UNSPECIFIED, C.RA_TRANSFER_CHARACTERISTICS_UNSPECIFIED)

	primaries := []gorav1e.RaChromaticityPoint{
		{X: scaleU16(0.68, 16), Y: scaleU16(0.32, 16)},
		{X: scaleU16(0.265, 16), Y: scaleU16(0.69, 16)},
		{X: scaleU16(0.15, 16), Y: scaleU16(0.06, 16)},
	}

	wp := gorav1e.RaChromaticityPoint{X: scaleU16(0.31268, 16), Y: scaleU16(0.329, 16)}

	ret := raConfig.SetMasteringDisplay(primaries, wp, scaleU32(1000, 8), scaleU32(0, 14))

	if ret != 0 {
		log.Printf("SetMasteringDisplay error %v", ret)
	}

	raCtx := raConfig.NewContext()

	raFrame := raCtx.NewFrame()

	limit := 30
	tries := 0
	for i := 0; i < limit; i++ {
		fmt.Printf("sending frame : %d\n", i)
		ret := raCtx.SendFrame(raFrame)
		if ret < 0 {
			goto clean
		} else {
			if ret > 0 {
				log.Printf("unable to append frame %s to internal queue", i)
			}
		}
	}

	raCtx.FlushFrames()

	for i := 0; i < limit+5; {
		tries++
		if tries > 2*limit {
			goto clean
		}
		packet, ret := raCtx.ReceivePacket()
		if ret < 0 {
			fmt.Printf("ReceivePacket error : %d\n", ret)
			goto clean
		} else {
			switch ret {
			case C.RA_ENCODER_STATUS_SUCCESS:
				log.Printf("ReceivePacket frame number %d\n", packet.InputFrameNo())
				packet.Unref()
				i++
			case C.RA_ENCODER_STATUS_LIMIT_REACHED:
				log.Printf("limit reached at %d\n", i)
				goto clean
			}
		}
	}

clean:
	raFrame.Unref()
	raConfig.Unref()
	raCtx.Unref()
}

func main() {

	fmt.Printf("starting!\n")

	test_setup()

}
