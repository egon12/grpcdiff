package bodygenerator

import "strings"

type (
	Source interface {
		GetValue() []Value
	}

	Value struct {
		L1 string
		L2 string
		L3 string
		L4 string
		L5 string
		L6 string
	}
)

type SeperatedValueSource struct {
	value []string
}

func (s *SeperatedValueSource) Parse(input string) {
	s.value = strings.Split(input, ",")
}

func (s *SeperatedValueSource) GetValue() []Value {
	result := make([]Value, len(s.value))
	for i, v := range s.value {
		result[i] = Value{L1: v}
	}
	return result
}
