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

	fig.Data = grob.Traces{}

	traces, err := seriesTraces(dataSets, pd.Series, cfg)
	if err != nil {
		return nil, fmt.Errorf("series traces: %w", err)
	}
	fig.Data = append(fig.Data, traces...)

	traces, err = scalarTraces(dataSets, pd.Scalars, cfg)
	if err != nil {
		return nil, fmt.Errorf("scalar tracess: %w", err)
	}
	fig.Data = append(fig.Data, traces...)

	return fig, nil
}

type LabeledSeries struct {
	Name      string
	SeriesDef *SeriesDef
	Labels    []any
	Values    []any
}

func seriesTraces(dataSets map[string]DataSet, seriesDefs []SeriesDef, cfg *PlotConfig) ([]grob.Trace, error) {
	var traces []grob.Trace

	seriesByDataSet := make(map[string][]SeriesDef)
	for i, s := range seriesDefs {
		if _, ok := dataSets[s.DataSet]; !ok {
			slog.Error(fmt.Sprintf("unknown dataset name %q in series %d", s.DataSet, i))
			continue
		}
		seriesByDataSet[s.DataSet] = append(seriesByDataSet[s.DataSet], s)
	}

	// data is ordered in the same way as the definition
	// TODO: fix ordering
	// if series are generated from a groupfield then it uses that ordering
	for dsname, series := range seriesByDataSet {
		ds := dataSets[dsname]

		data := make([]*LabeledSeries, 0)
		dataIndex := make(map[string]*LabeledSeries)

		slog.Info("reading dataset", "name", dsname)
		for ds.Next() {
			for _, s := range series {
				s := s
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
					slog.Debug("creating series", "series", name)
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
			if data[i].SeriesDef.order != data[j].SeriesDef.order {
				return data[i].SeriesDef.order < data[j].SeriesDef.order
			}
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

				if c := cfg.MaybeLookupColor(ls.SeriesDef.Color, ls.Name); c != "" {
					trace.Marker = &grob.BarMarker{
						Color: c,
					}
				}

				traces = append(traces, trace)
			case SeriesTypeHBar:
				trace := &grob.Bar{
					Type:        grob.TraceTypeBar,
					Name:        ls.Name,
					Orientation: grob.BarOrientationH,
					X:           ls.Values,
					Y:           ls.Labels,
				}
				if c := cfg.MaybeLookupColor(ls.SeriesDef.Color, ls.Name); c != "" {
					trace.Marker = &grob.BarMarker{
						Color: c,
					}
				}

				traces = append(traces, trace)
			case SeriesTypeLine:
				trace := &grob.Scatter{
					Type: grob.TraceTypeScatter,
					Name: ls.Name,
					X:    ls.Labels,
					Y:    ls.Values,
				}

				if c := cfg.MaybeLookupColor(ls.SeriesDef.Color, ls.Name); c != "" {
					trace.Marker = &grob.ScatterMarker{
						Color: c,
					}
				}
				traces = append(traces, trace)
			case SeriesTypeBox:
				trace := &grob.Box{
					Type: grob.TraceTypeBox,
					Name: ls.Name,
					Y:    ls.Values,
				}

				if c := cfg.MaybeLookupColor(ls.SeriesDef.Color, ls.Name); c != "" {
					trace.Marker = &grob.BoxMarker{
						Color: c,
					}
				}
				traces = append(traces, trace)
			case SeriesTypeHBox:
				trace := &grob.Box{
					Type: grob.TraceTypeBox,
					Name: ls.Name,
					X:    ls.Values,
				}

				if c := cfg.MaybeLookupColor(ls.SeriesDef.Color, ls.Name); c != "" {
					trace.Marker = &grob.BoxMarker{
						Color: c,
					}
				}
				traces = append(traces, trace)
			default:
				return nil, fmt.Errorf("unsupported series type: %s", ls.SeriesDef.Type)
			}
		}

	}

	return traces, nil
}

