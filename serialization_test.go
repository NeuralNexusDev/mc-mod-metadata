package mcmodmeta_test

import (
	mcmodmeta "mc-mod-metadata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBukkitDeserialization(t *testing.T) {
	bukkitString := `name: TaterLib
version: 0.1.0
author: p0t4t0sandwich
description: some words
website: https://some.url
main: dev.neuralnexus.taterloader.platforms.BukkitLoaderPlugin
depends: [ ]
softdepends: [ LuckPerms ]
folia-supported: true
`

	bukkitPlugin, err := mcmodmeta.NewBukkitPlugin(bukkitString)

	assert.Nil(t, err)

	assert.Equal(t, "TaterLib", bukkitPlugin.Name)
	assert.Equal(t, "0.1.0", bukkitPlugin.Version)
	assert.Equal(t, "p0t4t0sandwich", bukkitPlugin.Author)
	assert.Equal(t, "some words", bukkitPlugin.Description)
	assert.Equal(t, "https://some.url", bukkitPlugin.Website)
	assert.Equal(t, "dev.neuralnexus.taterloader.platforms.BukkitLoaderPlugin", bukkitPlugin.Main)
	assert.Equal(t, 0, len(bukkitPlugin.Depends))
	assert.Equal(t, 1, len(bukkitPlugin.SoftDepends))
	assert.Equal(t, "LuckPerms", bukkitPlugin.SoftDepends[0])
	assert.Equal(t, true, bukkitPlugin.FoliaSupported)
}

func TestBungeeCordDeserialization(t *testing.T) {
	bungeeCordString := `name: TaterLib
version: 0.1.0
author: p0t4t0sandwich
description: some words
website: https://some.url
main: dev.neuralnexus.taterloader.platforms.BungeeCordLoaderPlugin
depends: [ ]
softdepends: [ LuckPerms ]
`

	bungeeCordPlugin, err := mcmodmeta.NewBungeeCordPlugin(bungeeCordString)

	assert.Nil(t, err)

	assert.Equal(t, "TaterLib", bungeeCordPlugin.Name)
	assert.Equal(t, "0.1.0", bungeeCordPlugin.Version)
	assert.Equal(t, "p0t4t0sandwich", bungeeCordPlugin.Author)
	assert.Equal(t, "some words", bungeeCordPlugin.Description)
	assert.Equal(t, "https://some.url", bungeeCordPlugin.Website)
	assert.Equal(t, "dev.neuralnexus.taterloader.platforms.BungeeCordLoaderPlugin", bungeeCordPlugin.Main)
	assert.Equal(t, 0, len(bungeeCordPlugin.Depends))
	assert.Equal(t, 1, len(bungeeCordPlugin.SoftDepends))
	assert.Equal(t, "LuckPerms", bungeeCordPlugin.SoftDepends[0])
}

func TestFabricDeserialization(t *testing.T) {
	fabricString := `{
	"schemaVersion": 1,
  "id": "taterlib",
  "version": "0.1.0",
  "name": "TaterLib",
  "description": "some words stuffs",
  "authors": [
    {
      "name": "p0t4t0sandwich",
      "contact": {
        "sources": "https://github.com/p0t4t0sandwich/",
        "homepage": "https://links.sperrer.ca/"
      }
    }
  ],
  "contact": {
    "repo": "https://some.repo.url"
  },
  "license": "GPL-3.0",
  "icon": "TaterLib.png.gz",
  "environment": "*",
  "entrypoints": {
    "main": [
      "dev.neuralnexus.taterloader.platforms.FabricLoaderPlugin"
    ]
  },
  "mixins": [
    "taterlib.mixins.v1_7_10.fabric.json",
    "taterlib.mixins.v1_8_9.fabric.json",
    "taterlib.mixins.v1_9_4.fabric.json",
    "taterlib.mixins.v1_10_2.fabric.json",
    "taterlib.mixins.v1_11_2.fabric.json",
    "taterlib.mixins.v1_12_2.fabric.json",
    "taterlib.mixins.v1_14_4.fabric.json",
    "taterlib.mixins.v1_15.fabric.json",
    "taterlib.mixins.v1_16.fabric.json",
    "taterlib.mixins.v1_17.fabric.json",
    "taterlib.mixins.v1_18.fabric.json",
    "taterlib.mixins.v1_19.fabric.json",
    "taterlib.mixins.v1_19_1.fabric.json",
    "taterlib.mixins.v1_19_1.fabric.patch.json",
    "taterlib.mixins.v1_19_3.fabric.patch.json",
    "taterlib.mixins.v1_19_4.fabric.patch.json",
    "taterlib.mixins.v1_20.fabric.json",
    "taterlib.mixins.v1_20_2.fabric.json",
    "taterlib.mixins.v1_20_2.fabric.patch.json",
    "taterlib.mixins.v1_20_6.fabric.json"
  ],
  "depends": {
    "fabricloader": ">=0.9.0",
    "minecraft": "*"
  },
  "suggests": {
    "fabric-api-base": "*",
    "legacy-fabric-api-base": "*",
    "luckperms": "*"
  }
}`
	fabricMod, err := mcmodmeta.NewFabricMod(fabricString)

	assert.Nil(t, err)

	assert.Equal(t, 1, fabricMod.SchemaVersion)
	assert.Equal(t, "taterlib", fabricMod.ID)
	assert.Equal(t, "0.1.0", fabricMod.Version)
	assert.Equal(t, "TaterLib", fabricMod.Name)
	assert.Equal(t, "some words stuffs", fabricMod.Description)
	assert.Equal(t, 1, len(fabricMod.Authors))
	assert.Equal(t, "p0t4t0sandwich", fabricMod.Authors[0].Name)
}

