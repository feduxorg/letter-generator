module github.com/fedux-org/letter-generator-go

go 1.12

require (
	github.com/libgit2/git2go v0.28.4
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.2.2
	github.com/urfave/cli v1.22.2
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	gopkg.in/yaml.v2 v2.2.7
)

replace github.com/libgit2/git2go => github.com/maxmeyer/git2go v0.28.4
