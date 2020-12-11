package chezmoi

import (
	"testing"

	"github.com/stretchr/testify/require"
	vfs "github.com/twpayne/go-vfs"
	"github.com/twpayne/go-vfs/vfst"
)

func withTestFS(t *testing.T, root interface{}, f func(t *testing.T, fs vfs.FS)) {
	fs, cleanup, err := vfst.NewTestFS(root)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	f(t, fs)
}
