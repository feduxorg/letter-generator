package converter

type PdfFile struct {
	Path string
}

func (p PdfFile) GetPath() string {
	return p.Path
}
