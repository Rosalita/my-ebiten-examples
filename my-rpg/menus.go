package main

import (
	im "github.com/Rosalita/my-ebiten/pkgs/imagemenu"
	lm "github.com/Rosalita/my-ebiten/pkgs/listmenu"
	"github.com/Rosalita/my-ebiten/resources/avatars"
	"github.com/Rosalita/my-ebiten/resources/ui"
)

func initMenus() {

	mainMenuItems := []lm.Item{
		{Name: "playButton",
			Text:     "PLAY",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
		{Name: "optionButton",
			Text:     "OPTIONS",
			TxtX:     16,
			TxtY:     25,
			BgColour: white},
		{Name: "quitButton",
			Text:     "QUIT",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
	}

	mainMenuInput := lm.Input{
		Width:              140,
		Height:             36,
		Tx:                 24,
		Ty:                 24,
		Offy:               40,
		DefaultSelBgColour: pink,
		Items:              mainMenuItems,
	}

	mainMenu, _ = lm.NewMenu(mainMenuInput)

	optionsMenuItems := []lm.Item{
		{Name: "screen",
			Text:     "SCREEN",
			TxtX:     28,
			TxtY:     25,
			BgColour: white},
		{Name: "sound",
			Text:     "SOUND",
			TxtX:     32,
			TxtY:     25,
			BgColour: white},
		{Name: "language",
			Text:     "LANGUAGE",
			TxtX:     4,
			TxtY:     25,
			BgColour: white},
	}

	optionsMenuInput := lm.Input{
		Width:              140,
		Height:             36,
		Tx:                 24,
		Ty:                 24,
		Offy:               40,
		DefaultSelBgColour: pink,
		Items:              optionsMenuItems,
	}

	optionsMenu, _ = lm.NewMenu(optionsMenuInput)

	charGroupItems := []im.Item{
		{
			Name:  "human",
			Bytes: ui.Human_s,
		},
		{
			Name:  "creature",
			Bytes: ui.Creature_s,
		},
	}

	charGroupInput := im.Input{
		Tx:        100,
		Ty:        0,
		ImgWidth:  100,
		ImgHeight: 100,
		Items:     charGroupItems,
	}

	charGroupMenu, _ = im.NewMenu(charGroupInput)

	humanMenuItems := []im.Item{
		{
			Name:  "f1",
			Bytes: avatars.F_01_s,
		},
		{
			Name:  "m1",
			Bytes: avatars.M_01_s,
		},
		{
			Name:  "f2",
			Bytes: avatars.F_02_s,
		},
		{
			Name:  "m2",
			Bytes: avatars.M_02_s,
		},
		{
			Name:  "f3",
			Bytes: avatars.F_03_s,
		},
		{
			Name:  "m3",
			Bytes: avatars.M_03_s,
		},
		{
			Name:  "f4",
			Bytes: avatars.F_04_s,
		},
		{
			Name:  "m4",
			Bytes: avatars.M_04_s,
		},
		{
			Name:  "f5",
			Bytes: avatars.F_05_s,
		},
		{
			Name:  "m5",
			Bytes: avatars.M_05_s,
		},
		{
			Name:  "f6",
			Bytes: avatars.F_06_s,
		},
		{
			Name:  "m6",
			Bytes: avatars.M_06_s,
		},
		{
			Name:  "f7",
			Bytes: avatars.F_07_s,
		},
		{
			Name:  "m7",
			Bytes: avatars.M_07_s,
		},
		{
			Name:  "f8",
			Bytes: avatars.F_08_s,
		},
		{
			Name:  "m8",
			Bytes: avatars.M_08_s,
		},
		{
			Name:  "f9",
			Bytes: avatars.F_09_s,
		},
		{
			Name:  "m9",
			Bytes: avatars.M_09_s,
		},
		{
			Name:  "f10",
			Bytes: avatars.F_10_s,
		},
		{
			Name:  "m10",
			Bytes: avatars.M_10_s,
		},
	}

	humanMenuInput := im.Input{
		Tx:        100,
		Ty:        100,
		ImgWidth:  100,
		ImgHeight: 100,
		Items:     humanMenuItems,
	}

	humanMenu, _ = im.NewMenu(humanMenuInput)

	creatureMenuItems := []im.Item{
		{
			Name:  "c1",
			Bytes: avatars.C_01_s,
		},
		{
			Name:  "c2",
			Bytes: avatars.C_02_s,
		},
		{
			Name:  "c3",
			Bytes: avatars.C_03_s,
		},
		{
			Name:  "c4",
			Bytes: avatars.C_04_s,
		},
		{
			Name:  "c5",
			Bytes: avatars.C_05_s,
		},
		{
			Name:  "c6",
			Bytes: avatars.C_06_s,
		},
		{
			Name:  "c7",
			Bytes: avatars.C_07_s,
		},
		{
			Name:  "c8",
			Bytes: avatars.C_08_s,
		},
		{
			Name:  "c9",
			Bytes: avatars.C_09_s,
		},
		{
			Name:  "c10",
			Bytes: avatars.C_10_s,
		},
		{
			Name:  "c11",
			Bytes: avatars.C_11_s,
		},
		{
			Name:  "c12",
			Bytes: avatars.C_12_s,
		},
		{
			Name:  "c13",
			Bytes: avatars.C_13_s,
		},
		{
			Name:  "c14",
			Bytes: avatars.C_14_s,
		},
		{
			Name:  "c15",
			Bytes: avatars.C_15_s,
		},
		{
			Name:  "c16",
			Bytes: avatars.C_16_s,
		},
		{
			Name:  "c17",
			Bytes: avatars.C_17_s,
		},
		{
			Name:  "c18",
			Bytes: avatars.C_18_s,
		},
		{
			Name:  "c19",
			Bytes: avatars.C_19_s,
		},
		{
			Name:  "c20",
			Bytes: avatars.C_20_s,
		},
	}

	creatureMenuInput := im.Input{
		Tx:        100,
		Ty:        200,
		ImgWidth:  100,
		ImgHeight: 100,
		Items:     creatureMenuItems,
	}

	creatureMenu, _ = im.NewMenu(creatureMenuInput)

}
