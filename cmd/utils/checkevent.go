package main

import (
	"fmt"
	"time"

	"github.com/nbd-wtf/go-nostr"
)

func main() {
	e := nostr.Event{
		ID:        "bfe9a9f57c61d04493c74189dc6ffa73f54522d0bf470b6194dcf33c323eb3c0",
		PubKey:    "f7b642493ed5a462c2faebb5d6cd208e42b8ea36f0b479f2aa2f5f5bf6f67aa7",
		CreatedAt: time.Unix(1673931680, 0),
		Kind:      1,
		Tags:      nostr.Tags{},
		Content:   "running bitcoin",
		Sig:       "6815aaa87665e68a43639631d9dbac7c3f6b825ff14e6689a100eef2d307f54701ec6b28fa9ab9c9ea01215f07db4e507d8e06a0e01527035f0b0ae9e4c2779c",
	}
	//e2 := nostr.Event{
	//	ID:        "e9a415c41148b12c31aabba3a41b9cd1cfe713408305cc16f8c5ea6e0f99356b",
	//	PubKey:    "b0448252cddc47798e5e726b5c6de25f3c486a01427d736915071f6d320abaab",
	//	CreatedAt: time.Unix(1673931868, 0),
	//	Kind:      1,
	//	Tags: nostr.Tags{
	//		{"p", "f7b642493ed5a462c2faebb5d6cd208e42b8ea36f0b479f2aa2f5f5bf6f67aa7", "random"},
	//		{"e", "bfe9a9f57c61d04493c74189dc6ffa73f54522d0bf470b6194dcf33c323eb3c0", "other"},
	//	},
	//	Content: "It might make sense just to get some in case it catches on.",
	//	Sig:     "f0f6d65f6a2257b3af1a8e1760a697ed1dd2ad8839e3fb708e33555ad1d01206d7a68f9a944387c6505a7fd8cbfac29fc8151bf9109e62b5efbb1479893a8bca",
	//}
	fmt.Printf("%v\n", e.GetID())
	res, _ := e.CheckSignature()
	fmt.Printf("%v\n", res)
}
