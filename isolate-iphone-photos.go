package main

import (
	"log"
	"os"
	"path/filepath"

	processor "isolate-iphone-photos/processor"
)

func main() {
	log.SetFlags(log.Ltime)

	inputDir := os.Args[1]

	log.Printf("processing photo library export from %s\n", inputDir)

	processor.Run(&processor.Args{
		Root:                inputDir,
		TraverseDirTree:     filepath.Walk,
		MakeCollectMetadata: processor.MakeCollectMetadata,
		MakeWriteFile:       processor.MakeWriteFile,
	})

	log.Print("finished processing photo library export")
}