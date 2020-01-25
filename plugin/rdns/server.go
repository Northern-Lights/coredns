package rdns

import "context"

import "time"

type Server struct {
	cache         *Cache
	lookupTimeout time.Duration
}

func NewServer(cache *Cache) *Server {
	return &Server{
		cache:         cache,
		lookupTimeout: 100 * time.Millisecond,
	}
}

func (s *Server) Rdns(ctx context.Context, req *RdnsRequest) (resp *RdnsResponse, err error) {
	resp = new(RdnsResponse)
	ctx, cancel := context.WithTimeout(ctx, s.lookupTimeout)
	defer cancel()

	ch := make(chan []string, 1)
	go func() {
		ch <- []string{s.cache.Lookup(req.IpAddr)}
	}()

	select {
	case domains := <-ch:
		resp.Domains = domains
	case <-ctx.Done():
		err = ctx.Err()
	}

	return
}
