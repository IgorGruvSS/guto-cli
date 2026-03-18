package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockAudioRecorder struct {
	mock.Mock
}

func (m *MockAudioRecorder) Listen(filename string) error {
	args := m.Called(filename)
	return args.Error(0)
}

func (m *MockAudioRecorder) Stop() error {
	args := m.Called()
	return args.Error(0)
}

type MockScribe struct {
	mock.Mock
}

func (m *MockScribe) Transcribe(inputPath string) (string, error) {
	args := m.Called(inputPath)
	return args.String(0), args.Error(1)
}

type MockPress struct {
	mock.Mock
}

func (m *MockPress) Summarize(text string) (string, error) {
	args := m.Called(text)
	return args.String(0), args.Error(1)
}

func (m *MockPress) ListModels() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}
