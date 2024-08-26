package main

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	mcmodmeta "mc-mod-metadata/src"
	"os"
	"strings"
)

func stringFromFile(file *zip.File) (string, error) {
	fileReader, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileReader.Close()

	fileData, err := io.ReadAll(fileReader)
	if err != nil {
		return "", err
	}
	return string(fileData), nil
}

func readZipPart(file *zip.File) (string, error) {
	switch {
	case file.Name == "plugin.yml":
		{
			fileStr, err := stringFromFile(file)
			if err != nil {
				return "", err
			}

			plugin, err := mcmodmeta.NewBukkitPlugin(fileStr)
			if err != nil {
				return "", err
			}
			return plugin.Name, nil
		}
	case file.Name == "bungee.yml":
		{
			fileStr, err := stringFromFile(file)
			if err != nil {
				return "", err
			}

			plugin, err := mcmodmeta.NewBungeeCordPlugin(fileStr)
			if err != nil {
				return "", err
			}
			return plugin.Name, nil
		}
	case file.Name == "fabric.mod.json":
		{
			fileStr, err := stringFromFile(file)
			if err != nil {
				return "", err
			}

			fmt.Printf("%s\n", fileStr)

			plugin, err := mcmodmeta.NewFabricMod(fileStr)
			if err != nil {
				fmt.Println(err)
				return "", err
			}
			return plugin.ID, nil
		}
	case file.Name == "mcmod.info":
		{
			fileStr, err := stringFromFile(file)
			if err != nil {
				return "", err
			}

			plugin, err := mcmodmeta.NewForgeMod(fileStr)
			if err != nil {
				return "", err
			}
			return plugin.Mods[0].ModID, nil
		}
	case file.Name == "META-INF/mods.toml":
		{
			fileStr, err := stringFromFile(file)
			if err != nil {
				return "", err
			}

			plugin, err := mcmodmeta.NewForgeMod(fileStr)
			if err != nil {
				return "", err
			}
			return plugin.Mods[0].ModID, nil
		}
	case file.Name == "META-INF/neoforge.mods.toml":
		{
			fileStr, err := stringFromFile(file)
			if err != nil {
				return "", err
			}

			plugin, err := mcmodmeta.NewNeoForgeMod(fileStr)
			if err != nil {
				return "", err
			}

			return plugin.Mods[0].ModID, nil
		}
	case file.Name == "META-INF/sponge_plugins.json":
		{
			fileStr, err := stringFromFile(file)
			if err != nil {
				return "", err
			}

			plugin, err := mcmodmeta.NewSpongePlugin(fileStr)
			if err != nil {
				return "", err
			}
			return plugin.Plugins[0].ID, nil
		}
	case file.Name == "velocity-plugin.json":
		{
			fileStr, err := stringFromFile(file)
			if err != nil {
				return "", err
			}

			plugin, err := mcmodmeta.NewVelocityPlugin(fileStr)
			if err != nil {
				return "", err
			}
			return plugin.ID, nil
		}
	default:
		{
			return "", errors.New("unknown file")
		}
	}
}

func readJarFile(file string) ([]string, error) {
	zipListing, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer zipListing.Close()

	mods := make([]string, 0)
	for _, file := range zipListing.File {
		name, err := readZipPart(file)
		if err != nil {
			continue
		} else {
			mods = append(mods, name)
		}
	}

	return mods, nil
}

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
				mods, err := readJarFile(*inputDir + "/" + e.Name())
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(mods)
				}
			}
		}
	}

	if *file != "" {
		mods, err := readJarFile(*file)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(mods)
		}
	}
}
