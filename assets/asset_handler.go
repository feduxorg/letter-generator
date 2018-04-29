package assets

type AssetHandler interface {
	AddAsset(string)
	Init() error
	KnownAssets() []Asset
}
