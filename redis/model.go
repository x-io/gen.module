package redis

import (
	RS "github.com/go-redis/redis/v8"
)

//Options Options
type Options = RS.Options

//PubSub PubSub
type PubSub = RS.PubSub

//Pipeliner Pipeliner
type Pipeliner = RS.Pipeliner

//Cmder Cmder
type Cmder = RS.Cmder

//PoolState PoolStats
type PoolState = RS.PoolStats

//Cmd Cmd
type Cmd = RS.Cmd

//SliceCmd SliceCmd
type SliceCmd = RS.SliceCmd

//StatusCmd StatusCmd
type StatusCmd = RS.StatusCmd

//SlowLogCmd SlowLogCmd
type SlowLogCmd = RS.SlowLogCmd

//BoolCmd BoolCmd
type BoolCmd = RS.BoolCmd

//BoolSliceCmd BoolSliceCmd
type BoolSliceCmd = RS.BoolSliceCmd

//IntCmd IntCmd
type IntCmd = RS.IntCmd

//IntSliceCmd IntSliceCmd
type IntSliceCmd = RS.IntSliceCmd

//FloatCmd FloatCmd
type FloatCmd = RS.FloatCmd

//StringCmd StringCmd
type StringCmd = RS.StringCmd

//StringSliceCmd StringSliceCmd
type StringSliceCmd = RS.StringSliceCmd

//StringIntMapCmd StringIntMapCmd
type StringIntMapCmd = RS.StringIntMapCmd

//StringStructMapCmd StringStructMapCmd
type StringStructMapCmd = RS.StringStructMapCmd

//StringStringMapCmd StringStringMapCmd
type StringStringMapCmd = RS.StringStringMapCmd

//ScanCmd ScanCmd
type ScanCmd = RS.ScanCmd

//GeoPosCmd GeoPosCmd
type GeoPosCmd = RS.GeoPosCmd

//ZSliceCmd ZSliceCmd
type ZSliceCmd = RS.ZSliceCmd

//XStreamSliceCmd XStreamSliceCmd
type XStreamSliceCmd = RS.XStreamSliceCmd

//XPendingCmd XPendingCmd
type XPendingCmd = RS.XPendingCmd

//XPendingExtCmd XPendingExtCmd
type XPendingExtCmd = RS.XPendingExtCmd

//XMessageSliceCmd XMessageSliceCmd
type XMessageSliceCmd = RS.XMessageSliceCmd

//XInfoGroupsCmd XInfoGroupsCmd
type XInfoGroupsCmd = RS.XInfoGroupsCmd

//XInfoStreamCmd XInfoStreamCmd
type XInfoStreamCmd = RS.XInfoStreamCmd

//ZWithKeyCmd ZWithKeyCmd
type ZWithKeyCmd = RS.ZWithKeyCmd
