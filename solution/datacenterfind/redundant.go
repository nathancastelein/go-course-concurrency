package main

import (
	"context"
	"log/slog"
)

func Redundant(resourceName string, finders []Finder) {
	results := make(chan Result, len(finders))
	ctx, cancel := context.WithCancel(context.Background())

	for _, finder := range finders {
		go func() {
			results <- Result{
				datacenter: finder,
				found:      finder.FindWithContext(ctx, resourceName),
			}
		}()
	}

	result := <-results
	cancel()
	slog.Info("got result", slog.Any("datacenter", result.datacenter), slog.Bool("found", result.found))
}
