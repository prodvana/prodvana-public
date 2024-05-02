package sets

import "golang.org/x/exp/maps"

type Set[T comparable] struct {
	set map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		set: make(map[T]struct{}),
	}
}

func FromSlice[T comparable](sl []T) *Set[T] {
	s := NewSet[T]()
	for _, item := range sl {
		s.Add(item)
	}
	return s
}

func New[T comparable](sl ...T) *Set[T] {
	return FromSlice(sl)
}

func (s *Set[T]) Clone() *Set[T] {
	if s == nil {
		return &Set[T]{
			set: make(map[T]struct{}),
		}
	}
	return &Set[T]{
		set: maps.Clone(s.set),
	}
}

// get any object from the set. return empty value if set is empty.
func (s *Set[T]) GetAny() T {
	var empty T
	if s == nil {
		return empty
	}
	for item := range s.set {
		return item
	}
	return empty
}

func (s *Set[T]) Add(elem T) {
	s.set[elem] = struct{}{}
}

func (s *Set[T]) Remove(elem T) {
	delete(s.set, elem)
}

func (s *Set[T]) Contains(elem T) bool {
	if s == nil {
		return false
	}
	_, ok := s.set[elem]
	return ok
}

func (s *Set[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.set)
}

func (s *Set[T]) AddSet(set2 *Set[T]) *Set[T] {
	if set2 == nil {
		return s
	}

	for elem2 := range set2.set {
		s.set[elem2] = struct{}{}
	}
	return s
}

func (s *Set[T]) Iterate(f func(elem T) error) error {
	if s == nil {
		return nil
	}
	for elem := range s.set {
		if err := f(elem); err != nil {
			return err
		}
	}

	return nil
}

func Map[T comparable, ResultT any](s *Set[T], f func(elem T) (ResultT, error)) ([]ResultT, error) {
	sl := make([]ResultT, 0, s.Len())
	err := s.Iterate(func(elem T) error {
		mapped, err := f(elem)
		if err != nil {
			return err
		}
		sl = append(sl, mapped)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return sl, nil
}

// diff two sets, return elements of outer not in inner
func Diff[T comparable](outer, inner *Set[T]) *Set[T] {
	result := NewSet[T]()
	err := outer.Iterate(func(elem T) error {
		if !inner.Contains(elem) {
			result.Add(elem)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return result
}

// return a new set containing elements from both sets
func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	ret := s1.Clone()
	ret.AddSet(s2)
	return ret
}

// return a new set containing only elements that are present in both sets
func Intersect[T comparable](s1, s2 *Set[T]) *Set[T] {
	ret := s1.Clone()
	err := ret.Iterate(func(elem T) error {
		if !s2.Contains(elem) {
			ret.Remove(elem)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return ret
}

func (s *Set[T]) AsSlice() []T {
	if s == nil {
		return nil
	}
	sl := make([]T, 0, len(s.set))
	for elem := range s.set {
		sl = append(sl, elem)
	}
	return sl
}
