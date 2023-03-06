package router

import (
	"flag"
)

const (
	flagSlug            = "router-slug"
	flagSlugDefault     = ""
	flagSlugDescription = "root router slug"
)

func (s *Service) Flags() []string {
	return []string{
		flagSlug,
	}
}

func (s *Service) Parse(fs *flag.FlagSet, args []string) error {
	fs.StringVar(&s.flags.slug, flagSlug, flagSlugDefault, flagSlugDescription)

	return fs.Parse(args)
}
