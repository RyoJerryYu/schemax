//go:build tools

//lint:file-ignore U1000 Ignore all unused code, it's used in go:build tools

package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/xo/xo"
)
