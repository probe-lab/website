package main

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
)

func ExecuteTemplate(ctx context.Context, source string, cfg *PlotConfig) (string, error) {
	type ExecutionContext struct {
		Now         time.Time
		StartOfDay  time.Time
		StartOfWeek time.Time
	}

	t, err := template.New("").Funcs(template.FuncMap{
		"timestamptz": pgTimestampTZ,
		"timestamp":   pgTimestamp,
		"simpledate":  simpleDateFormat,
	}).Parse(source)
	if err != nil {
		return "", fmt.Errorf("parse query template: %w", err)
	}

	data := map[string]any{
		"Now":         cfg.BasisTime,
		"StartOfHour": cfg.BasisTime.Truncate(time.Hour),
		"StartOfDay":  cfg.BasisTime.Truncate(24 * time.Hour),
		"StartOfWeek": cfg.BasisTime.Truncate(7 * 24 * time.Hour),

		// The following are useful when formatting dates that are immediately before the start of the period
		// They are not really suitable for use as the end of a range in a query.
		"EndOfPreviousHour": cfg.BasisTime.Truncate(time.Hour).Add(-time.Nanosecond),
		"EndOfPreviousDay":  cfg.BasisTime.Truncate(24 * time.Hour).Add(-time.Nanosecond),
		"EndOfPreviousWeek": cfg.BasisTime.Truncate(7 * 24 * time.Hour).Add(-time.Nanosecond),
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return "", fmt.Errorf("execute query template: %w", err)
	}

	return buf.String(), nil
}

func pgTimestampTZ(t time.Time) string {
	return "'" + t.Format("2006-01-02 15:04:05 Z") + "'::timestamptz"
}

func pgTimestamp(t time.Time) string {
	return "'" + t.Format("2006-01-02 15:04:05") + "'::timestamp"
}

func simpleDateFormat(t time.Time) string {
	return t.Format("2 Jan 2006")
}

func ExecuteTemplateGrobString(ctx context.Context, str grob.String, cfg *PlotConfig) (grob.String, error) {
	switch tstr := str.(type) {
	case string:
		text, err := ExecuteTemplate(ctx, tstr, cfg)
		if err != nil {
			return str, fmt.Errorf("execute template: %w", err)
		}
		return text, nil
	case []string:
		for i, s := range tstr {
			text, err := ExecuteTemplate(ctx, s, cfg)
			if err != nil {
				return str, fmt.Errorf("execute template: %w", err)
			}
			tstr[i] = text
		}
		return tstr, nil
	}

	return str, nil
}
