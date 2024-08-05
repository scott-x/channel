package channel

import "sync"

// 扇入函数（组件），把多个channel中的数据发送到一个channel中
func Merge[T any](ins ...<-chan T) <-chan T {
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

// 工序: channel to channel
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

// 工序: channel to channel with sync.WaitGroup
func C2CWithWG[K, V any](in <-chan K, exchange func(k K, wg *sync.WaitGroup) V, wg *sync.WaitGroup) <-chan V {
	out := make(chan V)

	go func() {
		defer close(out)
		for c := range in {
			out <- exchange(c, wg)
		}
	}()

	return out
}
