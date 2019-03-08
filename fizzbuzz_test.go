package fizzbuzz

import (
	"runtime"
	"testing"
)

func BenchmarkDo(b *testing.B) {
	b.Run("num 10000", func(b *testing.B) {
		num := 10000
		b.Run("do sync", func(b *testing.B) {
			for h := 0; h < b.N; h++ {
				DoSync(num)
			}
		})
		b.Run("do async", func(b *testing.B) {
			runtime.GOMAXPROCS(1000)
			b.Run("goroutine 1", func(b *testing.B) {
				goroutineNum := 1
				for h := 0; h < b.N; h++ {
					DoAsync(num, goroutineNum)
				}
			})
			b.Run("goroutine 5", func(b *testing.B) {
				goroutineNum := 5
				for h := 0; h < b.N; h++ {
					DoAsync(num, goroutineNum)
				}
			})
			b.Run("goroutine 10", func(b *testing.B) {
				goroutineNum := 10
				for h := 0; h < b.N; h++ {
					DoAsync(num, goroutineNum)
				}
			})
			b.Run("goroutine 20", func(b *testing.B) {
				goroutineNum := 20
				for h := 0; h < b.N; h++ {
					DoAsync(num, goroutineNum)
				}
			})
			b.Run("goroutine 100", func(b *testing.B) {
				goroutineNum := 100
				for h := 0; h < b.N; h++ {
					DoAsync(num, goroutineNum)
				}
			})
			b.Run("goroutine 1000", func(b *testing.B) {
				goroutineNum := 1000
				for h := 0; h < b.N; h++ {
					DoAsync(num, goroutineNum)
				}
			})
		})
	})
}
