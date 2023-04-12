package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type StaticQueryJSON struct {
	X []any `json:"x"`
	Y []any `json:"y"`
}

type StaticDataSource struct{}

func (s *StaticDataSource) GetDataSet(_ context.Context, query string, params ...any) (DataSet, error) {
	var jq StaticQueryJSON
	err := json.NewDecoder(strings.NewReader(query)).Decode(&jq)
	if err != nil {
		return nil, fmt.Errorf("failed to decode static data: %w", err)
	}

	return NewStaticDataSet(map[string][]any{
		"x": jq.X,
		"y": jq.Y,
	}), nil
}

var _ DataSet = (*StaticDataSet)(nil)

type StaticDataSet struct {
	Data     map[string][]any
	rowcount int
	nextrow  int
	err      error
}

func NewStaticDataSet(data map[string][]any) *StaticDataSet {
	ds := &StaticDataSet{
		Data:     data,
		rowcount: -1,
	}
	for _, vs := range ds.Data {
		if ds.rowcount == -1 {
			ds.rowcount = len(vs)
			continue
		}
		if len(vs) != ds.rowcount {
			ds.err = fmt.Errorf("static dataset has unequal numbers of rows")
		}
	}

	return ds
}

func (s *StaticDataSet) ResetIterator() {
	s.nextrow = 0
}

func (s *StaticDataSet) Next() bool {
	if s.err != nil {
		return false
	}
	if s.nextrow == s.rowcount {
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
	if s.nextrow > s.rowcount {
		return nil
	}
	return col[s.nextrow-1]
}
