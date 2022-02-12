package setupkeys

import  "github.com/hajimehoshi/ebiten"


var KeyMap map[ebiten.Key]byte

func SetupKeys() {
  KeyMap = make(map[ebiten.Key]byte)
	KeyMap[ebiten.Key1] = 0x01
	KeyMap[ebiten.Key2] = 0x02
	KeyMap[ebiten.Key3] = 0x03
	KeyMap[ebiten.Key4] = 0x0C
	KeyMap[ebiten.KeyQ] = 0x04
	KeyMap[ebiten.KeyW] = 0x05
	KeyMap[ebiten.KeyE] = 0x06
	KeyMap[ebiten.KeyR] = 0x0D
	KeyMap[ebiten.KeyA] = 0x07
	KeyMap[ebiten.KeyS] = 0x08
	KeyMap[ebiten.KeyD] = 0x09
	KeyMap[ebiten.KeyF] = 0x0E
	KeyMap[ebiten.KeyZ] = 0x0A
	KeyMap[ebiten.KeyX] = 0x00
	KeyMap[ebiten.KeyC] = 0x0B
	KeyMap[ebiten.KeyV] = 0x0F

}
