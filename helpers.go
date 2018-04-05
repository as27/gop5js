package gop5js

import (
	"fmt"
	"strings"
)

func joinFloat64(fs ...float64) string {
	var ss []string
	for _, f := range fs {
		ss = append(ss, fmt.Sprintf("%f", f))
	}
	return strings.Join(ss, ",")
}
