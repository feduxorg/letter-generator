package assets

type Asset struct {
	Path string
}

func (a *Asset) GetPath() string {
	return a.Path
}
