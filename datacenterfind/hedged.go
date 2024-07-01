package main

import (
	"log/slog"
)

func Hedged(resourceName string, finders []Finder) {
	for _, finder := range finders {
		slog.Info("launching find", slog.Any("datacenter", finder))
		found := finder.Find(resourceName)

		slog.Info("got result", slog.Bool("found", found))
	}
}
