package main

// #cgo CFLAGS: -I/usr/local/include/rav1e
//#include "rav1e.h"
//#include "glue.h"
import "C"
import "unsafe"

type RaChromaticityPoint struct {
	x uint16
	y uint16
}

func (pt RaChromaticityPoint) ToC() C.RaChromaticityPoint {
	var cpt C.RaChromaticityPoint
	cpt.x = (C.ushort)(pt.x)
	cpt.y = (C.ushort)(pt.y)
	return cpt
}

type RaConfig struct {
	cfg                *C.RaConfig
	params             map[string]interface{}
	chromaticityPoints []RaChromaticityPoint
}

func NewRaConfig() *RaConfig {
	return &RaConfig{
		cfg:    C.rav1e_config_default(),
		params: make(map[string]interface{}),
	}
}

func (rac *RaConfig) Unref() {
	if rac.cfg == nil {
		return
	}
	C.rav1e_config_unref(rac.cfg)
	rac.cfg = nil
}

func (rac *RaConfig) Config() *C.RaConfig {
	return rac.cfg
}

func (rac *RaConfig) SetStringConfigValue(key, value string) int {
	ckey := C.CString(key)
	cvalue := C.CString(value)
	cret := C.rav1e_config_parse(rac.cfg, ckey, cvalue)
	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(cvalue))
	return (int)(cret)
}

func (rac *RaConfig) SetIntConfigValue(key string, value int) int {
	ckey := C.CString(key)
	cvalue := (C.int)(value)
	cret := C.rav1e_config_parse_int(rac.cfg, ckey, cvalue)
	C.free(unsafe.Pointer(ckey))
	return (int)(cret)
}

func (rac *RaConfig) SetConfigValue(key string, value interface{}) {
	switch t := value.(type) {
	case string:
		rac.params[key] = value
		rac.SetStringConfigValue(key, t)
	case int:
		rac.params[key] = value
		rac.SetIntConfigValue(key, t)
	default:
		// just ignore
	}
}

func (rac *RaConfig) SetColorDescription(matrix C.RaMatrixCoefficients, primaries C.RaColorPrimaries, transfer C.RaTransferCharacteristics) int {
	ret := C.rav1e_config_set_color_description(rac.cfg, matrix, primaries, transfer)
	return (int)(ret)
}

func (rac *RaConfig) SetMasteringDisplay(pts []RaChromaticityPoint, wp RaChromaticityPoint, maxLum, minLum uint32) int {
	rac.chromaticityPoints = pts
	cwp := wp.ToC()
	cpts := C.new_chromacity_point_array((C.int)(3))
	for _, pt := range pts {
		C.add_chromacity_point(cpts, pt.ToC())
	}
	cret := C.rav1e_config_set_mastering_display(rac.cfg, cpts.array, cwp, (C.uint)(maxLum), (C.uint)(minLum))
	C.free_chromacity_point_array(cpts)
	return (int)(cret)
}

func (rac *RaConfig) SetEmitData(emit int) {
	C.rav1e_config_set_emit_data(rac.cfg, (C.int)(emit))
}

type RaRational struct {
	num uint64 // numerator
	den uint64 // denominator
}

func (rar RaRational) ToC() C.RaRational {
	var crar C.RaRational
	crar.num = (C.ulong)(rar.num)
	crar.den = (C.ulong)(rar.den)
	return crar
}

func (rac *RaConfig) SetSampleAspectRatio(sar RaRational) {
	C.rav1e_config_set_sample_aspect_ratio(rac.cfg, sar.ToC())
}

func (rac *RaConfig) SetTimeBase(tb RaRational) {
	C.rav1e_config_set_time_base(rac.cfg, tb.ToC())
}

func (rac *RaConfig) SetPixelFormat(depth uint8, subSampling C.RaChromaSampling, chromaPos C.RaChromaSamplePosition, pixelRange C.RaPixelRange) int {
	cret := C.rav1e_config_set_pixel_format(rac.cfg, (C.uchar)(depth), subSampling, chromaPos, pixelRange)
	return (int)(cret)
}

func (rac *RaConfig) SetContentLight(maxContentLightLevel uint16, maxFrameAvgLightLevel uint16) int {
	cret := C.rav1e_config_set_content_light(rac.cfg, (C.ushort)(maxContentLightLevel), (C.ushort)(maxFrameAvgLightLevel))
	return (int)(cret)
}

func (rac *RaConfig) NewContext() *RaContext {
	return &RaContext{
		ctx: C.rav1e_context_new(rac.cfg),
	}
}
