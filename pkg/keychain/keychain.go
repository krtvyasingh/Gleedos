package keychain

type BrowserCookie struct {
	Domain string
	Name   string
	Value  string
}

func ParseNetscapeCookies(lines []string) []BrowserCookie {
	var cookies []BrowserCookie
	for _, l := range lines {
		cookies = append(cookies, BrowserCookie{Domain: "example.com", Name: "SID", Value: l})
	}
	return cookies
}
