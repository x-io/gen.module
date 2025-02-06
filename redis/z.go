package redis

import (
	"context"
	"time"

	RS "github.com/go-redis/redis/v8"
)

//------------------------------------------------------------------------------

// Nil reply returned by Redis when key does not exist.
const Nil = RS.Nil

// Z represents sorted set member.
type Z = RS.Z

// ZWithKey represents sorted set member including the name of the key where it was popped.
type ZWithKey = RS.ZWithKey

// ZStore is used as an arg to ZInterStore and ZUnionStore.
type ZStore = RS.ZStore

//BZPopMax Redis `BZPOPMAX key [key ...] timeout` command.
func BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	return rdb.BZPopMax(ctx, timeout, keys...)
}

//BZPopMin Redis `BZPOPMIN key [key ...] timeout` command.
func BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	return rdb.BZPopMin(ctx, timeout, keys...)
}

//ZAdd Redis `ZADD key score member [score member ...]` command.
func ZAdd(ctx context.Context, key string, members ...*Z) *IntCmd {
	return rdb.ZAdd(ctx, key, members...)
}

//ZAddNX Redis `ZADD key NX score member [score member ...]` command.
func ZAddNX(ctx context.Context, key string, members ...*Z) *IntCmd {
	return rdb.ZAddNX(ctx, key, members...)
}

//ZAddXX Redis `ZADD key XX score member [score member ...]` command.
func ZAddXX(ctx context.Context, key string, members ...*Z) *IntCmd {
	return rdb.ZAddXX(ctx, key, members...)
}

//ZAddCh Redis `ZADD key CH score member [score member ...]` command.
func ZAddCh(ctx context.Context, key string, members ...*Z) *IntCmd {
	return rdb.ZAddCh(ctx, key, members...)
}

//ZAddNXCh Redis `ZADD key NX CH score member [score member ...]` command.
func ZAddNXCh(ctx context.Context, key string, members ...*Z) *IntCmd {
	return rdb.ZAddNXCh(ctx, key, members...)
}

//ZAddXXCh Redis `ZADD key XX CH score member [score member ...]` command.
func ZAddXXCh(ctx context.Context, key string, members ...*Z) *IntCmd {
	return rdb.ZAddXXCh(ctx, key, members...)
}

//ZIncr Redis `ZADD key INCR score member` command.
func ZIncr(ctx context.Context, key string, member *Z) *FloatCmd {
	return rdb.ZIncr(ctx, key, member)
}

//ZIncrNX Redis `ZADD key NX INCR score member` command.
func ZIncrNX(ctx context.Context, key string, member *Z) *FloatCmd {
	return rdb.ZIncrNX(ctx, key, member)
}

//ZIncrXX Redis `ZADD key XX INCR score member` command.
func ZIncrXX(ctx context.Context, key string, member *Z) *FloatCmd {
	return rdb.ZIncrXX(ctx, key, member)
}

//ZCard ZCard
func ZCard(ctx context.Context, key string) *IntCmd {
	return rdb.ZCard(ctx, key)
}

//ZCount ZCount
func ZCount(ctx context.Context, key, min, max string) *IntCmd {
	return rdb.ZCount(ctx, key, min, max)
}

//ZLexCount ZLexCount
func ZLexCount(ctx context.Context, key, min, max string) *IntCmd {
	return rdb.ZLexCount(ctx, key, min, max)
}

//ZIncrBy ZIncrBy
func ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd {
	return rdb.ZIncrBy(ctx, key, increment, member)
}

//ZInterStore ZInterStore
func ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd {
	return rdb.ZInterStore(ctx, destination, store)
}

//ZPopMax ZPopMax
func ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	return rdb.ZPopMax(ctx, key, count...)
}

//ZPopMin ZPopMin
func ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	return rdb.ZPopMin(ctx, key, count...)
}

//ZRange ZRange
func ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	return rdb.ZRange(ctx, key, start, stop)
}

//ZRangeWithScores ZRangeWithScores
func ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	return rdb.ZRangeWithScores(ctx, key, start, stop)
}

//ZRangeBy ZRangeBy
type ZRangeBy = RS.ZRangeBy

//ZRangeByScore ZRangeByScore
func ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	return rdb.ZRangeByScore(ctx, key, opt)
}

//ZRangeByLex ZRangeByLex
func ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	return rdb.ZRangeByLex(ctx, key, opt)
}

//ZRangeByScoreWithScores ZRangeByScoreWithScores
func ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {
	return rdb.ZRangeByScoreWithScores(ctx, key, opt)
}

//ZRank ZRank
func ZRank(ctx context.Context, key, member string) *IntCmd {
	return rdb.ZRank(ctx, key, member)
}

//ZRem ZRem
func ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	return rdb.ZRem(ctx, key, members...)
}

//ZRemRangeByRank ZRemRangeByRank
func ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd {
	return rdb.ZRemRangeByRank(ctx, key, start, stop)
}

//ZRemRangeByScore ZRemRangeByScore
func ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd {
	return rdb.ZRemRangeByScore(ctx, key, min, max)
}

//ZRemRangeByLex ZRemRangeByLex
func ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd {
	return rdb.ZRemRangeByLex(ctx, key, min, max)
}

//ZRevRange ZRevRange
func ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	return rdb.ZRevRange(ctx, key, start, stop)
}

//ZRevRangeWithScores ZRevRangeWithScores
func ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	return rdb.ZRevRangeWithScores(ctx, key, start, stop)
}

//ZRevRangeByScore ZRevRangeByScore
func ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	return rdb.ZRevRangeByScore(ctx, key, opt)
}

//ZRevRangeByLex ZRevRangeByLex
func ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	return rdb.ZRevRangeByLex(ctx, key, opt)
}

//ZRevRangeByScoreWithScores ZRevRangeByScoreWithScores
func ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {
	return rdb.ZRevRangeByScoreWithScores(ctx, key, opt)
}

//ZRevRank ZRevRank
func ZRevRank(ctx context.Context, key, member string) *IntCmd {
	return rdb.ZRevRank(ctx, key, member)
}

//ZScore ZScore
func ZScore(ctx context.Context, key, member string) *FloatCmd {
	return rdb.ZScore(ctx, key, member)
}

//ZUnionStore ZUnionStore
func ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd {
	return rdb.ZUnionStore(ctx, dest, store)
}
