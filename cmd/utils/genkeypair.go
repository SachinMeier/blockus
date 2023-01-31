package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/SachinMeier/blockus/pkg/log"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {
	ctx := context.Background()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < 100; i++ {
		// cannot contain b, i, o, 1
		go findPrefixedNpub(ctx, wg, "halfin")
	}
	wg.Wait()
}

func findPrefixedNpub(ctx context.Context, wg *sync.WaitGroup, prefix string) {
	var sk, pk, nsec, npub string
	pfxlen := len(prefix)
	log.Infof("starting to search for npub with prefix: %s", prefix)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			sk = nostr.GeneratePrivateKey()
			pk, _ = nostr.GetPublicKey(sk)
			npub, _ = nip19.EncodePublicKey(pk)
			if npub[5:(5+pfxlen)] == prefix {
				nsec, _ = nip19.EncodePrivateKey(sk)

				fmt.Println("sk:", sk)
				fmt.Println("pk:", pk)
				fmt.Println("nsec: ", nsec)
				fmt.Println("npub: ", npub)
				wg.Done()
			}
		}
	}
}
