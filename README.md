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

	"maplestory-openapi/api/common"
	"maplestory-openapi/api/tms"
)

func main() {
	client := common.NewClient("YOUR_API_KEY", "maplestory", common.FixedOffsetLocation(540))
	date := client.ProperDefaultDate(common.LatestAPIUpdateTimeOptions{Hour: 2, Minute: 0})

	fmt.Println("Latest retrievable date:", date)

	// Region-specific clients will embed this common client; use Fetch directly for quick experiments.
	var payload map[string]any
	_ = client.Fetch(context.Background(), "maplestory/v1/id", map[string]string{
		"character_name": "example",
	}, &payload)

	// TMS example
	tmsClient := tms.NewClient("YOUR_API_KEY")
	character, err := tmsClient.GetCharacter(context.Background(), "characterName")
	if err != nil {
		panic(err)
	}
	fmt.Println("OCID:", character.OCID)
}
```
