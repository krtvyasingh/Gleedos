package hdr

import "bytes"

type HDRProfile string

const (
	HDR10       HDRProfile = "hdr10"
	DolbyVision HDRProfile = "dolby_vision"
	SDR         HDRProfile = "sdr"
)

func DetectHDRProfile(nalUnits []byte) HDRProfile {
	if bytes.Contains(nalUnits, []byte("dvh1")) || bytes.Contains(nalUnits, []byte("dvhe")) {
		return DolbyVision
	}
	if bytes.Contains(nalUnits, []byte("mdcv")) || bytes.Contains(nalUnits, []byte("clli")) {
		return HDR10
	}
	return SDR
}
