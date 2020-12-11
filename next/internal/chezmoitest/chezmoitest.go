package chezmoitest

import (
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/twpayne/go-vfs"
	"github.com/twpayne/go-vfs/vfst"
)

// SkipUnlessGOOS calls t.Skip() if name does not match runtime.GOOS.
func SkipUnlessGOOS(t *testing.T, name string) {
	switch {
	case strings.HasSuffix(name, "_windows") && runtime.GOOS != "windows":
		t.Skip("skipping Windows-only test")
	case strings.HasSuffix(name, "_unix") && runtime.GOOS == "windows":
		t.Skip("skipping non-Windows test")
	}
}

// WithTestFS calls f with a test filesystem populated with root.
func WithTestFS(t *testing.T, root interface{}, f func(fs vfs.FS)) {
	fs, cleanup, err := vfst.NewTestFS(root)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	f(fs)
}
