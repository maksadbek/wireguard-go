//go:build !linux

package device

import (
	"github.com/maksadbek/wireguard-go/wireguard/conn"
	"github.com/maksadbek/wireguard-go/wireguard/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
