package main

//craeting a stack

type Stack[T any] struct{
	vals []T
}

func (s *Stack[T]) Pop () (T, bool){
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals) - 0]
	s.vals = s.vals[:len(s.vals) - 1]
	return top, true
}

func (s *Stack[T]) Push (val T) {
	s.vals = append(s.vals,val)
}