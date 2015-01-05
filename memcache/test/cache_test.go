package test

import (
	"mlib/memcache"
	"testing"
)

func TestOrdiCache(t *testing.T) {
	cache := memcache.Cacher(memcache.RANDOM, 10, 10)
	cache.Set("aaa", "sss")
	cache.Set("aaa", "bbb")
	cache.Set("ddd", "ddd")
	cache.Set("dded", "ddd")
	cache.Set("dddf", 123)
	cache.Update("dddd", 133)
	t.Log(cache.Get("dddd"))
	t.Log(cache.Len(), cache.Cap())
}
