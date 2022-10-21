package rrdb

import (
	"fmt"

	"github.com/rueian/rueidis"
)

const lockName = "rrdb:lock"

func (c *Client) Unlock(key string) error {
	if err := c.client.Do(c.ctx, c.client.B().Spublish().Channel(lockName+":"+key).Message(key).Build()).Error(); err != nil {
		return fmt.Errorf("rrdb.Client.PublishLock: %w", err)
	}
	return nil
}

func (c *Client) Lock(key string) bool {
	if v, ok := c.lockTable.Load(key); ok {
		ch, ok := v.(chan struct{})
		if ok {
			<-ch
			return true
		}
	}

	ch := make(chan struct{}, 1)
	if err := c.client.Receive(c.ctx, c.client.B().Ssubscribe().Channel(lockName+":*").Build(), func(_ rueidis.PubSubMessage) {
	}); err != nil {
		return fmt.Errorf("rrdb.Client.SubscribeLock: %w", err)
	}
	return nil
}
