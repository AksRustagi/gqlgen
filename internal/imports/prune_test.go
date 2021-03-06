package imports

import (
	"io/ioutil"
	"testing"

	"github.com/99designs/gqlgen/internal/code"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"
)

func TestPrune(t *testing.T) {
	// prime the packages cache so that it's not considered uninitialized
	code.RecordPackagesList([]*packages.Package{})

	b, err := Prune("testdata/unused.go", mustReadFile("testdata/unused.go"))
	require.NoError(t, err)
	require.Equal(t, string(mustReadFile("testdata/unused.expected.go")), string(b))
}

func mustReadFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}
