module github.com/feduxorg/letter-generator

go 1.12

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/libgit2/git2go/v31 v31.4.7
	github.com/pkg/errors v0.9.1
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.2.2
	github.com/urfave/cli v1.22.5
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/sys v0.0.0-20210113181707-4bcb84eeeb78 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/libgit2/git2go/v31 => ./ext_deps/git2go
