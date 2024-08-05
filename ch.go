package channel

import "sync"

// 扇入函数（组件），把多个channel中的数据发送到一个channel中
func merge[T any](ins ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	//把一个channel中的数据发送到out中
	p := func(in <-chan T) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}

	wg.Add(len(ins))

	//扇入，需要启动多个goroutine用于处理多个channel中的数据
	for _, cs := range ins {
		go p(cs)
	}

	//等待所有输入的数据ins处理完，再关闭输出out
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// channel to channel: transform one channel to another
func C2C[K, V any](in <-chan K, exchange func(K) V) <-chan V {
	out := make(chan V)

	go func() {
		defer close(out)
		for c := range in {
			out <- exchange(c)
		}
	}()

	return out
}

// channel to channel n times then merge
// n workers to process the input channel
func C2CNM[K, V any](in <-chan K, exchange func(K) V, n int) <-chan V {
	outs := make([]<-chan V, n)

	for i := 0; i < n; i++ {
		outs[i] = C2C(in, exchange)
	}

	return merge(outs...)
}
