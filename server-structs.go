package main

// SiteConfig contains confiurations
// for individual sites
type SiteConfig struct {
	Incomming string // incomming url eg. www.abc.com
	Forward   string // local map eg. localhost:9081
}

// ServerConfig contains  configurations
// which are appied server wide
type ServerConfig struct {
	Sites []SiteConfig // Server will be hosting multiple website
}
