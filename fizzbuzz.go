package fizzbuzz

import (
	"strconv"
	"sync"
)

type Set struct {
	Input int
	Res   string // number | FizzBuzz | Fizz | Buzz
}

func check(x int) *Set {
	set := &Set{}
	set.Input = x
	if x%5 == 0 && x%3 == 0 {
		set.Res = "FizzBuzz"
	} else if x%3 == 0 {
		set.Res = "Fizz"
	} else if x%5 == 0 {
		set.Res = "Buzz"
	} else {
		set.Res = strconv.Itoa(x)
	}

	return set
}

func DoSync(num int) []*Set {
	var list []*Set
	for x := 1; x <= num; x++ {
		list = append(list, check(x))
	}

	return list
}

func DoAsync(num, goroutineNum int) []*Set {
	ch := make(chan int, goroutineNum)
	chSort := make(chan *Set, goroutineNum)

	// 対象となる数値を送る
	go func() {
		defer close(ch)
		for h := 1; h <= num; h++ {
			ch <- h
		}
	}()

	// FizzBuzzる
	// この時点では順不同
	go func() {
		wg := sync.WaitGroup{}
		defer close(chSort)
		for i := 0; i < goroutineNum; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for x := range ch {
					chSort <- check(x)
				}
			}()
		}
		wg.Wait()
	}()

	// FizzBuzzった奴をsortする
	mux := sync.Mutex{}
	wgSort := sync.WaitGroup{}
	sorted := make([]*Set, num)
	for h := 0; h < goroutineNum; h++ {
		wgSort.Add(1)
		go func() {
			defer wgSort.Done()
			counter := 0
			for s := range chSort {
				mux.Lock()
				sorted[s.Input-1] = s
				counter++
				mux.Unlock()
			}
		}()
	}
	wgSort.Wait()

	return sorted
}
