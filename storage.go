package main

import (
	"context"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

type storage struct {
	ec     *elasticsearch.Client
	logger *zap.Logger

	mu sync.RWMutex
	// TODO - словари
	phrases    map[string][]int
	marketing  map[int][]int
	categories map[int][]int
}

func (s *storage) repeat(ctx context.Context, interval time.Duration, index string) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := s.update(ctx, index); err != nil {
				s.logger.Error("storage.update", zap.Error(err))
			}
		}
	}
}

func (s *storage) update(ctx context.Context, index string) error {

	// TODO - наполнить словари
	phrases := make(map[string][]int, len(s.phrases))
	marketing := make(map[int][]int, len(s.categories))
	categories := make(map[int][]int, len(s.categories))

	// limit := 1000
	// var lastPriority, lastProduct int

	for {
		// Ищем limit items

		//if len(result.Hits.Hits) != limit {
		//	break
		//}
		break
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// TODO - fill dicionaries
	s.phrases = phrases
	s.marketing = marketing
	s.categories = categories

	return nil
}

// TODO - примени их получше, когда поймешь как

// Category - возвращает список товаров категории
func (s *storage) Category(category int) []int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.categories[category]
}

// Exists - возвращает true если товар существует в категории
func (s *storage) Exists(category, product int) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if products, ok := s.categories[category]; ok {
		for _, id := range products {
			if product == id {
				return true
			}
		}
	}

	return false
}
