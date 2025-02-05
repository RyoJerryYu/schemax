package schemax

import (
	"context"
	"os"
	"testing"

	// drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"buf.build/go/protoyaml"
	"github.com/RyoJerryYu/schemax/proto/schemax/schema"
	"github.com/stretchr/testify/require"
	xocmd "github.com/xo/xo/cmd"
	xotype "github.com/xo/xo/types"
)

func TestLoad(t *testing.T) {
	ctx := context.Background()
	var err error

	urlstr := "mysql://schemaxuser:123456@127.0.0.1:13306/schemax?loc=Local&parseTime=true&charset=utf8mb4"
	ctx, err = open(ctx, urlstr, "")
	require.NoError(t, err)
	set := xotype.Set{}
	err = xocmd.LoadSchema(ctx, &set, xocmd.NewArgs(""))
	require.NoError(t, err)

	protoSet := schema.SetFromXO(&set)
	require.NotNil(t, protoSet)

	protoSetRaw, err := protoyaml.Marshal(protoSet)
	require.NoError(t, err)

	err = os.WriteFile("schema.yaml", protoSetRaw, os.ModePerm)
	require.NoError(t, err)
}
