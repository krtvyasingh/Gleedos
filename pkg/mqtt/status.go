package mqtt

import "fmt"

func FormatClientStatusJSON(nodeID, state string) string {
	return fmt.Sprintf(`{"node_id":"%s","state":"%s"}`, nodeID, state)
}
