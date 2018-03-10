package judge

import (
	"strconv"
	"strings"
)

func parseLimits(limits []string) []*_Limit {
	result := make([]*_Limit, 0)
	for _, v := range limits {
		new := parseLimit(v)
		if new == nil {
			return nil
		}
		isNew := true
		for _, l := range result {
			if l.Equals(new) {
				isNew = false
				break
			}
		}
		if isNew {
			result = append(result, new)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func parseLimit(s string) *_Limit {
	index := strings.IndexAny(s, "dhms")
	if index < 0 {
		return nil
	}
	secondS := s[:index]
	countS := s[index + 1:]

	second, err := strconv.ParseInt(secondS, 10, 32)
	if err != nil {
		return nil
	}
	count, err := strconv.ParseInt(countS, 10, 32)
	if err != nil {
		return nil
	}

	switch s[index] {
	case 'd':
		second *= 24 * 60 * 60
	case 'h':
		second *= 60 * 60
	case 'm':
		second *= 60
	case 's':
		second *= 1
	}
	if second > 365 * 24 * 60 * 60 {
		return nil
	}

	return &_Limit{
		Second: int(second),
		Count: int(count),
	}
}

type _Limit struct {
	Second int
	Count int
}

func (l *_Limit) Equals(li *_Limit) bool {
	if l == li {
		return true
	}
	return l != nil && l.Second == li.Second && l.Count == li.Count
}
