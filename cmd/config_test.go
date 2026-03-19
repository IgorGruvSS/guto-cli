package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/IgorGruvSS/guto/internal/ports/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// captureStdout redirects os.Stdout during f() and returns the captured output.
// Needed because config.go uses fmt.Printf (not cmd.OutOrStdout()).
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r) //nolint:errcheck
	os.Stdout = old
	return buf.String()
}

// setStdin replaces os.Stdin with a pipe pre-loaded with input and returns a
// restore function. Used for commands that read via fmt.Scanln (not cobra's
// InOrStdin), such as scripta.go.
func setStdin(input string) func() {
	r, w, _ := os.Pipe()
	_, _ = w.WriteString(input)
	w.Close()
	old := os.Stdin
	os.Stdin = r
	return func() {
		r.Close()
		os.Stdin = old
	}
}

type ConfigTestSuite struct {
	suite.Suite
	mockPress *mocks.MockPress
	tempDir   string
}

func (suite *ConfigTestSuite) SetupTest() {
	suite.mockPress = new(mocks.MockPress)
	pressAdapter = suite.mockPress

	var err error
	suite.tempDir, err = os.MkdirTemp("", "guto-config-test-*")
	suite.Require().NoError(err)

	viper.Reset()
}

func (suite *ConfigTestSuite) TearDownTest() {
	os.RemoveAll(suite.tempDir)
	viper.Reset()
}

func (suite *ConfigTestSuite) TestConfigGet_AllSettings() {
	viper.Set("apikey", "abc123")
	viper.Set("timeout", "30")

	out := captureStdout(func() {
		configGetCmd.Run(configGetCmd, []string{})
	})

	assert.Contains(suite.T(), out, "apikey: abc123")
	assert.Contains(suite.T(), out, "timeout: 30")
}

func (suite *ConfigTestSuite) TestConfigGet_SpecificKey() {
	viper.Set("press.model", "gemini-pro")

	out := captureStdout(func() {
		configGetCmd.Run(configGetCmd, []string{"press.model"})
	})

	assert.Contains(suite.T(), out, "press.model: gemini-pro")
}

func (suite *ConfigTestSuite) TestConfigGet_KeyNotFound() {
	out := captureStdout(func() {
		configGetCmd.Run(configGetCmd, []string{"nonexistent"})
	})

	assert.Contains(suite.T(), out, "not found")
}

func (suite *ConfigTestSuite) TestConfigSet_WritesValue() {
	configFile := filepath.Join(suite.tempDir, "config.yaml")
	err := os.WriteFile(configFile, []byte(""), 0644)
	suite.Require().NoError(err)
	viper.SetConfigFile(configFile)

	captureStdout(func() {
		configSetCmd.Run(configSetCmd, []string{"press.model", "gemini-flash"})
	})

	assert.Equal(suite.T(), "gemini-flash", viper.GetString("press.model"))
}

func (suite *ConfigTestSuite) TestConfigModels_Success() {
	suite.mockPress.On("ListModels").Return([]string{"models/gemini-pro", "models/gemini-flash"}, nil)

	out := captureStdout(func() {
		configModelsCmd.Run(configModelsCmd, []string{})
	})

	assert.Contains(suite.T(), out, "gemini-pro")
	assert.Contains(suite.T(), out, "gemini-flash")
	suite.mockPress.AssertExpectations(suite.T())
}

func (suite *ConfigTestSuite) TestConfigModels_Error() {
	suite.mockPress.On("ListModels").Return([]string{}, fmt.Errorf("api error"))

	out := captureStdout(func() {
		configModelsCmd.Run(configModelsCmd, []string{})
	})

	assert.Contains(suite.T(), out, "Error listing models")
	suite.mockPress.AssertExpectations(suite.T())
}

func (suite *ConfigTestSuite) TestConfigAudioDevices() {
	// Just exercise the code path; output depends on whether pactl is available.
	captureStdout(func() {
		configAudioDevicesCmd.Run(configAudioDevicesCmd, []string{})
	})
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
