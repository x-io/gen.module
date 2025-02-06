package redis

import (
	"context"

	RS "github.com/go-redis/redis/v8"
)

//GeoLocation GeoLocation
type GeoLocation = RS.GeoLocation

//GeoRadiusQuery GeoRadiusQuery
type GeoRadiusQuery = RS.GeoRadiusQuery

//GeoLocationCmd GeoLocationCmd
type GeoLocationCmd = RS.GeoLocationCmd

//------------------------------------------------------------------------------

//GeoAdd
func GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) *IntCmd {
	return rdb.GeoAdd(ctx, key, geoLocation...)
}

// GeoRadius is a read-only GEORADIUS_RO command.
func GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd {
	return rdb.GeoRadius(ctx, key, longitude, latitude, query)
}

// GeoRadiusStore is a writing GEORADIUS command.
func GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *IntCmd {
	return rdb.GeoRadiusStore(ctx, key, longitude, latitude, query)
}

//GeoRadiusByMember GeoRadius is a read-only GEORADIUSBYMEMBER_RO command.
func GeoRadiusByMember(ctx context.Context, key, member string, query *GeoRadiusQuery) *GeoLocationCmd {
	return rdb.GeoRadiusByMember(ctx, key, member, query)
}

// GeoRadiusByMemberStore is a writing GEORADIUSBYMEMBER command.
func GeoRadiusByMemberStore(ctx context.Context, key, member string, query *GeoRadiusQuery) *IntCmd {
	return rdb.GeoRadiusByMemberStore(ctx, key, member, query)
}

//GeoDist GeoDist
func GeoDist(ctx context.Context, key string, member1, member2, unit string) *FloatCmd {
	return rdb.GeoDist(ctx, key, member1, member2, unit)
}

//GeoHash GeoHash
func GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd {
	return rdb.GeoHash(ctx, key, members...)
}

//GeoPos GeoPos
func GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd {
	return rdb.GeoPos(ctx, key, members...)
}
