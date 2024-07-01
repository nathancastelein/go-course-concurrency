package main

import (
	"log/slog"
	"time"
)

func Hedged(resourceName string, finders []Finder) {
	results := make(chan bool)

	for _, finder := range finders {
		timer := time.NewTimer(75 * time.Millisecond)
		go func() {
			slog.Info("launching find", slog.Any("datacenter", finder))
			results <- finder.Find(resourceName)
		}()

		select {
		case result := <-results:
			slog.Info("got result", slog.Bool("found", result))
			return
		case <-timer.C:
			continue
		}
	}
}
