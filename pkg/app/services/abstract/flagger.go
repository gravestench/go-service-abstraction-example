package abstract

import (
	"flag"
)

// FlagService uses command line flags
type FlagService interface {
	Service
	// Flags yields the flags that the service needs, to filter out
	// the unnecessary flags from the global os.Args
	Flags() []string

	// Parse the given args, which should be the filtered set of flags
	// that only this service needs
	Parse(flagSet *flag.FlagSet, args []string) error
}
