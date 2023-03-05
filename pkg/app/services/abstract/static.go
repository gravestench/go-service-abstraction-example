package abstract

import (
	"io/fs"
)

type StaticAssetsService interface {
	fs.ReadDirFS
	fs.ReadFileFS
}
