package main

import (
	"blockus/pkg/bitcoind"
	"blockus/pkg/config"
	"blockus/pkg/log"
	"context"
	"sync"
)

func main() {
	ctx := context.Background()
	cfg := config.NewEnvProvider()

	btcCfg := config.LoadBitcoindConfig(cfg)
	clientFactory := bitcoind.NewClientFactory(btcCfg)

	client, err := clientFactory.NewClient()
	if err != nil {
		log.Fatalf("failed to create bitcoind client : %v", err)
	}

	notifs, err := client.NewZMQSubscription(bitcoind.TopicSequence)
	if err != nil {
		log.Fatalf("failed to subscribe to zmq blocks : %v", err)
	}

	handleZMQMsg := func(topic string, msg bitcoind.ZMQMsg) error {
		log.Infof("recv block msg: %s:%s", topic, msg.Data())
		return nil
	}
	handleErr := func(err error) {
		log.Errorf("zmq err : %v", err)
	}
	log.Infof("listening to ZMQ notifications at %s", btcCfg.ZMQHost)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go bitcoind.ListenForZMQNotifications(ctx, notifs, handleZMQMsg, handleErr)
	wg.Wait()
}
