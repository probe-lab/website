package main

import (
	"errors"
)

var _ DataSet = (*StaticDataSet)(nil)

type StaticDataSet struct {
	Data    map[string][]any
	nextrow int
	err     error
}

func (s *StaticDataSet) ResetIterator() {
	s.nextrow = 0
}

func (s *StaticDataSet) Next() bool {
	if s.err != nil {
		return false
	}
	if s.nextrow == len(s.Data) {
		return false
	}
	s.nextrow++
	return true
}

func (s *StaticDataSet) Err() error {
	return s.err
}

func (s *StaticDataSet) Field(name string) any {
	if s.nextrow == 0 {
		return nil
	}
	col, ok := s.Data[name]
	if !ok {
		return errors.New("unknown field")
	}
	if s.nextrow >= len(col) {
		return nil
	}
	return col[s.nextrow-1]
}
