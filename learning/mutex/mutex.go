package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) increaseViews(wg *sync.WaitGroup) {
	// defer wg.Done()
	defer func() {
		wg.Done();
		p.mu.Unlock()
	} ()
	p.mu.Lock()
	p.views += 1;
	// p.mu.Unlock()
}

func main() {

	var wg sync.WaitGroup

	myPost := post{views: 0}
	// for i:= range(102) {

	// myPost.increaseViews();
	
	// }
	for i := 0; i <100; i++ {
		wg.Add(1)
		go myPost.increaseViews(&wg)
		
	}
	wg.Wait()

	fmt.Println(myPost.views)

	// time.Sleep(time.Second * 1)

}