package customdialroutingextension

import (
	"context"
	"fmt"
	"net"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/extensionmiddleware"
)

type customDialer struct {
	cfg *Config
}

func (c *customDialer) GetDialContext(_ context.Context) (extensionmiddleware.DialContextFunc, error) {
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		iface, err := net.InterfaceByName(c.cfg.NetworkInterface)
		if err != nil {
			return nil, fmt.Errorf("interface lookup failed: %w", err)
		}

		addrs, err := iface.Addrs()
		if err != nil || len(addrs) == 0 {
			return nil, fmt.Errorf("no addresses found on interface %s: %w", c.cfg.NetworkInterface, err)
		}

		// Find the first usable IPv4 address on the interface
		var localIP net.IP
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					localIP = ipNet.IP
					break
				}
			}
		}
		if localIP == nil {
			return nil, fmt.Errorf("no IPv4 address found on %s", c.cfg.NetworkInterface)
		}

		dialer := &net.Dialer{
			LocalAddr: &net.TCPAddr{
				IP: localIP,
			},
		}
		return dialer.DialContext(ctx, network, address)
	}, nil
}

func (c *customDialer) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (c *customDialer) Shutdown(_ context.Context) error {
	return nil
}
