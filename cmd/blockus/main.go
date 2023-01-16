package main

import (
	"context"
	"github.com/SachinMeier/blockus/pkg/bitcoind"
	"github.com/SachinMeier/blockus/pkg/config"
	"github.com/SachinMeier/blockus/pkg/log"
	"github.com/SachinMeier/blockus/pkg/nostr"
	zmq "github.com/pebbe/zmq4"
	"sync"
)

func main() {
	ctx := context.Background()
	env := config.NewEnvProvider()

	// setup bitcoind
	var notifs *zmq.Socket
	{
		btcCfg := bitcoind.LoadConfig(env)
		clientFactory := bitcoind.NewClientFactory(btcCfg)

		btcClient, err := clientFactory.NewClient()
		if err != nil {
			log.Fatalf("failed to create bitcoind btcClient : %v", err)
		}

		notifs, err = btcClient.NewZMQSubscription(bitcoind.TopicSequence)
		if err != nil {
			log.Fatalf("failed to subscribe to zmq blocks : %v", err)
		}
		log.Infof("listening to ZMQ notifications at %s", btcCfg.ZMQHost)
	}

	// setup nostr
	var nostrClient *nostr.Client
	{
		nostrCfg, err := nostr.LoadConfig(env)
		if err != nil {
			log.Fatalf("failed to load nostr config : %v", err)
		}

		nostrClient = nostr.NewClient(ctx, *nostrCfg)
	}

	handleZMQMsg := func(topic string, msg bitcoind.ZMQMsg) error {
		log.Infof("recv block msg: %s:%s", topic, msg.Data())

		event, err := nostrClient.NewBlockEvent(msg.Data())
		if err != nil {
			return err
		}
		return nostrClient.PublishEvent(*event)
	}
	handleErr := func(err error) {
		log.Errorf("zmq err : %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go bitcoind.ListenForZMQNotifications(ctx, notifs, handleZMQMsg, handleErr)
	wg.Wait()
}
