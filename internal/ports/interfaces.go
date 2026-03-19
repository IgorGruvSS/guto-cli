package ports

type AudioRecorder interface {
	Listen(filename string) error
	Stop() error
}

type Scribe interface {
	Transcribe(inputPath string) (string, error)
}

type Press interface {
	Summarize(text string) (string, error)
	ListModels() ([]string, error)
}
