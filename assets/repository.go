package assets

import (
	"os"
	"path/filepath"
)

type Repository struct {
	Path   string
	Assets []Asset
}

func (r *Repository) AddAsset(path string) {
	r.Assets = append(r.Assets, Asset{Path: path})
}

func (r *Repository) KnownAssets() []Asset {
	return r.Assets
}

func (r *Repository) Init() error {
	filepath.Walk(r.Path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		r.AddAsset(path)

		return nil
	})

	return nil
}

func NewRepository(path string) AssetHandler {
	return &Repository{Path: path}
}
