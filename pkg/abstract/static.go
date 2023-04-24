package abstract

import (
	"io/fs"
)

// StaticAssetsService is an abstraction for
// a service that provides access to the files embedded
// with go:embed
type StaticAssetsService interface {
	fs.ReadDirFS
	fs.ReadFileFS
}
