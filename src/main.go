package main

import (
	"chip8/cpu"
	"flag"
	"os"

	"image/color"

  setupkeys "chip8/setupkeys"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var chip8 cpu.Chip8
var audioPlayer *audio.Player
var square *ebiten.Image

func init() {
	square, _ = ebiten.NewImage(10, 10, ebiten.FilterNearest)
	square.Fill(color.White)
}

func getInput() bool {
	for key, value := range setupkeys.KeyMap {
		if ebiten.IsKeyPressed(key) {
			chip8.Keys[value] = 0x01
			return true
		}
	}
	return false
}

func update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	for i := 0; i < 10; i++ {
		chip8.Draw = false
		chip8.Inputflag = false
		gotInput := true
		chip8.Run()

		if chip8.Inputflag {
			gotInput = getInput()
			if !gotInput {
				chip8.Pc = chip8.Pc - 2
			}
		}

		if chip8.Draw || !gotInput {
			for i := 0; i < 32; i++ {
				for j := 0; j < 64; j++ {
					if chip8.Display[i][j] == 0x01 {

						opts := &ebiten.DrawImageOptions{}

						opts.GeoM.Translate(float64(j*10), float64(i*10))

						screen.DrawImage(square, opts)
					}
				}
			}
		}

		for key, value := range setupkeys.KeyMap {
			if ebiten.IsKeyPressed(key) {
				chip8.Keys[value] = 0x01
			} else {
				chip8.Keys[value] = 0x00
			}
		}

		if chip8.SoundTimer > 0 {
			audioPlayer.Play()
			audioPlayer.Rewind()
		}

	}

	return nil
}

func start() {
  rom := flag.String("rom", "", "")
  flag.Parse()

  if *rom == "" {
    os.Exit(1)
  } else {
    chip8 = cpu.NewCpu()
	  chip8.LoadProgram(*rom)

	  if err := ebiten.Run(update, 640, 320, 1, "CHIP-8"); err != nil {
		panic(err)
  }
}
}

func main() {
	audioContext, _ := audio.NewContext(48000)
	f, _ := ebitenutil.OpenFile("assets/beep.mp3")
	d, _ := mp3.Decode(audioContext, f)
	audioPlayer, _ = audio.NewPlayer(audioContext, d)
	setupkeys.SetupKeys()
  start()
	}

