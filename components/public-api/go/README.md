# Golang bindings for Khulnasoft API
This package contains API definitions and client bindings for interacting with Khulnasoft API.

## Usage
```bash
go get -u github.com/khulnasoft-com/khulnasoft/components/public-api/go
```

```golang
import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/bufbuild/connect-go"
    "github.com/khulnasoft-com/khulnasoft/components/public-api/go/client"
    v1 "github.com/khulnasoft-com/khulnasoft/components/public-api/go/experimental/v1"
)

func ExampleListTeams() {
    token := "khulnasoft_pat_example.personal-access-token"

    khulnasoft, err := client.New(client.WithCredentials(token))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to construct khulnasoft client %v", err)
        return
    }

    response, err := khulnasoft.Teams.ListTeams(context.Background(), connect.NewRequest(&v1.ListTeamsRequest{}))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to list teams %v", err)
        return
    }

    fmt.Fprintf(os.Stdout, "Retrieved teams %v", response.Msg.GetTeams())
}
```

For more examples, see [examples](./examples) directory.
