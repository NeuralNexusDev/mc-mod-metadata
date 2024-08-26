package mcmodmeta_test

import (
	mcmodmeta "mc-mod-metadata/src"
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
  },
  "apple"
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
  "taterlib.mixins.v1_7_10.fabric.json"
],
"accessWidener": "taterlib.accesswidener",
"depends": {
  "fabricloader": ">=0.9.0",
  "minecraft": "*"
},
"suggests": {
  "fabric-api-base": "*",
  "legacy-fabric-api-base": "*",
  "luckperms": "*"
},
"breaks": {
  "taterlib": "0.0.1"
},
"custom": {
  "some": {}
},
"jars": [
  {
	"file": "META-INF/jars/thing-0.1.0.jar"
  }
]
}`
	fabricMod, err := mcmodmeta.NewFabricMod(fabricString)

	assert.Nil(t, err)

	assert.Equal(t, 1, fabricMod.SchemaVersion)
	assert.Equal(t, "taterlib", fabricMod.ID)
	assert.Equal(t, "0.1.0", fabricMod.Version)
	assert.Equal(t, "TaterLib", fabricMod.Name)
	assert.Equal(t, "some words stuffs", fabricMod.Description)
	assert.Equal(t, 2, len(fabricMod.Authors))
	assert.Equal(t, "p0t4t0sandwich", fabricMod.Authors[0].(mcmodmeta.FabricPerson).Name)
}

func TestForgeLegacyMod(t *testing.T) {
	forgeLegacyString := `[{
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
}]`

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
loaderVersion = "[1,)"
license = "GPL-3.0"
issueTrackerURL = "https://some.issue.tracker"

[[mods]]
modId = "taterlib"
version = "0.1.0"
displayName = "TaterLib"
updateJSONURL = "https://some.update.url"
displayURL = "https://some.homepage.url"
logoFile = "TaterLib.png.gz"
credits = "p0t4t0sandwich"
authors = "p0t4t0sandwich"
displayTest = "IGNORE_SERVER_VERSION"
description = '''some more word stuffs'''

# Forge Dependency
[[dependencies.TaterLib]]
modId = "forge"
mandatory = false
type = "optional"
versionRange = "[30,)"
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
	assert.Equal(t, "[1,)", forgeMod.LoaderVersion)
	assert.Equal(t, "GPL-3.0", forgeMod.License)
	assert.Equal(t, "https://some.issue.tracker", forgeMod.IssueTrackerURL)
	assert.Equal(t, 1, len(forgeMod.Mods))
	assert.Equal(t, "taterlib", forgeMod.Mods[0].ModID)
	assert.Equal(t, "0.1.0", forgeMod.Mods[0].Version)
	assert.Equal(t, "TaterLib", forgeMod.Mods[0].DisplayName)
	assert.Equal(t, "https://some.update.url", forgeMod.Mods[0].UpdateJSONURL)
	assert.Equal(t, "https://some.homepage.url", forgeMod.Mods[0].DisplayURL)
	assert.Equal(t, "TaterLib.png.gz", forgeMod.Mods[0].LogoFile)
	assert.Equal(t, "p0t4t0sandwich", forgeMod.Mods[0].Credits)
	assert.Equal(t, "p0t4t0sandwich", forgeMod.Mods[0].Authors)
	assert.Equal(t, "IGNORE_SERVER_VERSION", forgeMod.Mods[0].DisplayTest)
	assert.Equal(t, "some more word stuffs", forgeMod.Mods[0].Description)
	assert.Equal(t, 2, len(forgeMod.Dependencies["TaterLib"]))
	assert.Equal(t, "forge", forgeMod.Dependencies["TaterLib"][0].ModID)
	assert.Equal(t, false, forgeMod.Dependencies["TaterLib"][0].Mandatory)
	assert.Equal(t, "[30,)", forgeMod.Dependencies["TaterLib"][0].VersionRange)
	assert.Equal(t, "NONE", forgeMod.Dependencies["TaterLib"][0].Ordering)
	assert.Equal(t, "BOTH", forgeMod.Dependencies["TaterLib"][0].Side)
}

func TestNeoForgeMod(t *testing.T) {
	neoForgeModString := `
