package bitcoind

import "fmt"

type ClientFactory struct {
	cfg Config
}

func NewClientFactory(cfg Config) *ClientFactory {
	return &ClientFactory{
		cfg: cfg,
	}
}

func (f *ClientFactory) NewClient() (*Client, error) {
	return NewClient(f.cfg)
}

func (f *ClientFactory) Host() string {
	return fmt.Sprintf("%s:%s", f.cfg.Host, f.cfg.Port)
}
