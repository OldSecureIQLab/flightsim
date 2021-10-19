package simulator

import (
	"context"
	"net"

	"github.com/pkg/errors"
)

// Hijack simulator.
type Hijack struct{}

// NewHijack creates port scan simulator.
func NewHijack() *Hijack {
	return &Hijack{}
}

// Simulate port scanning for given host.
func (*Hijack) Simulate(ctx context.Context, extIP net.IP, host string) error {
	d := &net.Dialer{
		LocalAddr: &net.UDPAddr{IP: extIP},
	}
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return d.DialContext(ctx, "udp", "dns.sandbox-services.alphasoc.xyz:53")
		},
	}

	addrs, err := r.LookupHost(ctx, host)
	if err != nil {
		return err
	}
	if len(addrs) == 0 {
		return errors.New("No DNS response")
	}

	return nil
}

// Hosts returns one domain to simulate dns query.
func (s *Hijack) Hosts(scope string, n int) ([]string, error) {
	return []string{"alphasoc.com"}, nil
}
