package cast

import (
	"context"

	"github.com/ikasamah/homecast"
)

type NewDevice struct {
	Devices []*homecast.CastDevice
}

func NewHomecast(ctx context.Context) *NewDevice {

	devices := homecast.LookupAndConnect(ctx)
	return &NewDevice{Devices: devices}
}
