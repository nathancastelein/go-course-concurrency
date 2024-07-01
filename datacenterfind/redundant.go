package main

import (
	"context"
	"log/slog"
	"sync"
)

func Redundant(resourceName string, finders []Finder) {
	results := make(chan Result, len(finders))
	var wg sync.WaitGroup

	for _, finder := range finders {
		wg.Add(1)
		go func() {
			defer wg.Done()
			slog.Info("starting find", slog.Any("datacenter", finder))
			results <- Result{
				datacenter: finder,
				found:      finder.FindWithContext(context.TODO(), resourceName),
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		slog.Info("got result", slog.Any("datacenter", result.datacenter), slog.Bool("found", result.found))
	}
}
