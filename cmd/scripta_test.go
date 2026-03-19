package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/IgorGruvSS/guto/internal/ports/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ScriptaTestSuite struct {
	suite.Suite
	mockAudio  *mocks.MockAudioRecorder
	mockScribe *mocks.MockScribe
	mockPress  *mocks.MockPress
	tempDir    string
}

func (suite *ScriptaTestSuite) SetupTest() {
	suite.mockAudio = new(mocks.MockAudioRecorder)
	suite.mockScribe = new(mocks.MockScribe)
	suite.mockPress = new(mocks.MockPress)
	audioRecorder = suite.mockAudio
	scribeAdapter = suite.mockScribe
	pressAdapter = suite.mockPress

	var err error
	suite.tempDir, err = os.MkdirTemp("", "guto-scripta-test-*")
	suite.Require().NoError(err)
	viper.Set("output.base_dir", suite.tempDir)
}

func (suite *ScriptaTestSuite) TearDownTest() {
	os.RemoveAll(suite.tempDir)
	viper.Set("output.base_dir", "")
}

// TestScriptaCommand_SkipTranscription exercises the Listen→Stop path when the
// user declines transcription ("n").
func (suite *ScriptaTestSuite) TestScriptaCommand_SkipTranscription() {
	suite.mockAudio.On("Listen", mock.AnythingOfType("string")).Return(nil)
	suite.mockAudio.On("Stop").Return(nil)

	// stdin: Enter (stop recording), Enter (no title), "n" (skip scribe)
	restore := setStdin("\n\nn\n")
	defer restore()

	captureStdout(func() {
		scriptaCmd.Run(scriptaCmd, []string{})
	})

	suite.mockAudio.AssertExpectations(suite.T())
}

// TestScriptaCommand_FullPipeline exercises the complete path: listen → title →
// scribe → keep audio → press summary.
func (suite *ScriptaTestSuite) TestScriptaCommand_FullPipeline() {
	suite.mockAudio.On("Listen", mock.AnythingOfType("string")).Return(nil)
	suite.mockAudio.On("Stop").Return(nil)

	// Prepare a transcript file that scribeAdapter will "return"
	txtFile := filepath.Join(suite.tempDir, "transcript.txt")
	err := os.WriteFile(txtFile, []byte("meeting notes"), 0644)
	suite.Require().NoError(err)

	suite.mockScribe.On("Transcribe", mock.AnythingOfType("string")).Return(txtFile, nil)
	suite.mockPress.On("Summarize", mock.AnythingOfType("string")).Return("# Summary\nKey points.", nil)

	// stdin: Enter (stop), "TestMeeting" (title), "y" (scribe), "n" (keep audio), "y" (press)
	restore := setStdin("\nTestMeeting\ny\nn\ny\n")
	defer restore()

	captureStdout(func() {
		scriptaCmd.Run(scriptaCmd, []string{})
	})

	suite.mockAudio.AssertExpectations(suite.T())
	suite.mockScribe.AssertExpectations(suite.T())
	suite.mockPress.AssertExpectations(suite.T())
}

func TestScriptaTestSuite(t *testing.T) {
	suite.Run(t, new(ScriptaTestSuite))
}
