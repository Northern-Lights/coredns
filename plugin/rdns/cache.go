package rdns

import "sync"

// Cache caches IP-to-domain mappings
type Cache struct {
	sync.RWMutex
	mapping map[string]string
}

// Lookup returns the domain mapping for the given IP address
func (c *Cache) Lookup(ip string) (domain string) {
	c.RLock()
	defer c.RUnlock()
	return c.mapping[ip]
}

// Add adds an IP-to-domain mapping, overwriting anything previous for that
// domain
func (c *Cache) Add(ip, domain string) {
	c.Lock()
	defer c.Unlock()
	c.mapping[ip] = domain
}
