//go:build tools

package tools

// tool dependencies
import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql"
	_ "golang.org/x/tools/cmd/godoc"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
)

//go:generate go install -v github.com/golang/mock/mockgen@v1.6.0
//go:generate go install -v golang.org/x/tools/cmd/goimports
//go:generate go install -v golang.org/x/tools/cmd/godoc
//go:generate go install -v honnef.co/go/tools/cmd/staticcheck
//go:generate go install -v github.com/volatiletech/sqlboiler/v4
//go:generate go install -v github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
//go:generate go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest
