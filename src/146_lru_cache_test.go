package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	//var obj LRUCache
	//fmt.Println(obj.kv)
	capacity := 2
	obj := NewLRUCache(capacity)
	obj.Put(1, 11)
	obj.Put(2, 22)
	ret := obj.Get(1)
	fmt.Println("ret:", ret)
	assert.Equal(t, 11, ret)
	obj.Put(3, 3)
	obj.Put(3, 3)
	ret = obj.Get(2)
	fmt.Println("ret:", ret)
	assert.Equal(t, -1, ret)
	obj.Get(3)
	obj.Get(4)
	ret = obj.Get(1)
	fmt.Println("ret:", ret)
	assert.Equal(t, 11, ret)
	obj.Get(6)
	assert.Equal(t, 11, ret)
}
