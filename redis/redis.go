package redis

import (
	"context"

	RS "github.com/go-redis/redis/v8"
)

var rdb *RS.Client

// Init connects to the database.
func Init(uri, password string, db int) error {

	rdb = RS.NewClient(&RS.Options{
		Addr:     uri,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return nil
}

// Close closes the client, releasing any open resources.
//
// It is rare to Close a Client, as the Client is meant to be
// long-lived and shared between many goroutines.
func Close() error {
	return rdb.Close()
}

// Do creates a Cmd from the args and processes the cmd.
func Do(ctx context.Context, args ...interface{}) *Cmd {
	return rdb.Do(ctx, args...)
}

func Process(ctx context.Context, cmd Cmder) error {
	return rdb.Process(ctx, cmd)
}

// Options returns read-only Options that were used to create the client.
func GetOptions() *Options {
	return rdb.Options()
}

// PoolStats returns connection pool stats.
func PoolStats() *PoolState {
	return rdb.PoolStats()
}

func Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	return rdb.Pipelined(ctx, fn)
}

func Pipeline() Pipeliner {
	return rdb.Pipeline()
}

func TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	return rdb.TxPipelined(ctx, fn)
}

// TxPipeline acts like Pipeline, but wraps queued commands with MULTI/EXEC.
func TxPipeline() Pipeliner {
	return rdb.TxPipeline()
}

// Subscribe subscribes the client to the specified channels.
// Channels can be omitted to create empty subscription.
// Note that this method does not wait on a response from Redis, so the
// subscription may not be active immediately. To force the connection to wait,
// you may call the Receive() method on the returned *PubSub like so:
//
//	sub := client.Subscribe(queryResp)
//	iface, err := sub.Receive()
//	if err != nil {
//	    // handle error
//	}
//
//	// Should be *Subscription, but others are possible if other actions have been
//	// taken on sub since it was created.
//	switch iface.(type) {
//	case *Subscription:
//	    // subscribe succeeded
//	case *Message:
//	    // received first message
//	case *Pong:
//	    // pong received
//	default:
//	    // handle error
//	}
//
//	ch := sub.Channel()
func Subscribe(ctx context.Context, channels ...string) *PubSub {
	return rdb.Subscribe(ctx, channels...)
}

// PSubscribe subscribes the client to the given patterns.
// Patterns can be omitted to create empty subscription.
func PSubscribe(ctx context.Context, channels ...string) *PubSub {
	return rdb.PSubscribe(ctx, channels...)
}
