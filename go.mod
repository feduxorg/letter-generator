module github.com/feduxorg/letter-generator-go

go 1.12

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/feduxorg/letter-generator-go v1.0.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/libgit2/git2go v0.28.4
	github.com/pkg/errors v0.9.1
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.5
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	golang.org/x/sys v0.0.0-20210113181707-4bcb84eeeb78 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/libgit2/git2go => github.com/maxmeyer/git2go v0.28.4
