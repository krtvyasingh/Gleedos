package k8s

import "fmt"

func FormatResourceRequirements(cpuLimit, memLimit string) string {
	return fmt.Sprintf("resources:\n  limits:\n    cpu: %s\n    memory: %s\n", cpuLimit, memLimit)
}
