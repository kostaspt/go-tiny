package cache

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/google/wire"
	"github.com/gosimple/slug"
)

var ProviderSet = wire.NewSet(NewInMemory, NewRedis)

type Cache interface {
	Get(ctx context.Context, key string, data any) error
	Set(ctx context.Context, key string, data any, ttl ...time.Duration) error
	Delete(ctx context.Context, key string) error
	Close() error
}

func GenerateKey(parts ...any) string {
	return strings.Join(interfacesToString(parts), "-")
}

func GeneratePrefixedKey(parts ...any) string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()

	prefix := slug.Make(funcName)

	return strings.Join(append([]string{prefix}, interfacesToString(parts)...), "-")
}

func interfacesToString(parts []any) []string {
	var s string
	sp := make([]string, len(parts))
	for i, p := range parts {
		switch p.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
			s = fmt.Sprintf("%d", p)
		case string:
			s = p.(string)
		default:
			s = fmt.Sprintf("%s", p)
		}
		sp[i] = slug.Make(s)
	}
	return sp
}
