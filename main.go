package main

import (
	"flag"
	"fmt"
	"log"
	mcmodmeta "mc-mod-metadata/src"
	"os"
	"strings"
)

func main() {
	// Flags
	// -h - Help
	// -f <file> - File to read
	// -j <jar> - Jar to read
	// -o <file> - File to write
	// -v - Verbose output

	file := flag.String("f", "", "File to read")
	inputDir := flag.String("i", "", "Directory to read")
	// outputDir := flag.String("o", "", "Directory to dump metadata to")

	flag.Parse()

	if *inputDir != "" {
		entries, err := os.ReadDir(*inputDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range entries {
			if !e.IsDir() && strings.Contains(e.Name(), ".jar") && strings.Contains(e.Name(), "fabric") {
				fmt.Println(e.Name())
				mods, err := mcmodmeta.ReadJarFile(*inputDir + "/" + e.Name())
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(mods)
				}
			}
		}
	}

	if *file != "" {
		mods, err := mcmodmeta.ReadJarFile(*file)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(mods)
		}
	}
}
