// Package rdns provides reverse DNS support, mapping hostnames to IP addresses.
// It also provides a service on the side to retrieve hostnames from IP
// addresses
package rdns

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

const name = "rdns"

// Rdns is a reverse DNS cache
type Rdns struct {
	next  plugin.Handler
	Zones []string

	cache *Cache
}

// New returns a new rDNS plugin
func New(cache *Cache) *Rdns {
	return &Rdns{
		cache: cache,
	}
}

// ServeDNS passes the request to the next plugin and stores the IP-to-domain
// mapping of any response
func (r *Rdns) ServeDNS(ctx context.Context, w dns.ResponseWriter, req *dns.Msg) (int, error) {
	w = &responseWriter{
		ResponseWriter: w,
		cache:          r.cache,
	}
	return plugin.NextOrFailure(r.next.Name(), r.next, ctx, w, req)
}

// Name returns the name of the plugin
func (*Rdns) Name() string {
	return name
}

type responseWriter struct {
	dns.ResponseWriter
	cache *Cache
}

func (w *responseWriter) WriteMsg(msg *dns.Msg) error {
	for _, rr := range msg.Answer {
		switch rr := rr.(type) {
		case *dns.A:
			w.cache.Add(rr.A.String(), rr.Header().Name)
		case *dns.AAAA:
			w.cache.Add(rr.AAAA.String(), rr.Header().Name)
		default:
			// TODO: log?
		}
	}
	return w.ResponseWriter.WriteMsg(msg)
}
