# AI Integration Roadmap (Guto CLI)

## 🎯 Goal
Transform the `Press` adapter into a provider-agnostic interface, allowing the user to choose between different AI models (Gemini, Claude, GPT-4, local Llama) via configuration.

## 🏗️ Proposed Architecture

### Current Interface
```go
type Press interface {
    Summarize(text string) (string, error)
}
```

### New Plugin Architecture
Create a provider registry in `internal/adapters/press/factory.go`:

```go
type AIProvider string

const (
    ProviderGemini  AIProvider = "gemini"
    ProviderOpenAI  AIProvider = "openai"
    ProviderClaude  AIProvider = "claude"
    ProviderOllama  AIProvider = "ollama"
)

func NewPressAdapter(provider AIProvider, config AIConfig) (ports.Press, error) {
    switch provider {
    case ProviderGemini:
        return &GeminiAdapter{Key: config.Key}, nil
    case ProviderOpenAI:
        return &OpenAIAdapter{Key: config.Key, Model: "gpt-4-turbo"}, nil
    // ...
    }
}
```

## 🛠️ Implementation Steps

### Phase 1: Abstraction (Short Term)
- [ ] Refactor `cmd/press.go` to instantiate the adapter based on `viper.GetString("ai.provider")`.
- [ ] Move the current `GeminiAdapter` logic to support the API key via an environment variable, in addition to the CLI wrapper.

### Phase 2: OpenAI Integration (Medium Term)
- [ ] Implement `OpenAIAdapter` using the official API (`https://api.openai.com/v1/chat/completions`).
- [ ] Required configuration: `ai.openai_key`.

### Phase 3: Claude Integration (Medium Term)
- [ ] Implement `ClaudeAdapter` (Anthropic API).
- [ ] Support for long context (100k+ tokens) for extensive meetings.

### Phase 4: Local Models (Long Term)
- [ ] Support for `ollama` or `llama.cpp` for total privacy (run `guto press` offline).

## 📝 Example Configuration (Future)
```yaml
ai:
  provider: "openai"
  model: "gpt-4"
  temperature: 0.3
  openai_key: "sk-..."
```
