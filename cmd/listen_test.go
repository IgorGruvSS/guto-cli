package cmd

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/IgorGruvSS/guto/internal/ports/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ListenTestSuite struct {
	suite.Suite
	mockAudio *mocks.MockAudioRecorder
	tempDir   string
}

func (suite *ListenTestSuite) SetupTest() {
	suite.mockAudio = new(mocks.MockAudioRecorder)
	audioRecorder = suite.mockAudio

	var err error
	suite.tempDir, err = os.MkdirTemp("", "guto-listen-test-*")
	suite.Require().NoError(err)
	viper.Set("output.base_dir", suite.tempDir)
}

func (suite *ListenTestSuite) TearDownTest() {
	os.RemoveAll(suite.tempDir)
	viper.Set("output.base_dir", "")
}

func (suite *ListenTestSuite) TestListenCommand_Success() {
	suite.mockAudio.On("Listen", mock.AnythingOfType("string")).Return(nil)
	suite.mockAudio.On("Stop").Return(nil)

	buf := new(bytes.Buffer)
	inBuf := bytes.NewBufferString("\n")
	listenCmd.SetOut(buf)
	listenCmd.SetIn(inBuf)

	listenCmd.Run(listenCmd, []string{})

	assert.Contains(suite.T(), buf.String(), "Guto is listening")
	assert.Contains(suite.T(), buf.String(), "Recording finished")
	suite.mockAudio.AssertExpectations(suite.T())
}

func (suite *ListenTestSuite) TestListenCommand_ListenError() {
	suite.mockAudio.On("Listen", mock.AnythingOfType("string")).Return(fmt.Errorf("device not found"))

	buf := new(bytes.Buffer)
	listenCmd.SetOut(buf)
	listenCmd.SetIn(bytes.NewBufferString(""))

	listenCmd.Run(listenCmd, []string{})

	assert.Contains(suite.T(), buf.String(), "Error starting recording")
	suite.mockAudio.AssertExpectations(suite.T())
}

func TestListenTestSuite(t *testing.T) {
	suite.Run(t, new(ListenTestSuite))
}
