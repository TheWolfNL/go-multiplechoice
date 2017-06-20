# Go-MultipleChoice [![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/thewolfnl/go-multiplechoice)
If you're building a CLI application in GO and require the user to make a selection (1) out of given options, you can use this package.

> This currently requires the user to have `fzf` installed

## Install :package:
```
import (
    "github.com/thewolfnl/go-multiplechoice"
)
```

## Usage :radio_button:
```
selection := MultipleChoice.Selection("Select one: ", []string{"option1", "option2", "option3"}])
```

## Roadmap :rocket:
Our roadmap is located here: [Roadmap](./roadmap.md)