package main

import (
	"fmt"
	"github.com/SachinMeier/blockus/pkg/log"
	"github.com/nbd-wtf/go-nostr"
	"time"
)

func main() {
	sk := nostr.GeneratePrivateKey()
	pk, err := nostr.GetPublicKey(sk)
	if err != nil {
		log.Fatalf("failed to derive pubkey: %v", err)
	}

	ev := nostr.Event{
		PubKey:    pk,
		CreatedAt: time.Now(),
		Kind:      nostr.KindTextNote,
		Tags:      nil,
		Content:   "It might make sense just to get some in case it catches on.",
	}
	err = ev.Sign(sk)
	if err != nil {
		log.Fatalf("failed to sign event: %v", err)
	}

	fmt.Printf("sk: %s\n", sk)
	fmt.Printf("pk: %s\n", pk)
	fmt.Printf("ev: {\n")
	fmt.Printf("  id: \"%s\",\n", ev.ID)
	fmt.Printf("  pubkey: \"%s\",\n", ev.PubKey)
	fmt.Printf("  created_at: %d,\n", ev.CreatedAt.Unix())
	fmt.Printf("  kind: %d,\n", ev.Kind)
	fmt.Printf("  tags: %v,\n", ev.Tags)
	fmt.Printf("  content: \"%s\",\n", ev.Content)
	fmt.Printf("  sig: \"%s\"\n", ev.Sig)
	fmt.Printf("}\n")
}
