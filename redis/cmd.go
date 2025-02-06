package redis

import (
	"context"
	"time"
)

// Get `GET key` command. It returns redis.Nil error when key does not exist.
func Get(ctx context.Context, key string) *StringCmd {
	return rdb.Get(ctx, key)
}

func Exists(ctx context.Context, keys ...string) *IntCmd {
	return rdb.Exists(ctx, keys...)
}

func Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	return rdb.Expire(ctx, key, expiration)
}

func ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd {
	return rdb.ExpireAt(ctx, key, tm)
}

func Incr(ctx context.Context, key string) *IntCmd {
	return rdb.Incr(ctx, key)
}

func IncrBy(ctx context.Context, key string, value int64) *IntCmd {
	return rdb.IncrBy(ctx, key, value)
}

// MSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet([]string{"key1", "value1", "key2", "value2"})
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func MSet(ctx context.Context, values ...interface{}) *StatusCmd {
	return rdb.MSet(ctx, values)
}

// MSetNX is like SetNX but accepts multiple values:
//   - MSetNX("key1", "value1", "key2", "value2")
//   - MSetNX([]string{"key1", "value1", "key2", "value2"})
//   - MSetNX(map[string]interface{}{"key1": "value1", "key2": "value2"})
func MSetNX(ctx context.Context, values ...interface{}) *BoolCmd {
	return rdb.MSetNX(ctx, values)
}

// Set Redis `SET key value [expiration]` command.
// Use expiration for `SETEX`-like behavior.
//
// Zero expiration means the key has no expiration time.
// KeepTTL(-1) expiration is a Redis KEEPTTL option to keep existing TTL.
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd {
	return rdb.Set(ctx, key, value, expiration)
}

// SetNX Redis `SET key value [expiration] NX` command.
//
// Zero expiration means the key has no expiration time.
// KeepTTL(-1) expiration is a Redis KEEPTTL option to keep existing TTL.
func SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd {
	return rdb.SetNX(ctx, key, value, expiration)
}

// SetXX Redis `SET key value [expiration] XX` command.
//
// Zero expiration means the key has no expiration time.
// KeepTTL(-1) expiration is a Redis KEEPTTL option to keep existing TTL.
func SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd {
	return rdb.SetXX(ctx, key, value, expiration)
}

// SetRange SetRange
func SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd {
	return rdb.SetRange(ctx, key, offset, value)
}

// StrLen StrLen
func StrLen(ctx context.Context, key string) *IntCmd {
	return rdb.StrLen(ctx, key)
}

func Append(ctx context.Context, key, value string) *IntCmd {
	return rdb.Append(ctx, key, value)
}

//------------------------------------------------------------------------------

// Scan Scan
func Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd {
	return rdb.Scan(ctx, cursor, match, count)
}

// SScan SScan
func SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	return rdb.SScan(ctx, key, cursor, match, count)
}

// HScan HScan
func HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	return rdb.HScan(ctx, key, cursor, match, count)
}

// ZScan ZScan
func ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	return rdb.ZScan(ctx, key, cursor, match, count)
}

//------------------------------------------------------------------------------

// HDel HDel
func HDel(ctx context.Context, key string, fields ...string) *IntCmd {
	return rdb.HDel(ctx, key, fields...)
}

// HExists HExists
func HExists(ctx context.Context, key, field string) *BoolCmd {
	return rdb.HExists(ctx, key, field)
}

// HGet HGet
func HGet(ctx context.Context, key, field string) *StringCmd {
	return rdb.HGet(ctx, key, field)
}

// HGetAll HGetAll
func HGetAll(ctx context.Context, key string) *StringStringMapCmd {
	return rdb.HGetAll(ctx, key)
}

// HIncrBy HIncrBy
func HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd {
	return rdb.HIncrBy(ctx, key, field, incr)
}

// HIncrByFloat HIncrByFloat
func HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd {
	return rdb.HIncrByFloat(ctx, key, field, incr)
}

// HKeys HKeys
func HKeys(ctx context.Context, key string) *StringSliceCmd {
	return rdb.HKeys(ctx, key)
}

// HLen HLen
func HLen(ctx context.Context, key string) *IntCmd {
	return rdb.StrLen(ctx, key)
}

// HMGet returns the values for the specified fields in the hash stored at key.
// It returns an interface{} to distinguish between empty string and nil value.
func HMGet(ctx context.Context, key string, fields ...string) *SliceCmd {
	return rdb.HMGet(ctx, key, fields...)
}

// HSet accepts values in following formats:
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
//
// Note that it requires Redis v4 for multiple field/value pairs support.
func HSet(ctx context.Context, key string, values ...interface{}) *IntCmd {
	return rdb.HSet(ctx, key, values...)
}

// HMSet is a deprecated version of HSet left for compatibility with Redis 3.
func HMSet(ctx context.Context, key string, values ...interface{}) *BoolCmd {
	return rdb.HMSet(ctx, key, values...)
}

