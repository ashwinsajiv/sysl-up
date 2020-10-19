package main

import (
    "context"
    "log"

    "github.com/anz-bank/sysl-go/core"

    "github.com/ashwinsajiv/sysl-up/gen/pkg/servers/bff"
)

func main() {
    type AppConfig struct {
        // Define app-level config fields here.
    }
    log.Fatal(bff.Serve(context.Background(),
        func(ctx context.Context, config AppConfig) (*bff.ServiceInterface, *core.Hooks, error) {
            // Perform one-time setup based on config here.
            return &bff.ServiceInterface{
                // Add handlers here.
            }, nil, nil
        },
    ))
}
