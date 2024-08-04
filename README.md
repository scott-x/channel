# channel

utils for golang channel

### api

- `func Merge[T any](ins ...<-chan T) <-chan T`: merge multiple channels into one channel
- `func C2C[K, V any](in <-chan K, exchange func(K) V) <-chan V`: receive one channel and return another; exchange: change K instance to V instance 