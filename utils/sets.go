package utils

type Int64Set map[int64]struct{}

func (s Int64Set) Has(key int64) bool {
	_, ok := s[key]
	return ok
}

func (s Int64Set) Add(key int64) {
	s[key] = struct{}{}
}

func (s Int64Set) Items() []int64 {
	arr := make([]int64, 0, len(s))
	for k, _ := range s {
		arr = append(arr, k)
	}
	return arr
}
