# go-terminal

golang cross platform terminal

## Usage

```go
package main

import (
    "github.com/chenhg5/go-terminal"
)

func main() {
	term, err := terminal.NewTerminal("-> ")
	if err != nil {
		panic(err)
	}
	defer term.Close()

	for {
		line, err := term.ReadLine()

		if err == io.EOF {
			fmt.Println()
			return
		}
		if err != nil {
			return
		}

		term.Write([]byte(line))
	}
}
```