package settings

import (
	"github.com/wieku/danser-go/framework/env"
	"os"
	"path/filepath"
	"runtime"
)

var General = initGeneral()

func initGeneral() *general {
	osuBaseDir := ""
	if runtime.GOOS == "windows" {
		osuBaseDir = filepath.Join(os.Getenv("localappdata"), "osu!")
	} else {
		dir, _ := os.UserHomeDir()
		osuBaseDir = filepath.Join(dir, ".osu")
	}

	return &general{
		OsuSongsDir:       filepath.Join(osuBaseDir, "Songs"),
		OsuSkinsDir:       filepath.Join(osuBaseDir, "Skins"),
		DiscordPresenceOn: true,
		UnpackOszFiles:    true,
	}
}

type general struct {

	// Directory that contains osu! songs,
	OsuSongsDir string

	// Directory that contains osu! skins,
	OsuSkinsDir string

	// Whether discord should show that danser is on
	DiscordPresenceOn bool

	// Whether danser should unpack .osz files in Songs folder, osu! may complain about it
	UnpackOszFiles bool

	songsDir *string
	skinsDir *string
}

func (g *general) GetSongsDir() string {
	if g.songsDir == nil {
		dir := filepath.Join(env.DataDir(), g.OsuSongsDir)

		if filepath.IsAbs(g.OsuSongsDir) {
			dir = g.OsuSongsDir
		}

		g.songsDir = &dir
	}

	return *g.songsDir
}

func (g *general) GetSkinsDir() string {
	if g.skinsDir == nil {
		dir := filepath.Join(env.DataDir(), g.OsuSkinsDir)

		if filepath.IsAbs(g.OsuSkinsDir) {
			dir = g.OsuSkinsDir
		}

		g.skinsDir = &dir
	}

	return *g.skinsDir
}
