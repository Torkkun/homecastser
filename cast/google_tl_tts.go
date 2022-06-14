package cast

import (
	"context"
	"log"

	"github.com/ikasamah/homecast"
)

type NewDevice struct {
	Devices []*homecast.CastDevice
}

func NewHomecast(ctx context.Context) *NewDevice {
	devices := homecast.LookupAndConnect(ctx)
	if devices == nil {
		log.Fatalln("not connect devices")
	}
	return &NewDevice{Devices: devices}
}
