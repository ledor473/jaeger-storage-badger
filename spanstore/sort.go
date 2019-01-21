package spanstore

import (
	"github.com/jaegertracing/jaeger/model"
	"sort"
)

type byTraceID []*model.TraceID

func (s byTraceID) Len() int      { return len(s) }
func (s byTraceID) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byTraceID) Less(i, j int) bool {
	if s[i].High < s[j].High {
		return true
	} else if s[i].High > s[j].High {
		return false
	}
	return s[i].Low < s[j].Low
}

// SortTraceIDs sorts a list of TraceIDs
func SortTraceIDs(traceIDs []*model.TraceID) {
	sort.Sort(byTraceID(traceIDs))
}