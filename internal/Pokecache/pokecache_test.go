package pokecache

import (
	"testing"
	"time"
)

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)

		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(cas.inputVal))
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Microsecond * 10
	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))
	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("%s should have been reaped", "key1")
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Microsecond * 10
	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))
	time.Sleep(interval / 2)

	_, ok := cache.Get("key1")
	if !ok {
		t.Errorf("%s should have been reaped", "key1")
	}
}