modLoader = "javafml"
loaderVersion = "[1,)"
license = "GPL-3.0"
issueTrackerURL = "https://some.issue.tracker"

[[mods]]
modId = "taterlib"
version = "0.1.0"
displayName = "TaterLib"
updateJSONURL = "https://some.update.url"
displayURL = "https://some.home.url"
logoFile = "TaterLib.png.gz"
credits = "p0t4t0sandwich"
authors = "p0t4t0sandwich"
displayTest = "IGNORE_SERVER_VERSION"
description = '''some descriptive words'''

[[mixins]]
config = "taterlib.mixins.v1_20.vanilla.json"
[[mixins]]
config = "taterlib.mixins.v1_20_2.vanilla.patch.json"
[[mixins]]
config = "taterlib.mixins.v1_20_2.vanilla.json"
[[mixins]]
config = "taterlib.mixins.v1_20_6.vanilla.patch.json"
[[mixins]]
config = "taterlib.mixins.v1_20_6.vanilla.json"
[[mixins]]
config = "taterlib.mixins.v1_21.vanilla.json"

# NeoForge Dependency
[[dependencies.taterlib]]
modId = "neoforge"
type = "required"
versionRange = "${neo_version_range}"
ordering = "NONE"
side = "BOTH"

# Minecraft Dependency
[[dependencies.taterlib]]
modId = "minecraft"
type = "required"
versionRange = "${minecraft_version_range}"
ordering = "NONE"
side = "BOTH"`

	neoForgeMod, err := mcmodmeta.NewNeoForgeMod(neoForgeModString)

	assert.Nil(t, err)

	assert.Equal(t, "javafml", neoForgeMod.ModLoader)
	assert.Equal(t, "[1,)", neoForgeMod.LoaderVersion)
	assert.Equal(t, "GPL-3.0", neoForgeMod.License)
	assert.Equal(t, "https://some.issue.tracker", neoForgeMod.IssueTrackerURL)
	assert.Equal(t, 1, len(neoForgeMod.Mods))
	assert.Equal(t, "taterlib", neoForgeMod.Mods[0].ModID)
	assert.Equal(t, "0.1.0", neoForgeMod.Mods[0].Version)
	assert.Equal(t, "TaterLib", neoForgeMod.Mods[0].DisplayName)
	assert.Equal(t, "https://some.update.url", neoForgeMod.Mods[0].UpdateJSONURL)
	assert.Equal(t, "https://some.home.url", neoForgeMod.Mods[0].DisplayURL)
	assert.Equal(t, "TaterLib.png.gz", neoForgeMod.Mods[0].LogoFile)
	assert.Equal(t, "p0t4t0sandwich", neoForgeMod.Mods[0].Credits)
	assert.Equal(t, "p0t4t0sandwich", neoForgeMod.Mods[0].Authors)
	assert.Equal(t, "IGNORE_SERVER_VERSION", neoForgeMod.Mods[0].DisplayTest)
	assert.Equal(t, "some descriptive words", neoForgeMod.Mods[0].Description)
	assert.Equal(t, 6, len(neoForgeMod.Mixins))
	assert.Equal(t, "taterlib.mixins.v1_20.vanilla.json", neoForgeMod.Mixins[0].Config)
	assert.Equal(t, 2, len(neoForgeMod.Dependencies["taterlib"]))
	assert.Equal(t, "neoforge", neoForgeMod.Dependencies["taterlib"][0].ModID)
	assert.Equal(t, "required", neoForgeMod.Dependencies["taterlib"][0].Type)
	assert.Equal(t, "${neo_version_range}", neoForgeMod.Dependencies["taterlib"][0].VersionRange)
	assert.Equal(t, "NONE", neoForgeMod.Dependencies["taterlib"][0].Ordering)
	assert.Equal(t, "BOTH", neoForgeMod.Dependencies["taterlib"][0].Side)
}

func TestSpongeDeserialization(t *testing.T) {
	spongeString := `{
  "loader": {
    "name": "java_plain",
    "version": "1.0"
  },
  "license": "GPL-3.0",
  "plugins": [
    {
      "id": "taterlib",
      "entrypoint": "dev.neuralnexus.taterloader.platforms.Sponge8LoaderPlugin",
      "name": "TaterLib",
      "description": "some more rando descripto words",
      "version": "0.1.0",
      "branding": {},
      "links": {
        "homepage": "https://some.homepage",
        "source": "https://some.repo",
        "issues": "https://some.issue.tracker"
      },
      "dependencies": [
        {
          "id": "spongeapi",
          "version": "8.0.0",
          "load-order": "after",
          "optional": false
        }
      ]
    }
  ]
}`

	spongePlugin, err := mcmodmeta.NewSpongePlugin(spongeString)

	assert.Nil(t, err)

	assert.Equal(t, "java_plain", spongePlugin.Loader.Name)
	assert.Equal(t, "1.0", spongePlugin.Loader.Version)
	assert.Equal(t, "GPL-3.0", spongePlugin.License)
	assert.Equal(t, 1, len(spongePlugin.Plugins))
	assert.Equal(t, "taterlib", spongePlugin.Plugins[0].ID)
	assert.Equal(t, "dev.neuralnexus.taterloader.platforms.Sponge8LoaderPlugin", spongePlugin.Plugins[0].Entrypoint)
	assert.Equal(t, "TaterLib", spongePlugin.Plugins[0].Name)
	assert.Equal(t, "some more rando descripto words", spongePlugin.Plugins[0].Description)
	assert.Equal(t, "0.1.0", spongePlugin.Plugins[0].Version)
	assert.Equal(t, mcmodmeta.SpongeBranding{}, spongePlugin.Plugins[0].Branding)
	assert.Equal(t, "https://some.homepage", spongePlugin.Plugins[0].Links.Homepage)
	assert.Equal(t, "https://some.repo", spongePlugin.Plugins[0].Links.Source)
	assert.Equal(t, "https://some.issue.tracker", spongePlugin.Plugins[0].Links.Issues)
	assert.Equal(t, 1, len(spongePlugin.Plugins[0].Dependencies))
	assert.Equal(t, "spongeapi", spongePlugin.Plugins[0].Dependencies[0].ID)
	assert.Equal(t, "8.0.0", spongePlugin.Plugins[0].Dependencies[0].Version)
	assert.Equal(t, "after", spongePlugin.Plugins[0].Dependencies[0].LoadOrder)
	assert.Equal(t, false, spongePlugin.Plugins[0].Dependencies[0].Optional)
}

func TestVelocityDeserialization(t *testing.T) {
	velocityString := `{
  "id": "taterlib",
  "name": "TaterLib",
  "version": "0.1.0",
  "description": "some words are here",
  "url": "https://some.homepage",
  "authors": [
    "p0t4t0sandwich"
  ],
  "dependencies": [],
  "main": "dev.neuralnexus.taterloader.platforms.VelocityLoaderPlugin"
}`

	velocityPlugin, err := mcmodmeta.NewVelocityPlugin(velocityString)

	assert.Nil(t, err)

	assert.Equal(t, "taterlib", velocityPlugin.ID)
	assert.Equal(t, "TaterLib", velocityPlugin.Name)
	assert.Equal(t, "0.1.0", velocityPlugin.Version)
	assert.Equal(t, "some words are here", velocityPlugin.Description)
	assert.Equal(t, "https://some.homepage", velocityPlugin.URL)
	assert.Equal(t, 1, len(velocityPlugin.Authors))
	assert.Equal(t, "p0t4t0sandwich", velocityPlugin.Authors[0])
	assert.Equal(t, 0, len(velocityPlugin.Dependencies))
	assert.Equal(t, "dev.neuralnexus.taterloader.platforms.VelocityLoaderPlugin", velocityPlugin.Main)
}
