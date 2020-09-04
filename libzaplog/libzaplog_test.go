package libzaplog

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestZapInitial(t *testing.T) {
	logger, err := ZapInitial()
	if nil != err {
		t.Error(err)
	}
	if logger == nil {
		t.Error("logger: nil")
	}
}

func TestSugar(t *testing.T) {
	sugar := zap.NewExample().Sugar()
	defer func() {
		_ = sugar.Sync()
	}()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")
}

func TestLogger(t *testing.T) {
	logger := zap.NewExample()
	defer func() {
		_ = logger.Sync()
	}()
	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

// go test -bench=. -benchmem -run="BenchmarkSugar0"
//  351733	      3702 ns/op	     194 B/op	       2 allocs/op
func BenchmarkSugar0(b *testing.B) {
	sugar := zap.NewExample().Sugar()
	defer func() {
		_ = sugar.Sync()
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sugar.Infow("failed to fetch URL",
			"url", "http://example.com",
			"attempt", 3,
			"backoff", time.Second,
		)
	}
}

// go test -bench=. -benchmem -run="BenchmarkSugar1"
//  330472	      3657 ns/op	     194 B/op	       2 allocs/op
func BenchmarkSugar1(b *testing.B) {
	sugar := zap.NewExample().Sugar()
	defer func() {
		_ = sugar.Sync()
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sugar.Info("failed to fetch URL",
			zap.String("url", "http://example.com"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second))
	}
}

// go test -bench=. -benchmem -run="BenchmarkSugar2"
//  352872	      4442 ns/op	     194 B/op	       2 allocs/op
func BenchmarkSugar2(b *testing.B) {
	sugar := zap.NewExample().Sugar()
	defer func() {
		_ = sugar.Sync()
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sugar.Infof("failed to fetch URL: url:%s, attempt:%d, backoff:%d",
			"http://example.com", 3, time.Second)
	}
}

// go test -bench=. -benchmem -run="BenchmarkLogger"
//  338037	      3690 ns/op	     194 B/op	       2 allocs/op
func BenchmarkLogger(b *testing.B) {
	logger := zap.NewExample()
	defer func() {
		_ = logger.Sync()
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("failed to fetch URL",
			zap.String("url", "http://example.com"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
}
