package upscale

import "testing"

func TestBuildUpscaleArgs(t *testing.T) {
	args := BuildUpscaleArgs(UpscaleJob{Engine: RealESRGAN, Scale: 4, InputFile: "in.mp4", OutputFile: "out.mp4"})
	if len(args) != 6 || args[1] != "in.mp4" {
		t.Errorf("unexpected args: %+v", args)
	}
}
