package main

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (s *StringIntMap) Add(key string, value int) {
	s.data[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}

func (s *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(s.data))
	for k, v := range s.data {
		newMap[k] = v
	}
	return newMap
}

func (s *StringIntMap) Exists(key string) bool {
	if _, ok := s.data[key]; ok {
		return true
	}
	return false
}

func (s *StringIntMap) Get(key string) (int, bool) {
	if val, ok := s.data[key]; ok {
		return val, ok
	}
	return 0, false
}
