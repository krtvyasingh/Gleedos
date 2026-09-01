package upscale

type UpscaleEngine string

const (
	RealESRGAN UpscaleEngine = "realesrgan"
	Waifu2x    UpscaleEngine = "waifu2x"
)

type UpscaleJob struct {
	Engine    UpscaleEngine
	Scale     int
	InputFile string
	OutputFile string
}

func BuildUpscaleArgs(job UpscaleJob) []string {
	return []string{"-i", job.InputFile, "-o", job.OutputFile, "-s", "4"}
}
