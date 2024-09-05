package validate

import "fmt"

type Errors map[string][]string

func (e Errors) Add(key, message string) {
	e[key] = append(e[key], message)
}

func (e Errors) HasAny() bool {
	return len(e) > 0
}

func (e Errors) Required(key string, value interface{}) {
	if fmt.Sprintf("%v", value) == "" {
		e.Add(key, fmt.Sprintf("%s is required.", key))
	}
}

func (e Errors) GreatherThan(key string, value, min float64) {
	if value <= min {
		e.Add(key, fmt.Sprintf("%s must be greater than %.1f.", key, min))
	}
}