func scalarTraces(dataSets map[string]DataSet, scalarDefs []ScalarDef, cfg *PlotConfig) ([]grob.Trace, error) {
	// work out which dataset fields need to be read
	datasetFieldsUsed := make(map[string][]string)
	for _, s := range scalarDefs {
		if _, ok := dataSets[s.DataSet]; !ok {
			slog.Error(fmt.Sprintf("unknown dataset name %q for scalar %s", s.DataSet, s.Name))
			continue
		}
		datasetFieldsUsed[s.DataSet] = append(datasetFieldsUsed[s.DataSet], s.Value)

		if s.DeltaDataSet != "" {
			if _, ok := dataSets[s.DeltaDataSet]; !ok {
				slog.Error(fmt.Sprintf("unknown delta dataset name %q for scalar %s", s.DeltaDataSet, s.Name))
				continue
			}
			datasetFieldsUsed[s.DeltaDataSet] = append(datasetFieldsUsed[s.DeltaDataSet], s.DeltaValue)
		}
	}

	// read one row from each referenced dataset and record the relevant fields
	dsValues := make(map[string]map[string]float64)
	for dsname, fields := range datasetFieldsUsed {
		ds := dataSets[dsname]

		if !ds.Next() {
			if ds.Err() != nil {
				slog.Error(fmt.Sprintf("error reading dataset %q: %v", dsname, ds.Err()))
				continue
			}
			slog.Error(fmt.Sprintf("no rows found for dataset %q", dsname))
			continue
		}

		dsValues[dsname] = make(map[string]float64)

		for _, f := range fields {
			v := ds.Field(f)
			switch tv := v.(type) {
			case float64:
				dsValues[dsname][f] = tv
			case int64:
				dsValues[dsname][f] = float64(tv)
			default:
				slog.Error(fmt.Sprintf("field %q not read from dataset %q: (type %T)", f, dsname, v))
				dsValues[dsname][f] = 0
			}
		}
	}

	var traces []grob.Trace

	domainX := 1.0 / float64(len(scalarDefs))
	for idx, s := range scalarDefs {
		switch s.Type {
		case ScalarTypeNumber:
			trace := &grob.Indicator{
				Type: grob.TraceTypeIndicator,
				Name: s.Name,
				Mode: "number",
				Number: &grob.IndicatorNumber{
					Suffix: s.ValueSuffix,
				},
				Domain: &grob.IndicatorDomain{
					Column: int64(idx),
					X:      []float64{domainX * float64(idx), domainX * float64(idx+1)},
				},
				Title: &grob.IndicatorTitle{
					Text: s.Name,
				},
			}

			v, ok := dsValues[s.DataSet][s.Value]
			if !ok {
				slog.Error(fmt.Sprintf("missing value field for scalar %s", s.Name))
				continue
			}
			trace.Value = v

			if s.DeltaDataSet != "" {
				dv, ok := dsValues[s.DeltaDataSet][s.DeltaValue]
				if !ok {
					slog.Error(fmt.Sprintf("missing delta value field for scalar %s", s.Name))
					continue
				}
				switch s.DeltaType {
				case DeltaTypeRelative:
					trace.Delta = &grob.IndicatorDelta{
						Reference:   dv,
						Relative:    grob.True,
						Valueformat: ".2%",
					}
					trace.Mode = "number+delta"
					if c := cfg.MaybeLookupColor(s.IncreaseColor, ""); c != "" {
						trace.Delta.Increasing = &grob.IndicatorDeltaIncreasing{
							Color: c,
						}
					}
					if c := cfg.MaybeLookupColor(s.DecreaseColor, ""); c != "" {
						trace.Delta.Decreasing = &grob.IndicatorDeltaDecreasing{
							Color: c,
						}
					}
				default:
					return nil, fmt.Errorf("unsupported delta type: %s", s.DeltaType)
				}
			}

			traces = append(traces, trace)
		default:
			return nil, fmt.Errorf("unsupported scalar type: %s", s.Type)

		}
	}
	return traces, nil
}
