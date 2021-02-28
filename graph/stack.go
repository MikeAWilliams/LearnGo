package graph

import "sync"

type Stack struct {
	storage []ContainerItem
	mutex   sync.Mutex
}

func (s *Stack) Add(item ContainerItem) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.storage = append(s.storage, item)
}

func (s *Stack) Remove() ContainerItem {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	len := len(s.storage)
	if 0 == len {
		return nil
	}
	last := len - 1
	result := s.storage[last]
	s.storage = s.storage[:last]
	return result
}

func (s *Stack) IsEmpty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return 0 == len(s.storage)
}
