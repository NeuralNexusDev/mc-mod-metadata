package mcmodmeta

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
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

			plugin, err := NewBukkitPlugin(fileStr)
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

			plugin, err := NewBungeeCordPlugin(fileStr)
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

			plugin, err := NewFabricMod(fileStr)
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

			plugin, err := NewForgeMod(fileStr)
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

			plugin, err := NewForgeMod(fileStr)
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

			plugin, err := NewNeoForgeMod(fileStr)
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

			plugin, err := NewSpongePlugin(fileStr)
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

			plugin, err := NewVelocityPlugin(fileStr)
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

func ReadJarFile(file string) ([]string, error) {
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
