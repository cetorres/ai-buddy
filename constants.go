package main

import "fmt"

const GOOGLE_API_KEY_NAME = "GEMINI_API_KEY"
const OPENAI_API_KEY_NAME = "OPENAI_API_KEY"
const PATTERNS_DIR_ENV = "AI_BUDDY_PATTERNS"
const DEFAULT_MODEL_ENV = "AI_BUDDY_MODEL"
const TITLE = "AI Buddy 1.0 - Copyright © 2024 Carlos E. Torres (https://github.com/cetorres)"
var DESCRIPTION = fmt.Sprintf(`%s
An AI tool to help solving problems using a set of crowdsourced AI prompts.

Example usage:
	echo "Text to summarize..." | ai-buddy -p summarize
	ai-buddy -p summarize "Text to summarize..."
	ai-buddy -p summarize -m gemini-1.5-pro "Text to summarize..."
	ai-buddy -p summarize -m gpt-3.5-turbo "Text to summarize..."
	cat MyEssayText.txt | ai-buddy -p analyze_claims
	pbpaste | ai-buddy -p summarize | pbcopy
	cat text.txt | ai-buddy -p summarize -m gemini-1.5-pro

Commands:
  -p, --pattern <pattern_name> <prompt>  Specify a pattern and send prompt to model. Requires pattern name and prompt (also receive via pipe).
	-m, --model <name>                     Specify the model name to use.
	-l, --list                             List all available patterns.
	-v, --view <pattern_name>              View pattern prompt. Requires pattern name.
	-lm, --list-models                     List all available models.
	-h, --help                             Show this help.

Google Gemini API:
	- Get your API key at https://aistudio.google.com/app/apikey
	- Set an environment variable: export %s=<your_key_here>

OpenAI ChatGPT API:
	- Get your API key at https://platform.openai.com/api-keys
	- Set an environment variable: export %s=<your_key_here>

Default model to use:
	- By default, the model "gemini-1.5-pro" from Google or "gpt-3.5-turbo" from OpenAI are used, depending on the API KEY entered.
	- But you can set a custom default model via an environment variable: export %s=<model_name>

Patterns directory:
	- You can use the patterns directory in the same location of the binary (./patterns), this is by default.
	- Or you can set an environment variable if you want to move the binary to another directory.
	- Set the environment variable: export %s=<your_dir>/patterns`, TITLE, GOOGLE_API_KEY_NAME, OPENAI_API_KEY_NAME, DEFAULT_MODEL_ENV, PATTERNS_DIR_ENV)