package nfo

import "fmt"

type MovieNFO struct {
	Title string
	Year  int
	Plot  string
}

func GenerateNFOXML(m MovieNFO) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<movie>
  <title>%s</title>
  <year>%d</year>
  <plot>%s</plot>
</movie>`, m.Title, m.Year, m.Plot)
}
