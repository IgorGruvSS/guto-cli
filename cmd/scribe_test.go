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

type ScribeTestSuite struct {
	suite.Suite
	mockScribe *mocks.MockScribe
	tempDir    string
}

func (suite *ScribeTestSuite) SetupTest() {
	suite.mockScribe = new(mocks.MockScribe)
	scribeAdapter = suite.mockScribe
	
	// Create temp directory for output
	var err error
	suite.tempDir, err = os.MkdirTemp("", "guto-scribe-test-*")
	assert.NoError(suite.T(), err)
	viper.Set("output.base_dir", suite.tempDir)
}

func (suite *ScribeTestSuite) TearDownTest() {
	os.RemoveAll(suite.tempDir)
	viper.Set("output.base_dir", "")
}

func (suite *ScribeTestSuite) TestScribeCommand_Success() {
	// Setup input file
	inputPath := filepath.Join(suite.tempDir, "test.wav")
	err := os.WriteFile(inputPath, []byte("fake audio"), 0644)
	assert.NoError(suite.T(), err)

	// Mock behavior
	expectedTxt := filepath.Join(suite.tempDir, "test.txt")
	err = os.WriteFile(expectedTxt, []byte("transcribed text"), 0644)
	assert.NoError(suite.T(), err)
	
	suite.mockScribe.On("Transcribe", inputPath).Return(expectedTxt, nil)

	// Execute command with input for "n" to keep audio
	buf := new(bytes.Buffer)
	inBuf := bytes.NewBufferString("n\n")
	scribeCmd.SetOut(buf)
	scribeCmd.SetIn(inBuf)
	
	scribeCmd.Run(scribeCmd, []string{inputPath})

	// Verify output
	scribeDir := filepath.Join(suite.tempDir, "scribe")
	assert.DirExists(suite.T(), scribeDir)
	assert.FileExists(suite.T(), inputPath) // Audio should still exist
	
	suite.mockScribe.AssertExpectations(suite.T())
}

func (suite *ScribeTestSuite) TestScribeCommand_Success_DeleteAudio() {
	// Setup input file
	inputPath := filepath.Join(suite.tempDir, "delete.wav")
	err := os.WriteFile(inputPath, []byte("fake audio"), 0644)
	assert.NoError(suite.T(), err)

	// Mock behavior
	expectedTxt := filepath.Join(suite.tempDir, "delete.txt")
	err = os.WriteFile(expectedTxt, []byte("transcribed text"), 0644)
	assert.NoError(suite.T(), err)
	
	suite.mockScribe.On("Transcribe", inputPath).Return(expectedTxt, nil)

	// Execute command with input for "y" to delete audio
	buf := new(bytes.Buffer)
	inBuf := bytes.NewBufferString("y\n")
	scribeCmd.SetOut(buf)
	scribeCmd.SetIn(inBuf)
	
	scribeCmd.Run(scribeCmd, []string{inputPath})

	// Verify audio is gone
	_, err = os.Stat(inputPath)
	assert.True(suite.T(), os.IsNotExist(err))
	
	suite.mockScribe.AssertExpectations(suite.T())
}

func TestScribeTestSuite(t *testing.T) {
	suite.Run(t, new(ScribeTestSuite))
}
