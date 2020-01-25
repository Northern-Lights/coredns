package rdns

import (
	"net"
	"os"

	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"google.golang.org/grpc"
)

var null *Rdns = nil

func init() { plugin.Register(name, setup) }

func setup(c *caddy.Controller) error {
	if !c.Next() || !c.Next() {
		return c.ArgErr()
	}
	sockLocation := c.Val()

	// cache is shared by the plugin and the service
	const defaultCacheSize = 1024
	cache := &Cache{
		mapping: make(map[string]string, defaultCacheSize),
	}

	svc := NewServer(cache)
	var svr *grpc.Server

	var e *chan error
	c.OnStartup(func() error {
		// remove any previous socket
		os.Remove(sockLocation)

		// serving the listener blocks, so get any error over a channel
		ech := make(chan error)
		e = &ech

		// do it here because svr.Serve() cannot be called twice
		svr = grpc.NewServer()
		RegisterRdnsCacheServer(svr, svc)

		lis, err := net.Listen("unix", sockLocation)
		if err != nil {
			return plugin.Error(name, err)
		}
		const userWritePermissions = 0766
		err = os.Chmod(sockLocation, userWritePermissions)
		if err != nil {
			return plugin.Error(name, err)
		}

		go runExternalService(svr, lis, e)

		return nil
	})

	c.OnShutdown(func() error {
		err := os.Remove(sockLocation)
		svr.GracefulStop()
		if err != nil {
			return err
		} else if err := <-*e; err != nil {
			return err
		}
		return nil
	})

	// Add the Plugin to CoreDNS, so Servers can use it in their plugin chain.
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		rdns := New(cache)
		rdns.next = next
		return rdns
	})

	return nil
}

func runExternalService(svr *grpc.Server, lis net.Listener, e *chan error) {
	*e <- svr.Serve(lis)
}
