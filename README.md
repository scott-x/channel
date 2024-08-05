# channel

utils for golang channel

### api

- `func Merge[T any](ins ...<-chan T) <-chan T`: merge multiple channels into one channel
- `func C2C[K, V any](in <-chan K, exchange func(K) V) <-chan V`: receive one channel and return another; exchange: change K instance to V instance 
- `func C2CWithWG[K, V any](in <-chan K, exchange func(k K, wg *sync.WaitGroup) V, wg *sync.WaitGroup) <-chan V `: you can do goroutine task when C2C