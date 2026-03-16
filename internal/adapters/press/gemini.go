package press

import (
    "os/exec"
    "fmt"
)

type GeminiAdapter struct{}

func (a *GeminiAdapter) Summarize(text string) (string, error) {
    prompt := fmt.Sprintf("Aja como um assistente de reuniões sênior. Baseado no texto abaixo, gere um sumário em Markdown (Resumo, Pontos Chave, Decisões e Próximos Passos). Responda apenas com o Markdown. Transcrição: %s", text)
    
    cmd := exec.Command("gemini", "ask", prompt)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", fmt.Errorf("erro ao chamar gemini: %v, output: %s", err, string(output))
    }
    
    return string(output), nil
}
