package nostr

import (
	"context"
	"fmt"
	"github.com/SachinMeier/blockus/pkg/log"
	"github.com/nbd-wtf/go-nostr"
	"time"
)

type Client struct {
	ctx context.Context
	cfg Config
}

func NewClient(ctx context.Context, cfg Config) *Client {
	return &Client{
		ctx: ctx,
		cfg: cfg,
	}
}

func (client *Client) NewBlockEvent(text string) (*nostr.Event, error) {
	event := nostr.Event{
		PubKey:    client.cfg.Pubkey,
		CreatedAt: time.Now(),
		Kind:      nostr.KindTextNote,
		Tags:      nil,
		Content:   text,
	}
	err := event.Sign(client.cfg.Privkey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign block event : %w", err)
	}
	return &event, nil
}

func (client *Client) PublishEvent(event nostr.Event) error {
	published := false
	for _, url := range client.cfg.Relays {
		relay, err := nostr.RelayConnect(client.ctx, url)
		if err != nil {
			log.Warnf("failed to connect to relay %s : %w", url, err)
			continue
		}
		relay.Publish(client.ctx, event)
		published = true
	}
	if !published {
		return fmt.Errorf("failed to publish event to any relays : %s", event.ID)
	}
	return nil
}
