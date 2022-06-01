package cache

import (
	"sync"
	"testing"
)

type KeyValue struct {
	key int64
	value int64
}

type KeyValuePadding struct {
	Key int64
	p1,p2,p3,p4,p5,p6,p7 int64
	Value int64
}

func BenchmarkCacheLineBaseline(b *testing.B) {
	for i:=0; i<b.N; i++ {
		kv := KeyValue{key:0, value: 0}
		var wg sync.WaitGroup
		times := 100000000
		wg.Add(2)
		go func ()  {
			for n:=0;n<times;n++{
				kv.key++
			}
			wg.Done()
		}()
		go func ()  {
			// for n:=0;n<times;n++{
			// 	kv.value++
			// }
			wg.Done()
		}()
		wg.Wait()
	}
}

func BenchmarkCacheLineFalseSharing(b *testing.B) {
	for i:=0; i<b.N; i++ {
		kv := KeyValue{key:0, value: 0}
		var wg sync.WaitGroup
		times := 100000000
		wg.Add(2)
		go func ()  {
			for n:=0;n<times;n++{
				kv.key++
			}
			wg.Done()
		}()
		go func ()  {
			for n:=0;n<times;n++{
				kv.value++
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

func BenchmarkCacheLinePadding(b *testing.B) {
	for i:=0; i<b.N; i++ {
		kv := KeyValuePadding{Key:0, Value: 0}
		var wg sync.WaitGroup
		times := 100000000
		wg.Add(2)
		go func ()  {
			for n:=0;n<times;n++{
				kv.Key++
			}
			wg.Done()
		}()
		go func ()  {
			for n:=0;n<times;n++{
				kv.Value++
			}
			wg.Done()
		}()
		wg.Wait()
	}
}