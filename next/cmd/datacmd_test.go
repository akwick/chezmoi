package cmd

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twpayne/go-vfs"

	"github.com/twpayne/chezmoi/next/internal/chezmoi"
	"github.com/twpayne/chezmoi/next/internal/chezmoitest"
)

func TestDataCmd(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		format string
		root   map[string]interface{}
	}{
		{
			format: "json",
			root: map[string]interface{}{
				"/home/user/.config/chezmoi/chezmoi.json": joinLines(
					`{`,
					`  "sourceDir": "/tmp/source",`,
					`  "data": {`,
					`    "test": true`,
					`  }`,
					`}`,
				),
			},
		},
		{
			format: "toml",
			root: map[string]interface{}{
				"/home/user/.config/chezmoi/chezmoi.toml": joinLines(
					`sourceDir = "/tmp/source"`,
					`[data]`,
					`  test = true`,
				),
			},
		},
		{
			format: "yaml",
			root: map[string]interface{}{
				"/home/user/.config/chezmoi/chezmoi.yaml": joinLines(
					`sourceDir: /tmp/source`,
					`data:`,
					`  test: true`,
				),
			},
		},
	} {
		tc := tc

		t.Run(tc.format, func(t *testing.T) {
			t.Parallel()

			chezmoitest.WithTestFS(t, tc.root, func(fs vfs.FS) {
				args := []string{
					"data",
					"--format", tc.format,
				}
				c := newTestConfig(t, fs)
				var sb strings.Builder
				c.stdout = &sb
				require.NoError(t, c.execute(args))

				var data struct {
					Chezmoi struct {
						SourceDir string `json:"sourceDir" toml:"sourceDir" yaml:"sourceDir"`
					} `json:"chezmoi" toml:"chezmoi" yaml:"chezmoi"`
					Test bool `json:"test" toml:"test" yaml:"test"`
				}
				assert.NoError(t, chezmoi.Formats[tc.format].Decode([]byte(sb.String()), &data))
				absSourceDir, err := filepath.Abs("/tmp/source")
				require.NoError(t, err)
				assert.Equal(t, filepath.ToSlash(absSourceDir), data.Chezmoi.SourceDir)
				assert.True(t, data.Test)
			})
		})
	}
}
