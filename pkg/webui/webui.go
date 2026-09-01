package webui

import (
	"fmt"
	"net/http"
)

const IndexHTML = `<!DOCTYPE html>
<html>
<head><title>Gleedos Web Hub</title></head>
<body>
  <h1>Gleedos Media Engine</h1>
  <video controls width="640" src="/stream/video.mp4"></video>
</body>
</html>`

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, IndexHTML)
	})
}
