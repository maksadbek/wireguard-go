//go:build !linux

package device

import (
	"github.com/maksadbek/wireguard-go/conn"
	"github.com/maksadbek/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
