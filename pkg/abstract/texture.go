package abstract

import rl "github.com/gen2brain/raylib-go/raylib"

type TextureService interface {
	LoadImageToTexture(filepath string) rl.Texture2D
}