// HSetNX HSetNX
func HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd {
	return rdb.HSetNX(ctx, key, field, value)
}

// HVals HVals
func HVals(ctx context.Context, key string) *StringSliceCmd {
	return rdb.HVals(ctx, key)
}

//------------------------------------------------------------------------------

// BLPop BLPop
func BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	return rdb.BLPop(ctx, timeout, keys...)
}

// BRPop BRPop
func BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	return rdb.BRPop(ctx, timeout, keys...)
}

// BRPopLPush BRPopLPush
func BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd {
	return rdb.BRPopLPush(ctx, source, destination, timeout)
}

// LIndex LIndex
func LIndex(ctx context.Context, key string, index int64) *StringCmd {
	return rdb.LIndex(ctx, key, index)
}

// LInsert LInsert
func LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd {
	return rdb.LInsert(ctx, key, op, pivot, value)
}

// LInsertBefore LInsertBefore
func LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *IntCmd {
	return rdb.LInsertBefore(ctx, key, pivot, value)
}

// LInsertAfter LInsertAfter
func LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *IntCmd {
	return rdb.LInsertAfter(ctx, key, pivot, value)
}

// LLen LLen
func LLen(ctx context.Context, key string) *IntCmd {
	return rdb.LLen(ctx, key)
}

// LPop LPop
func LPop(ctx context.Context, key string) *StringCmd {
	return rdb.LPop(ctx, key)
}

// LPush LPush
func LPush(ctx context.Context, key string, values ...interface{}) *IntCmd {
	return rdb.LPush(ctx, key, values...)
}

// LPushX LPushX
func LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd {
	return rdb.LPushX(ctx, key, values...)
}

// LRange LRange
func LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	return rdb.LRange(ctx, key, start, stop)
}

// LRem LRem
func LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd {
	return rdb.LRem(ctx, key, count, value)
}

// LSet LSet
func LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd {
	return rdb.LSet(ctx, key, index, value)
}

// LTrim LTrim
func LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd {
	return rdb.LTrim(ctx, key, start, stop)
}

// RPop RPop
func RPop(ctx context.Context, key string) *StringCmd {
	return rdb.RPop(ctx, key)
}

// RPopLPush RPopLPush
func RPopLPush(ctx context.Context, source, destination string) *StringCmd {
	return rdb.RPopLPush(ctx, source, destination)
}

// RPush RPush
func RPush(ctx context.Context, key string, values ...interface{}) *IntCmd {
	return rdb.RPush(ctx, key, values...)
}

// RPushX RPushX
func RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd {
	return rdb.RPushX(ctx, key, values...)
}

//------------------------------------------------------------------------------

// SAdd SAdd
func SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd {
	return rdb.SAdd(ctx, key, members...)
}

// SCard SCard
func SCard(ctx context.Context, key string) *IntCmd {
	return rdb.SCard(ctx, key)
}

// SDiff SDiff
func SDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	return rdb.SDiff(ctx, keys...)
}

// SDiffStore SDiffStore
func SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	return rdb.SDiffStore(ctx, destination, keys...)
}

// SInter SInter
func SInter(ctx context.Context, keys ...string) *StringSliceCmd {
	return rdb.SInter(ctx, keys...)
}

// SInterStore SInterStore
func SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	return rdb.SInterStore(ctx, destination, keys...)
}

// SIsMember SIsMember
func SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd {
	return rdb.SIsMember(ctx, key, member)
}

// SMembers Redis `SMEMBERS key` command output as a slice.
func SMembers(ctx context.Context, key string) *StringSliceCmd {
	return rdb.SMembers(ctx, key)
}

// SMembersMap Redis `SMEMBERS key` command output as a map.
func SMembersMap(ctx context.Context, key string) *StringStructMapCmd {
	return rdb.SMembersMap(ctx, key)
}

// SMove SMove
func SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd {
	return rdb.SMove(ctx, source, destination, member)
}

// SPop Redis `SPOP key` command.
func SPop(ctx context.Context, key string) *StringCmd {
	return rdb.SPop(ctx, key)
}

// SPopN Redis `SPOP key count` command.
func SPopN(ctx context.Context, key string, count int64) *StringSliceCmd {
	return rdb.SPopN(ctx, key, count)
}

// SRandMember Redis `SRANDMEMBER key` command.
func SRandMember(ctx context.Context, key string) *StringCmd {
	return rdb.SRandMember(ctx, key)
}

// SRandMemberN Redis `SRANDMEMBER key count` command.
func SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd {
	return rdb.SRandMemberN(ctx, key, count)
}

// SRem SRem
func SRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	return rdb.SRem(ctx, key, members...)
}

// SUnion SUnion
func SUnion(ctx context.Context, keys ...string) *StringSliceCmd {
	return rdb.SUnion(ctx, keys...)
}

// SUnionStore SUnionStore
func SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	return rdb.SUnionStore(ctx, destination, keys...)
}
