package tagger

type Box struct {
	Name     string
	Children []Box
}

func NewBox(name string) Box {
	return Box{Name: name, Children: make([]Box, 0)}
}
