package redis

import (
	"context"
)

//------------------------------------------------------------------------------

// Publish posts the message to the channel.
func Publish(ctx context.Context, channel string, message interface{}) *IntCmd {
	return rdb.Publish(ctx, channel, message)
}

//PubSubChannels PubSubChannels
func PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd {
	return rdb.PubSubChannels(ctx, pattern)
}

//PubSubNumSub PubSubNumSub
func PubSubNumSub(ctx context.Context, channels ...string) *StringIntMapCmd {
	return rdb.PubSubNumSub(ctx, channels...)
}

//PubSubNumPat PubSubNumPat
func PubSubNumPat(ctx context.Context) *IntCmd {
	return rdb.PubSubNumPat(ctx)
}
