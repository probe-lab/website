package main

import (
	"context"
	"fmt"
	"sort"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"golang.org/x/exp/slog"
)

func generateFig(ctx context.Context, pd *PlotDef, cfg *PlotConfig) (*grob.Fig, error) {
	fig := &grob.Fig{
		Layout: &pd.Layout,
	}

	dataSets := make(map[string]DataSet)
	for _, ds := range pd.Datasets {
		src, exists := cfg.Sources[ds.Source]
		if !exists {
			return nil, fmt.Errorf("unknown dataset source: %s", ds.Source)
		}
		var err error
		slog.Debug("getting dataset", "name", ds.Name, "source", ds.Source, "query", ds.Query)
		dataSets[ds.Name], err = src.GetDataSet(ctx, ds.Query)
		if err != nil {
			return nil, fmt.Errorf("failed to get dataset from source %q: %w", ds.Source, err)
		}
	}

	seriesByDataSet := make(map[string][]SeriesDef)
	for i, s := range pd.Series {
		if _, ok := dataSets[s.DataSet]; !ok {
			slog.Error(fmt.Sprintf("unknown dataset name %q in series %d", s.DataSet, i))
			continue
		}
		seriesByDataSet[s.DataSet] = append(seriesByDataSet[s.DataSet], s)
	}

	type LabeledSeries struct {
		Name      string
		SeriesDef *SeriesDef
		Labels    []any
		Values    []any
	}

	fig.Data = grob.Traces{}
	for dsname, series := range seriesByDataSet {
		ds := dataSets[dsname]

		// data is ordered in the same way as the definition
		// if series are generated from a groupfield then it uses that ordering
		data := make([]*LabeledSeries, 0)
		dataIndex := make(map[string]*LabeledSeries)

		slog.Info("reading dataset", "name", dsname)
		for ds.Next() {
			for _, s := range series {
				name := s.Name
				if s.GroupField != "" {
					if s.GroupValue == "*" {
						if name != "" {
							name = fmt.Sprintf("%s-%s", name, ds.Field(s.GroupField))
						} else {
							name = fmt.Sprintf("%s", ds.Field(s.GroupField))
						}
					} else if ds.Field(s.GroupField) != s.GroupValue {
						continue
					}
				}

				ls, ok := dataIndex[name]
				if !ok {
					slog.Debug("creating series", "plot", pd.Name, "series", name)
					ls = &LabeledSeries{
						Name:      name,
						SeriesDef: &s,
					}
					data = append(data, ls)
					dataIndex[ls.Name] = ls
				}
				if s.Labels != "" {
					ls.Labels = append(ls.Labels, ds.Field(s.Labels))
				}
				ls.Values = append(ls.Values, ds.Field(s.Values))
			}
		}
		if ds.Err() != nil {
			return nil, fmt.Errorf("dataset iteration ended with an error: %w", ds.Err())
		}

		sort.Slice(data, func(i, j int) bool {
			return data[i].Name < data[j].Name
		})

		for _, ls := range data {
			ls := ls
			switch ls.SeriesDef.Type {
			case SeriesTypeBar:
				trace := &grob.Bar{
					Type:        grob.TraceTypeBar,
					Name:        ls.Name,
					Orientation: grob.BarOrientationV,
					X:           ls.Labels,
					Y:           ls.Values,
				}

				if ls.SeriesDef.Color != "" {
					trace.Marker = &grob.BarMarker{
						Color: ls.SeriesDef.Color,
					}
				}

				fig.Data = append(fig.Data, trace)
			case SeriesTypeHBar:
				trace := &grob.Bar{
					Type:        grob.TraceTypeBar,
					Name:        ls.Name,
					Orientation: grob.BarOrientationH,
					X:           ls.Values,
					Y:           ls.Labels,
				}
				if ls.SeriesDef.Color != "" {
					trace.Marker = &grob.BarMarker{
						Color: ls.SeriesDef.Color,
					}
				}

				fig.Data = append(fig.Data, trace)
			case SeriesTypeLine:
				trace := &grob.Scatter{
					Type: grob.TraceTypeScatter,
					Name: ls.Name,
					X:    ls.Labels,
					Y:    ls.Values,
				}

				if ls.SeriesDef.Color != "" {
					trace.Marker = &grob.ScatterMarker{
						Color: ls.SeriesDef.Color,
					}
				}
				fig.Data = append(fig.Data, trace)
			case SeriesTypeBox:
				trace := &grob.Box{
					Type: grob.TraceTypeBox,
					Name: ls.Name,
					Y:    ls.Values,
				}

				if ls.SeriesDef.Color != "" {
					trace.Marker = &grob.BoxMarker{
						Color: ls.SeriesDef.Color,
					}
				}
				fig.Data = append(fig.Data, trace)
			case SeriesTypeHBox:
				trace := &grob.Box{
					Type: grob.TraceTypeBox,
					Name: ls.Name,
					X:    ls.Values,
				}

				if ls.SeriesDef.Color != "" {
					trace.Marker = &grob.BoxMarker{
						Color: ls.SeriesDef.Color,
					}
				}
				fig.Data = append(fig.Data, trace)
			default:
				return nil, fmt.Errorf("unsupported series type: %s", ls.SeriesDef.Type)
			}
		}

	}

	return fig, nil
}
