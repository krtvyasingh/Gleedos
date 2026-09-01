package cloud

import "fmt"

type CloudProvider string

const (
	S3  CloudProvider = "s3"
	R2  CloudProvider = "r2"
	B2  CloudProvider = "b2"
	GCS CloudProvider = "gcs"
)

type UploadSpec struct {
	Provider CloudProvider
	Bucket   string
	Key      string
}

func FormatS3Endpoint(spec UploadSpec) string {
	return fmt.Sprintf("https://%s.%s.amazonaws.com/%s", spec.Bucket, spec.Provider, spec.Key)
}
