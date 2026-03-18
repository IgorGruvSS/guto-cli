package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/IgorGruvSS/guto/internal/ports/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PressTestSuite struct {
	suite.Suite
	mockPress *mocks.MockPress
	tempDir   string
}

func (suite *PressTestSuite) SetupTest() {
	suite.mockPress = new(mocks.MockPress)
	pressAdapter = suite.mockPress
	
	// Create temp directory for output
	var err error
	suite.tempDir, err = os.MkdirTemp("", "guto-test-*")
	assert.NoError(suite.T(), err)
	viper.Set("output.base_dir", suite.tempDir)
}

func (suite *PressTestSuite) TearDownTest() {
	os.RemoveAll(suite.tempDir)
	viper.Set("output.base_dir", "")
}

func (suite *PressTestSuite) TestPressCommand_Success() {
	// Setup input file
	inputPath := filepath.Join(suite.tempDir, "test.txt")
	content := "This is a meeting transcript."
	err := os.WriteFile(inputPath, []byte(content), 0644)
	assert.NoError(suite.T(), err)

	// Mock behavior
	expectedSummary := "# Meeting Summary\nPoints discussed..."
	suite.mockPress.On("Summarize", content).Return(expectedSummary, nil)

	// Execute command
	buf := new(bytes.Buffer)
	pressCmd.SetOut(buf)
	
	pressCmd.Run(pressCmd, []string{inputPath})

	// Verify output file
	outputPath := filepath.Join(suite.tempDir, "press", "test.md")
	assert.FileExists(suite.T(), outputPath)
	
	savedContent, _ := os.ReadFile(outputPath)
	assert.Equal(suite.T(), expectedSummary, string(savedContent))
	
	suite.mockPress.AssertExpectations(suite.T())
}

func (suite *PressTestSuite) TestPressCommand_FileNotFound() {
	buf := new(bytes.Buffer)
	pressCmd.SetOut(buf)
	
	pressCmd.Run(pressCmd, []string{"nonexistent.txt"})
	
	assert.Contains(suite.T(), buf.String(), "Error reading file")
}

func TestPressTestSuite(t *testing.T) {
	suite.Run(t, new(PressTestSuite))
}
