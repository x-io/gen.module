package redis

import (
	"context"
)

//------------------------------------------------------------------------------

//Eval Eval
func Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd {
	return rdb.Eval(ctx, script, keys, args...)
}

//EvalSha EvalSha
func EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd {
	return rdb.EvalSha(ctx, sha1, keys, args...)
}

//ScriptExists ScriptExists
func ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd {
	return rdb.ScriptExists(ctx, hashes...)
}

//ScriptFlush ScriptFlush
func ScriptFlush(ctx context.Context) *StatusCmd {
	return rdb.ScriptFlush(ctx)
}

//ScriptKill ScriptKill
func ScriptKill(ctx context.Context) *StatusCmd {
	return rdb.ScriptKill(ctx)
}

//ScriptLoad ScriptLoad
func ScriptLoad(ctx context.Context, script string) *StringCmd {
	return rdb.ScriptLoad(ctx, script)
}
