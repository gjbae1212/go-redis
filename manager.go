package redis

import (
	"fmt"
	"time"

	"github.com/gjbae1212/go-redis/internal"

	"github.com/gjbae1212/consistent"
	redigo "github.com/gomodule/redigo/redis"
)

type Manager interface {
	Do(command string, args ...interface{}) (interface{}, error)
	DoWithTimeout(timeout time.Duration, command string, args ...interface{}) (interface{}, error)
	Close() error
	Stats() (map[string]redigo.PoolStats, error)
}

// NewManager creates an object implemented Manager interface.
func NewManager(addrs []string) (Manager, error) {
	if len(addrs) == 0 {
		return nil, fmt.Errorf("[err] NewManager %w", internal.ErrorRedisEmpty)
	}
	g := &Group{
		Opt:   defaultOption,
		Ring:  consistent.New(),
		Pools: map[string]*Pool{},
	}
	for _, addr := range addrs {
		if err := g.addPool(addr); err != nil {
			return nil, err
		}
	}
	return Manager(g), nil
}
