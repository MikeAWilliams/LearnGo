package maw

import "sync"

type Queue struct {
	storage []ContainerItem
	mutex   sync.Mutex
}

func (s *Queue) Add(item ContainerItem) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.storage = append(s.storage, item)
}

func (s *Queue) Remove() ContainerItem {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	len := len(s.storage)
	if 0 == len {
		return nil
	}
	result := s.storage[0]
	s.storage = s.storage[1:len]
	return result
}

func (s *Queue) IsEmpty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return 0 == len(s.storage)
}
