package main

import (
	"context"
	"fmt"
	"time"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
)

// PlotConfig provides external configuration and context to the generation
// of a plot.
type PlotConfig struct {
	// BasisTime is the time of execution of the queries in a plot
	// Generally it is the current time but can be set to a time in the past
	BasisTime time.Time

	// Sources is a mapping of names to datasources. The names can be
	// referenced in a dataset definition
	Sources map[string]DataSource
}

type PlotDef struct {
	Name      string       `yaml:"name"`
	Frequency string       `yaml:"frequency"`
	Datasets  []DataSetDef `yaml:"datasets"`
	Series    []SeriesDef  `yaml:"series"`
	Scalars   []ScalarDef  `yaml:"scalars"`
	Layout    grob.Layout  `yaml:"layout"`
}

func (pd *PlotDef) ExecuteTemplates(ctx context.Context, cfg *PlotConfig) error {
	for i := range pd.Datasets {
		if err := pd.Datasets[i].ExecuteTemplates(ctx, cfg); err != nil {
			return err
		}
	}

	return nil
}

type DataSetDef struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
	Query  string `yaml:"query"`
}

func (ds *DataSetDef) ExecuteTemplates(ctx context.Context, cfg *PlotConfig) error {
	q, err := ExecuteTemplate(ctx, ds.Query, cfg)
	if err != nil {
		return fmt.Errorf("failed to expand query for dataset %q: %w", ds.Name, err)
	}
	ds.Query = q
	return nil
}

type SeriesDef struct {
	Type       SeriesType `yaml:"type"`
	Name       string     `yaml:"name"` // name of the series
	Color      string     `yaml:"color"`
	DataSet    string     `yaml:"dataset"`
	Labels     string     `yaml:"labels"`     // the name of the field the series should use for labels
	Values     string     `yaml:"values"`     // the name of the field the series should use for values
	GroupField string     `yaml:"groupfield"` // optional name of a field the series should use for grouping into related series
	GroupValue string     `yaml:"groupvalue"` // optional value of a field the series should use for grouping into related series
	Percent    bool       `yaml:"percent"`
}

type SeriesType string

const (
	SeriesTypeBar  SeriesType = "bar"  // vertical bars
	SeriesTypeHBar SeriesType = "hbar" // horizontal bars
	SeriesTypeLine SeriesType = "line" // lines
	SeriesTypeBox  SeriesType = "box"  // vertical box plot
	SeriesTypeHBox SeriesType = "hbox" // horizontal box plot
)

type ScalarDef struct {
	Type         ScalarType `yaml:"type"`
	Name         string     `yaml:"name"` // name of the scalar
	Color        string     `yaml:"color"`
	DataSet      string     `yaml:"dataset"`
	Value        string     `yaml:"value"`        // the name of the field in the dataset that should be used for the scalar value
	ValueSuffix  string     `yaml:"valueSuffix"`  // a string to append after the value
	ValuePrefix  string     `yaml:"valuePrefix"`  // a string to prepend to the value
	DeltaDataSet string     `yaml:"deltaDataset"` // the name of a dataset to use for a delta value
	DeltaValue   string     `yaml:"deltaValue"`   // the name of the field in the delta dataset that should be used for the scalar value
	DeltaType    DeltaType  `yaml:"deltaType"`    // the type of delta contained in the value field
}

type ScalarType string

const (
	ScalarTypeNumber ScalarType = "number" // display the scalar value as a number
)

type DeltaType string

const (
	DeltaTypeRelative DeltaType = "relative" // the delta is an absolute value and should be displayed with a relative % change to the scalar
)

type DataSource interface {
	GetDataSet(ctx context.Context, query string, params ...any) (DataSet, error)
}

type DataSeries struct {
	Labels []string
	Values []float64
}

type DataSet interface {
	Next() bool
	Err() error
	Field(name string) any
}

// ColorDoc represents a document that defines a set of named colors
type ColorDoc struct {
	Colors map[string]string `yaml:"name"`
}
