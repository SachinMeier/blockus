package nostr

import (
	"fmt"
	"github.com/SachinMeier/blockus/pkg/config"
	"github.com/nbd-wtf/go-nostr"
)

type Config struct {
	Relays  []string
	Pubkey  string
	Privkey string
}

func NewConfig(relays []string, privkey string) (*Config, error) {
	pk, err := nostr.GetPublicKey(privkey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse privkey and generate pubkey : %w", err)
	}

	return &Config{
		Relays:  relays,
		Pubkey:  pk,
		Privkey: privkey,
	}, nil
}

const (
	nostrRelaysKey  = "NOSTR_RELAYS"
	nostrPrivkeyKey = "NOSTR_PRIVKEY"
)

func LoadConfig(cfg config.Provider) (*Config, error) {
	relays := cfg.GetStrings(nostrRelaysKey, nil)
	if len(relays) < 1 {
		return nil, fmt.Errorf("nostr relays unset")
	}
	privkey := cfg.GetString(nostrPrivkeyKey, "")
	if privkey == "" {
		return nil, fmt.Errorf("nostr privkey unset")
	}
	return NewConfig(relays, privkey)
}
