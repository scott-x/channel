# channel

utils for golang channel

### api

- `func Merge[T any](ins ...<-chan T) <-chan T`: merge multiple channels into one channel
- `func Proceed[K, V any](in <-chan K, fn func(K) V) <-chan V`: receive one channel and return another