package constants

import "fmt"

const APP_VERSION = "1.1.8"
const OLLAMA_HOST_ENV = "OLLAMA_HOST"
const AI_BUDDY_SERVER_PORT = 8080
var TITLE = fmt.Sprintf("ai-buddy %s - Created by Carlos E. Torres (https://github.com/cetorres)", APP_VERSION)
var DESCRIPTION = fmt.Sprintf(`%s
AI tool to help solving problems using prompt engineering from a set of crowdsourced AI prompts.

Example usage:
	echo "Text to summarize..." | ai-buddy -p summarize
	ai-buddy -p summarize "Text to summarize..."
	ai-buddy -p summarize -m gemini-1.5-pro "Text to summarize..."
	ai-buddy -p summarize -m gpt-3.5-turbo "Text to summarize..."
	cat MyEssayText.txt | ai-buddy -p analyze_claims
	pbpaste | ai-buddy -p summarize | pbcopy
	cat text.txt | ai-buddy -p summarize -m gemini-1.5-pro
	ai-buddy -p summarize -o -m llama3 "Text to summarize..."
	ai-buddy --webui

Commands:
	-s, --setup                            Set up the app with the necessary configuration.
	-p, --pattern <pattern_name> <prompt>  Specify a pattern and send prompt to model. Requires pattern name and prompt (also receive via pipe).
	-o, --ollama                           Use Ollama local server. You should specify the model name available on your local Ollama server.
	-m, --model <model_name>               Specify the model name to use.
	-l, --list                             List all available patterns.
	-v, --view <pattern_name>              View pattern prompt. Requires pattern name.
	-lm, --list-models                     List all available models.
	-w, --webui [--port <port_number>]     Start an HTTP server with the web UI of the app. Optional argument: --port <port_number>.
	-h, --help                             Show this help.

Google Gemini API:
	- To use this API, enter the key on the setup command.
	- Get your API key at https://aistudio.google.com/app/apikey

OpenAI ChatGPT API:
	- To use this API, enter the key on the setup command.
	- Get your API key at https://platform.openai.com/api-keys

Anthropic Claude API:
	- To use this API, enter the key on the setup command.
	- Get your API key at https://console.anthropic.com/settings/keys

Ollama:
	- To use Ollama (https://ollama.com), please download the Ollama app, install it and download an AI model.
	- It runs locally on your machine and can use free and open source models like llama3 or gemma2.
	- A list of all available models can be accessed at https://ollama.com/library.
	- You can set a custom Ollama URL via an environment variable: export %s=<ollama_url>

Web UI:
	- You can use the web UI as an alternative to interact with the app.
	- By default, it starts a local HTTP server on port 8080, but you can change the port with --port <port_number>.`, TITLE, OLLAMA_HOST_ENV)