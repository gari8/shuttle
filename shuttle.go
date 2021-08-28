package shuttle

import (
	"fmt"
	"strings"
)

const (
	headerPart = "Content-"
)

type Shuttle struct {
	Boundary string `json:"boundary,omitempty"`
	Text string `json:"text,omitempty"`
}

func New(text, boundary string) Shuttle {
	return Shuttle{Boundary: boundary, Text: text}
}

func (s Shuttle) Launch (name string) string {
	reading := false
	var lines []string
	for _, line := range strings.Split(s.Text, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, headerPart) {
			if strings.Contains(line, fmt.Sprintf("name=\"%s\"", name)) {
				reading = !reading
			}
			continue
		}
		if reading {
			if strings.Contains(trimmed, s.Boundary) {
				return strings.TrimSpace(strings.Join(lines, "\n"))
			} else {
				lines = append(lines, line)
			}
		}
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}
