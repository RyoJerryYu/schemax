package main

import (
	_ "embed"
	"testing"

	"buf.build/go/protoyaml"
	"github.com/RyoJerryYu/schemax/proto/schemax/recipe"
	"github.com/stretchr/testify/require"
)

//go:embed config.yaml
var cfg []byte

func TestUnmarshalConfig(t *testing.T) {
	var cfgSet recipe.Recipe
	err := protoyaml.Unmarshal(cfg, &cfgSet)
	require.NoError(t, err)
	t.Logf("%+v", &cfgSet)
}
