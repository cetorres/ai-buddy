# ai-buddy

A terminal command in Go that is an AI tool to help solving problems using a set of crowdsourced AI prompts.

It's heavily inspired by [Daniel Miessler](https://github.com/danielmiessler)'s tool [Fabric](https://github.com/danielmiessler/fabric). I created this as a simplified Go version, that's compiled, and probably a bit faster.

It's currently using only Google Gemini API.

## Build and usage

To build and run the program, just run:

```sh
$ go build
$ ./ai-buddy
```

## Google Gemini API

I used the [Google Gemini](https://gemini.google.com/app) API and the model `gemini-1.5-pro` to send the prompts to be processed.

You will need to obtain your own API key and set an environment variable with it:

```sh
$ export GEMINI_API_KEY=<your_key_here>
```

Get your free API key at <https://aistudio.google.com/app/apikey>.

## Patterns

Patterns are crowdsourced curated special prompts that improve the quality of the model's response for a given request.

Take a look at the [patterns](./patterns/) folder and check how they are created and work.

The current list of patterns was copied from the [Fabric](https://github.com/danielmiessler/fabric) project.

## Output

```
AI Buddy 1.0 - Copyright © 2024 Carlos E. Torres (https://github.com/cetorres)
An AI tool to help solving problems using a set of crowdsourced AI prompts.

Example usage:
        $ echo "Text to summarize..." | ai-buddy -p summarize
        $ ai-buddy -p summarize "Text to summarize..."
        $ cat MyEssayText.txt | ai-buddy -p analyze_claims

Commands:
        -p or --pattern : Specify a pattern and send prompt to model. Requires pattern name and prompt.
        -l or --list    : List all available patterns.
        -v or --view    : View pattern prompt. Requires pattern name.
        -h or --help    : Show this help.

Uses the Google Gemini API:
        - Get your API key at https://aistudio.google.com/app/apikey
        - Set an environment variable: export GEMINI_API_KEY=<your_key_here>
```

## More info

- Carlos E. Torres (<cetorres@cetorres.com>)
  - <https://cetorres.com>
  - <https://x.com/cetorres>

## Thanks

- [Daniel Miessler](https://github.com/danielmiessler) and all contributors from the [Fabric](https://github.com/danielmiessler/fabric) project for the great tool that inspired this one.
