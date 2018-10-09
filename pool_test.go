package gpool

import (
	"io"
	"os"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Init()
	if DefaultGpool == nil {
		t.Errorf("DefaultGpool can`t be nil")
	}

	if DefaultGpool.maxSize != MaxSize {
		t.Errorf("DefaultGpool max size error")
	}

	if DefaultGpool.minSize != MinSize {
		t.Errorf("DefaultGpool min size error")
	}
}

func TestGo(t *testing.T) {

	var res = make(chan int, 2)

	var funcs = []func(){
		func() {
			res <- 1
		},
		func() {
			res <- 1
		},
	}

	Init()

	go func() {
		i := <-res
		n := <-res

		if i != 1 || n != 1 {
			t.Errorf("Go Error, does`t run func() ")
		}

		t.Log(i, n)
	}()

	for _, v := range funcs {
		Go(v)
	}

	time.Sleep(1 * time.Second)

	Stop()
}

func TestNew(t *testing.T) {
	g := New(10, 100)

	if g == nil {
		t.Errorf("g can`t be nil")
	}
}

func BenchmarkGo(b *testing.B) {
	g := New(100, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Go(func() {
			// do some thing ? 1006960925
			// 1003118524
			// 1002560599
			_ = 1 + 1
		})
	}
	time.Sleep(1*time.Second)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.Go(func() {
			// do some thing ? 1006960925
			// 1003118524
			// 1002560599
			_ = 1 + 1
		})
	}
	g.Stop()
}


func BenchmarkGo2(b *testing.B) {
	g := New(100, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Go(func() {
			// do some thing ? 1006960925
			// 1003118524
			// 1002560599
			_ = 1 + 1
		})
	}
	time.Sleep(1*time.Second)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.Go(func() {
			io.WriteString(os.Stderr,"")
		})
	}
	g.Stop()
}
