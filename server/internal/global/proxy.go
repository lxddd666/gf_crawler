package global

import (
	"math/rand"
	"sync"
	"time"
)

// SafeProxyList 线程安全的代理列表
type SafeProxyList struct {
	mu   sync.RWMutex
	list []string
	rand *rand.Rand
}

// NewSafeProxyList 创建新的线程安全代理列表
func NewSafeProxyList(proxies []string) *SafeProxyList {
	source := rand.NewSource(time.Now().UnixNano())
	return &SafeProxyList{
		list: proxies,
		rand: rand.New(source),
	}
}

// Get 获取当前代理列表的副本
func (s *SafeProxyList) Get() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 返回切片的深拷贝
	result := make([]string, len(s.list))
	copy(result, s.list)
	return result
}

// GetRandom 随机获取一个代理
func (s *SafeProxyList) GetRandom() (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.list) == 0 {
		return "", false
	}

	// 生成随机索引
	index := s.rand.Intn(len(s.list))
	return s.list[index], true
}

// Remove 移除指定的代理（如果存在）
func (s *SafeProxyList) Remove(proxy string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, p := range s.list {
		if p == proxy {
			// 通过连接移除元素前后的切片来移除元素
			s.list = append(s.list[:i], s.list[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveAll 移除所有指定的代理（支持批量移除）
func (s *SafeProxyList) RemoveAll(proxies []string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	count := 0
	for _, proxy := range proxies {
		for i := 0; i < len(s.list); i++ {
			if s.list[i] == proxy {
				s.list = append(s.list[:i], s.list[i+1:]...)
				count++
				i-- // 因为移除了一个元素，索引需要调整
			}
		}
	}
	return count
}

// Add 添加代理到列表
func (s *SafeProxyList) Add(proxy string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.list = append(s.list, proxy)
}

// Size 获取代理列表大小
func (s *SafeProxyList) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.list)
}
