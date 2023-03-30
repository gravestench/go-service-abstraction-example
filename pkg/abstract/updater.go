package abstract

import "time"

type Updater interface {
	Service
	Update(time.Duration)
	Ready() bool
}

type UpdateService interface {
	Update()
}
