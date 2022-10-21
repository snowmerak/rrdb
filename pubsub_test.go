package rrdb_test

import (
	"context"
	"testing"

	"github.com/snowmerak/rrdb"
)

const TestRedisAddr = "127.0.0.1:6379"
const TestUsername = ""
const TestPassword = ""

func TestPubSubLock(t *testing.T) {
	client, err := rrdb.NewClient(context.Background(), TestUsername, TestPassword, TestRedisAddr)
	if err != nil {
		t.Error(err)
	}

}