func TestForgeLegacyMod(t *testing.T) {
	forgeLegacyString := `{[
  {
    "modid": "taterlib",
    "name": "TaterLib",
    "license": "GPL-3.0",
    "description": "some words and stuffs",
    "version": "0.1.0",
    "mcversion": "1.12.2",
    "url": "https://some.random.url",
    "updateUrl": "https://some.update.url",
    "authorList": [
      "p0t4t0sandwich"
    ],
    "credits": "p0t4t0sandwich",
    "logoFile": "./TaterLib.png.gz",
    "screenshots": [],
    "useDependencyInformation": true,
    "dependencies": [],
    "dependants": []
  }
]`

	forgeLegacyMod, err := mcmodmeta.NewForgeLegacyMod(forgeLegacyString)

	assert.Nil(t, err)

	assert.Equal(t, "taterlib", forgeLegacyMod[0].ModID)
	assert.Equal(t, "TaterLib", forgeLegacyMod[0].Name)
	assert.Equal(t, "GPL-3.0", forgeLegacyMod[0].License)
	assert.Equal(t, "some words and stuffs", forgeLegacyMod[0].Description)
	assert.Equal(t, "0.1.0", forgeLegacyMod[0].Version)
	assert.Equal(t, "1.12.2", forgeLegacyMod[0].MCVersion)
	assert.Equal(t, "https://some.random.url", forgeLegacyMod[0].URL)
	assert.Equal(t, "https://some.update.url", forgeLegacyMod[0].UpdateURL)
	assert.Equal(t, "p0t4t0sandwich", forgeLegacyMod[0].AuthorList[0])
	assert.Equal(t, "p0t4t0sandwich", forgeLegacyMod[0].Credits)
	assert.Equal(t, "./TaterLib.png.gz", forgeLegacyMod[0].LogoFile)
	assert.Equal(t, 0, len(forgeLegacyMod[0].Screenshots))
	assert.Equal(t, true, forgeLegacyMod[0].UseDependencyInformation)
	assert.Equal(t, 0, len(forgeLegacyMod[0].Dependencies))
	assert.Equal(t, 0, len(forgeLegacyMod[0].Dependants))
}

func TestForgeMod(t *testing.T) {
	forgeModString := `
	modLoader = "javafml"
loaderVersion = "${loader_version_range}"
license = "${license}"
issueTrackerURL = "${issue_tracker_url}"

[[mods]]
modId = "${project_id}"
version = "${version}"
displayName = "${project_name}"
updateJSONURL = "${update_json_url}"
displayURL = "${homepage_url}"
logoFile = "${project_name}.png.gz"
credits = "${authors}"
authors = "${authors}"
displayTest = "IGNORE_SERVER_VERSION"
description = '''${project_description}'''

# Forge Dependency
[[dependencies.TaterLib]]
modId = "forge"
mandatory = false
type = "optional"
versionRange = "${forge_version_range}"
ordering = "NONE"
side = "BOTH"

# NeoForge Dependency
[[dependencies.TaterLib]]
modId = "neoforge"
mandatory = false
versionRange = "${neo_version_range}"
ordering = "NONE"
side = "BOTH"

# Minecraft Dependency
[[dependencies.TaterLib]]
modId = "minecraft"
mandatory = true
versionRange = "${minecraft_version_range}"
ordering = "NONE"
side = "BOTH"
`

	forgeMod, err := mcmodmeta.NewForgeMod(forgeModString)

	assert.Nil(t, err)

	assert.Equal(t, "javafml", forgeMod.ModLoader)
	assert.Equal(t, "${loader_version_range}", forgeMod.LoaderVersion)
	assert.Equal(t, "${license}", forgeMod.License)
	assert.Equal(t, "${issue_tracker_url}", forgeMod.IssueTrackerURL)
	assert.Equal(t, 1, len(forgeMod.Mods))
	assert.Equal(t, "${project_id}", forgeMod.Mods[0].ModID)
	assert.Equal(t, "${version}", forgeMod.Mods[0].Version)
	assert.Equal(t, "${project_name}", forgeMod.Mods[0].DisplayName)
	assert.Equal(t, "${update_json_url}", forgeMod.Mods[0].UpdateJSONURL)
	assert.Equal(t, "${homepage_url}", forgeMod.Mods[0].DisplayURL)
	assert.Equal(t, "${project_name}.png.gz", forgeMod.Mods[0].LogoFile)
	assert.Equal(t, "${authors}", forgeMod.Mods[0].Credits)
	assert.Equal(t, "${authors}", forgeMod.Mods[0].Authors)
	assert.Equal(t, "IGNORE_SERVER_VERSION", forgeMod.Mods[0].DisplayTest)
	assert.Equal(t, "${project_description}", forgeMod.Mods[0].Description)
	assert.Equal(t, 3, len(forgeMod.Dependencies))
	assert.Equal(t, "forge", forgeMod.Dependencies["TaterLib"][0].ModID)
	assert.Equal(t, false, forgeMod.Dependencies["TaterLib"][0].Mandatory)
	assert.Equal(t, "${forge_version_range}", forgeMod.Dependencies["TaterLib"][0].VersionRange)
	assert.Equal(t, "NONE", forgeMod.Dependencies["TaterLib"][0].Ordering)
	assert.Equal(t, "BOTH", forgeMod.Dependencies["TaterLib"][0].Side)
}
