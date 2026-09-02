package chameleon

import "net/http"

func ApplyMobileWireProfile(req *http.Request, appName string) {
	if appName == "youtube_ios" {
		req.Header.Set("User-Agent", "com.google.ios.youtube/19.34.2 (iPhone15,2; U; CPU iOS 17_6 like Mac OS X; en_US)")
		req.Header.Set("X-YouTube-Client-Name", "5")
		req.Header.Set("X-YouTube-Client-Version", "19.34.2")
	}
}
