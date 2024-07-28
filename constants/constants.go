package constants

import "fmt"

const APP_VERSION = "1.1"
const GOOGLE_API_KEY_NAME = "GEMINI_API_KEY"
const OPENAI_API_KEY_NAME = "OPENAI_API_KEY"
const PATTERNS_DIR_ENV = "AI_BUDDY_PATTERNS"
const DEFAULT_MODEL_ENV = "AI_BUDDY_MODEL"
const OLLAMA_HOST_ENV = "OLLAMA_HOST"
const AI_BUDDY_SERVER_PORT = 8080
const AI_BUDDY_SERVER_PORT_ENV = "AI_BUDDY_SERVER_PORT"
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
	-p, --pattern <pattern_name> <prompt>  Specify a pattern and send prompt to model. Requires pattern name and prompt (also receive via pipe).
	-o, --ollama                           Use Ollama local server. You should specify the model name available on your local Ollama server.
	-m, --model <model_name>               Specify the model name to use.
	-l, --list                             List all available patterns.
	-v, --view <pattern_name>              View pattern prompt. Requires pattern name.
	-lm, --list-models                     List all available models.
	-w, --webui                            Start an HTTP server with the web UI of the app.
	-h, --help                             Show this help.

Google Gemini API:
	- Get your API key at https://aistudio.google.com/app/apikey
	- Set an environment variable: export %s=<your_key_here>

OpenAI ChatGPT API:
	- Get your API key at https://platform.openai.com/api-keys
	- Set an environment variable: export %s=<your_key_here>

Default model to use:
	- By default, the model "gemini-1.5-pro" from Google or "gpt-3.5-turbo" from OpenAI are used, depending on the API KEY entered.
	- You can set a custom default model via an environment variable: export %s=<model_name>

Ollama:
	- To use Ollama (https://ollama.com), please download the Ollama app, install it and download an AI model.
	- It runs locally on your machine and can use free and open source models like llama3 or gemma2.
	- A list of all available models can be accessed at https://ollama.com/library.
	- You can set a custom Ollama URL via an environment variable: export %s=<ollama_url>

Patterns directory:
	- You can use the patterns directory in the same location of the binary (./patterns), this is by default.
	- Or you can set an environment variable if you want to move the binary to another directory.
	- Set the environment variable: export %s=<your_dir>/patterns
	
Web UI:
	- You can use the web UI as an alternative to interact with the app.
	- By default, it starts a local HTTP server on port 8080, but you can change the port.
	- Set an environment variable: export %s=<port>`, TITLE, GOOGLE_API_KEY_NAME, OPENAI_API_KEY_NAME, DEFAULT_MODEL_ENV, OLLAMA_HOST_ENV, PATTERNS_DIR_ENV, AI_BUDDY_SERVER_PORT_ENV)