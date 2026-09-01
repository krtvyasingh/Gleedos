package metrics

const GrafanaDashboardJSON = `{
  "title": "Gleedos High-Performance Engine Telemetry",
  "panels": [
    {
      "title": "Download Throughput (Bytes/sec)",
      "type": "graph",
      "targets": [{"expr": "rate(gleedos_bytes_total[1m])"}]
    }
  ]
}`

func GetDashboardTemplate() string {
	return GrafanaDashboardJSON
}
