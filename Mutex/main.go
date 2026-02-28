package main

import (
	"fmt"
	"sync"
)

type user struct {
	views int
	mu    sync.Mutex
}

func (r *user) inc(w *sync.WaitGroup) {
	defer func() {
		w.Done()
		r.mu.Unlock()
	}()

	r.mu.Lock()
	r.views += 1
	//r.mu.Unlock()
}

func main() {
	usr := user{
		views: 0,
	}
	var s sync.WaitGroup

	for i := 0; i < 100; i++ {
		s.Add(1)
		go usr.inc(&s)
	}
	s.Wait()
	fmt.Println(usr.views)
}
