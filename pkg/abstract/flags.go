package abstract

import (
	"flag"
)

// ServiceThatUsesFlags is what a service should implement if it uses CLI flags.
// This is intended to be picked up by an additional service which knows how
// to pass the right flags into this service Parse method
type ServiceThatUsesFlags interface {
	Service
	// Flags yields the flags that the service needs, to filter out
	// the unnecessary flags from the global os.Args
	Flags() []string

	// Parse the given args, which should be the filtered set of flags
	// that only this service needs
	Parse(flagSet *flag.FlagSet, args []string) error
}
