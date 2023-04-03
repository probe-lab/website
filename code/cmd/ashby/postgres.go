package main

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgDataSource struct {
	connstr  string
	poolOnce sync.Once
	err      error
	pool     *pgxpool.Pool
}

func NewPgDataSource(connstr string) *PgDataSource {
	return &PgDataSource{
		connstr: connstr,
	}
}

func (p *PgDataSource) GetDataSet(ctx context.Context, query string, params ...any) (DataSet, error) {
	p.poolOnce.Do(func() {
		conf, err := pgxpool.ParseConfig(p.connstr)
		if err != nil {
			p.err = fmt.Errorf("unable to parse connection string: %w", err)
			return
		}
		pool, err := pgxpool.NewWithConfig(context.Background(), conf)
		if err != nil {
			p.err = fmt.Errorf("unable to connect to database: %w", err)
			return
		}
		p.pool = pool
	})

	if p.err != nil {
		return nil, p.err
	}

	conn, err := p.pool.Acquire(ctx)
	if err != nil {
		p.err = fmt.Errorf("unable to connect to database: %w", err)
		return nil, fmt.Errorf("execute query: %w", err)
	}

	rows, err := conn.Query(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("execute query: %w", err)
	}

	return &PgDataSet{
		rows: rows,
	}, nil
}

var _ DataSet = (*PgDataSet)(nil)

type PgDataSet struct {
	rows    pgx.Rows
	rowdata []any
	fields  map[string]int
	err     error
}

func (s *PgDataSet) Next() bool {
	if s.err != nil {
		return false
	}
	s.rowdata = nil
	return s.rows.Next()
}

func (s *PgDataSet) Err() error {
	if s.err != nil {
		return s.err
	}
	return s.rows.Err()
}

func (s *PgDataSet) Field(name string) any {
	s.rowdata, s.err = s.rows.Values()
	if s.err != nil || s.rowdata == nil {
		return nil
	}
	if s.fields == nil {
		fds := s.rows.FieldDescriptions()
		s.fields = make(map[string]int, len(fds))
		for i, fd := range fds {
			s.fields[fd.Name] = i
		}
	}

	col, ok := s.fields[name]
	if !ok {
		return errors.New("unknown field")
	}

	return s.rowdata[col]
}
