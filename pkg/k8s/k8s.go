package k8s

import "fmt"

func GenerateCRDYAML(jobName, url string) string {
	return fmt.Sprintf(`apiVersion: gleedos.io/v1
kind: GleedosDownloadJob
metadata:
  name: %s
spec:
  url: "%s"
  threads: 8
`, jobName, url)
}
