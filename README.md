# 🎙️ Guto (Your Personal Archivist)

**Guto** é uma ferramenta de linha de comando (CLI) em Go, inspirada no legado de Johannes Gutenberg. Sua missão é capturar, transcrever e "prensar" (resumir) o conhecimento gerado em reuniões e conversas, transformando áudio volátil em registros permanentes e organizados.

> *"Verba volant, scripta manent."* (As palavras voam, o que é escrito permanece.)

---

## ✨ Funcionalidades

- **`guto listen`**: Captura áudio do sistema (sink monitor) e do microfone simultaneamente usando FFmpeg.
- **`guto scribe`**: Transcreve arquivos de áudio para texto usando OpenAI Whisper (via `whisper-ctranslate2`).
- **`guto press`**: Processa transcrições e gera sumários executivos em Markdown via Google Gemini.
- **`guto scripta`**: Fluxo interativo completo (Gravar -> Titular -> Transcrever -> Resumir).
- **`guto config`**: Gerenciamento de configurações local via CLI.

## 🚀 Instalação Rápida

O Guto possui um instalador universal para as principais distribuições Linux (**Ubuntu/Debian, Fedora, Arch, openSUSE**).

### 1. Requisitos do Sistema
Certifique-se de ter o `git` e `go` instalados (o script tentará instalar se não houver).

### 2. Rodar o Instalador
```bash
git clone https://github.com/IgorGruvSS/guto-cli.git
cd guto-cli
chmod +x install.sh
sudo ./install.sh
```

O script irá:
1. Instalar dependências de sistema (`ffmpeg`, `pulseaudio-utils`, `python3`).
2. Compilar o binário Go e mover para `/usr/local/bin/guto`.
3. Criar um ambiente virtual isolado para o Whisper em `/opt/guto/whisper-env`.

---

## ⚙️ Configuração Pós-Instalação

Após rodar o instalador, você precisa configurar as chaves e caminhos no seu perfil de usuário (sem sudo):

### 1. Configurar o Scribe (Whisper)
Aponte para o ambiente Python criado pelo instalador:
```bash
guto config set scribe.python_bin /opt/guto/whisper-env/bin/python3
```

### 2. Configurar o Press (IA Gemini)
1. Obtenha uma chave de API gratuita no [Google AI Studio](https://aistudio.google.com/app/apikey).
2. Configure no Guto:
```bash
guto config set press.api_key SUA_CHAVE_AQUI
```

### 3. Listar e Escolher um Modelo de IA
O Guto permite que você escolha qual modelo do Gemini deseja usar (o padrão é `gemini-2.5-flash`). Para ver os modelos disponíveis:
```bash
guto config models
```
Para escolher um modelo da lista:
```bash
guto config set press.model gemini-1.5-pro
```

### 4. Verificar Configurações
```bash
guto config get
```

---

## 📖 Como Usar (O Fluxo Scripta)

O comando principal para o dia a dia é o `guto scripta`, que guia você por todo o processo:

1. **Gravação**: O Guto começa a ouvir o áudio do seu microfone e do sistema (perfeito para chamadas no Zoom/Teams/Meet).
2. **Encerramento**: Pressione `Enter` para parar a gravação.
3. **Titularização**: Dê um nome para a reunião (ex: `Daily-Sync`).
4. **Transcrição**: O Guto pergunta se deseja transcrever agora.
5. **Sumário**: O Guto gera um arquivo `.md` com o resumo executivo, decisões e próximos passos.

Os arquivos são organizados automaticamente no diretório `Output/`:
- `Output/audio/`: Matrizes originais `.wav`.
- `Output/scribe/`: Transcrições brutas `.txt`.
- `Output/press/`: Sumários finais `.md`.

---

## 🏗️ Arquitetura

O projeto segue os princípios da **Arquitetura Hexagonal (Ports and Adapters)**:
- `internal/ports/`: Define as interfaces de áudio, transcrição e IA.
- `internal/adapters/`: Implementações concretas (FFmpeg, Whisper, Gemini).
- `cmd/`: Orquestração via Cobra CLI.

Isso facilita a troca de provedores no futuro (ex: usar OpenAI GPT em vez de Gemini, ou Ollama local).

---

## 🛠️ Desenvolvimento Local

Para compilar manualmente:
```bash
go build -o guto main.go
./guto --help
```

---

## 📄 Licença
Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para detalhes.

---
*Guto: Devolvendo a permanência à palavra falada.*
