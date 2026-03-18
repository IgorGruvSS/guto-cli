package cmd

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetOutputPath(t *testing.T) {
	tests := []struct {
		name     string
		baseDir  string
		subDir   string
		expected string
	}{
		{
			name:     "Default output path",
			baseDir:  "",
			subDir:   "audio",
			expected: filepath.Join("Output", "audio"),
		},
		{
			name:     "Custom base directory",
			baseDir:  "/tmp/guto",
			subDir:   "scribe",
			expected: filepath.Join("/tmp/guto", "scribe"),
		},
		{
			name:     "Empty subdirectory",
			baseDir:  "/home/user/meetings",
			subDir:   "",
			expected: "/home/user/meetings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Set("output.base_dir", tt.baseDir)
			result := getOutputPath(tt.subDir)
			assert.Equal(t, tt.expected, result)
		})
	}
	// Reset viper for other tests
	viper.Set("output.base_dir", "")
}
