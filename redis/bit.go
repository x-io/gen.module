package redis

import (
	"context"

	RS "github.com/go-redis/redis/v8"
)

//GetBit GetBit
func GetBit(ctx context.Context, key string, offset int64) *IntCmd {
	return rdb.GetBit(ctx, key, offset)
}

//SetBit SetBit
func SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd {
	return rdb.SetBit(ctx, key, offset, value)
}

//BitCount BitCount
func BitCount(ctx context.Context, key string, start, end int64) *IntCmd {
	return rdb.BitCount(ctx, key, &RS.BitCount{start, end})
}

//BitOpAnd BitOpAnd
func BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd {
	return rdb.BitOpAnd(ctx, destKey, keys...)
}

//BitOpOr BitOpOr
func BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd {
	return rdb.BitOpOr(ctx, destKey, keys...)
}

//BitOpXor BitOpXor
func BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd {
	return rdb.BitOpXor(ctx, destKey, keys...)
}

//BitOpNot BitOpNot
func BitOpNot(ctx context.Context, destKey string, key string) *IntCmd {
	return rdb.BitOpNot(ctx, destKey, key)
}

//BitPos BitPos
func BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd {
	return rdb.BitPos(ctx, key, bit, pos...)
}

//BitField BitField
func BitField(ctx context.Context, key string, args ...interface{}) *IntSliceCmd {
	return rdb.BitField(ctx, key, args...)
}
