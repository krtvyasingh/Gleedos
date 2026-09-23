package themes

type Theme struct {
	Name       string
	Background string
	Foreground string
	Accent     string
}

func GetTheme(name string) Theme {
	switch name {
	case "nord":
		return Theme{Name: "Nord", Background: "#2E3440", Foreground: "#D8DEE9", Accent: "#88C0D0"}
	case "dracula":
		return Theme{Name: "Dracula", Background: "#282A36", Foreground: "#F8F8F2", Accent: "#BD93F9"}
	default: // Catppuccin Mocha
		return Theme{Name: "Catppuccin Mocha", Background: "#1E1E2E", Foreground: "#CDD6F4", Accent: "#89B4FA"}
	}
}
