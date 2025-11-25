# MapleStory OpenAPI (Go)

> Work in progress: this folder hosts the Go implementation of the MapleStory OpenAPI client. The first drop provides the shared `common` module (errors, enums, utilities, and the HTTP base client) so region-specific clients can be layered on next.

## Getting Started

```bash
cd go
go test ./...
```

## Usage Snapshot

```go
package main

import (
	"context"
	"fmt"

	"maplestory-openapi/api/tms"
)

const (
	apiKey    = "YOUR_API_KEY"
	character = "CHARACTER_NAME"
)

func main() {
	client := tms.NewClient(apiKey)
	character, err := client.GetCharacter(context.Background(), character)
	if err != nil {
		panic(err)
	}
	fmt.Println("OCID:", character.OCID)

	basic, err := client.GetCharacterBasic(context.Background(), character.OCID, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Basic:", basic)
}
```
