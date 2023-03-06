package abstract

import (
	"flag"
)

// FlagService is what a service should implement if it uses CLI flags.
// This is inteded to be picked up by an additional service which knows how
// to filter the flags in `os.Args` to be just the ones that the target service needs.
type FlagService interface {
	Service
	// Flags yields the flags that the service needs, to filter out
	// the unnecessary flags from the global os.Args
	Flags() []string

	// Parse the given args, which should be the filtered set of flags
	// that only this service needs
	Parse(flagSet *flag.FlagSet, args []string) error
}
