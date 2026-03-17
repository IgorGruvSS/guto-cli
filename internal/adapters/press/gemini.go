package press

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type GeminiAdapter struct{}

type GeminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

type GeminiModelsResponse struct {
	Models []struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Methods     []string `json:"supportedGenerationMethods"`
	} `json:"models"`
}

func (a *GeminiAdapter) ListModels() ([]string, error) {
	apiKey := viper.GetString("press.api_key")
	if apiKey == "" {
		return nil, fmt.Errorf("Guto Press: API Key não configurada. Use: guto config set press.api_key <sua_chave>")
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models?key=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com a API do Gemini: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta do Gemini: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API do Gemini retornou status %d: %s", resp.StatusCode, string(body))
	}

	var modelsResp GeminiModelsResponse
	if err := json.Unmarshal(body, &modelsResp); err != nil {
		return nil, fmt.Errorf("erro ao decodificar lista de modelos: %v", err)
	}

	var modelNames []string
	for _, m := range modelsResp.Models {
		// Filtra apenas modelos que suportam geração de conteúdo
		canGenerate := false
		for _, method := range m.Methods {
			if method == "generateContent" {
				canGenerate = true
				break
			}
		}
		if canGenerate {
			modelNames = append(modelNames, m.Name)
		}
	}

	return modelNames, nil
}

func (a *GeminiAdapter) Summarize(text string) (string, error) {
	apiKey := viper.GetString("press.api_key")
	if apiKey == "" {
		return "", fmt.Errorf("Guto Press: API Key não configurada. Use: guto config set press.api_key <sua_chave>")
	}

	model := viper.GetString("press.model")
	if model == "" {
		model = "gemini-2.5-flash"
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/%s:generateContent?key=%s", model, apiKey)

	prompt := fmt.Sprintf("Aja como um assistente de reuniões sênior. Baseado no texto abaixo, gere um sumário em Markdown (Resumo, Pontos Chave, Decisões e Próximos Passos). Responda apenas com o Markdown. Transcrição: %s", text)

	reqBody := GeminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: prompt},
				},
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("erro ao preparar requisição Gemini: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("erro ao conectar com a API do Gemini: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta do Gemini: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API do Gemini retornou status %d: %s", resp.StatusCode, string(body))
	}

	var geminiResp GeminiResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		return "", fmt.Errorf("erro ao decodificar resposta do Gemini: %v", err)
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		return geminiResp.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("Gemini não retornou nenhum conteúdo no sumário")
}
