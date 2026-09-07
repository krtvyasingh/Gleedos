package torrentv2

type FileNode struct {
	Length   int64
	PiecesRoot string
}

type FileTree struct {
	Files map[string]FileNode
}

func NewFileTree() *FileTree {
	return &FileTree{Files: make(map[string]FileNode)}
}

func (f *FileTree) AddFile(name string, length int64, piecesRoot string) {
	f.Files[name] = FileNode{Length: length, PiecesRoot: piecesRoot}
}
