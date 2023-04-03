package main

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"
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
	}).Parse(source)
	if err != nil {
		return "", fmt.Errorf("parse query template: %w", err)
	}

	data := map[string]any{
		"Now":         cfg.BasisTime,
		"StartOfHour": cfg.BasisTime.Truncate(time.Hour),
		"StartOfDay":  cfg.BasisTime.Truncate(24 * time.Hour),
		"StartOfWeek": cfg.BasisTime.Truncate(7 * 24 * time.Hour),
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
