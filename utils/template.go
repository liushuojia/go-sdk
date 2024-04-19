package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
)

func TemplateExecute(title, content string, data interface{}) (string, error) {
	tmpl := template.New(title)
	tmpl, err := tmpl.Funcs(template.FuncMap{
		"TimeToDate": func(d float64, s string) string {
			return TimeToDate(int64(d), s)
		},
		"Status": func(d float64) bool {
			return int64(d) == 1
		},
		"Divisor": func(p, u float64) string {
			return fmt.Sprintf("%.2f", p/u)
		},
	}).Parse(content)
	if err != nil {
		return "", err
	}
	var m map[string]interface{}
	j, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(j, &m); err != nil {
		return "", err
	}

	var doc bytes.Buffer
	if err := tmpl.Execute(&doc, m); err != nil {
		return "", err
	}
	return doc.String(), nil
}
