package internal

import (
	"fmt"
	"path"
	"strings"
)

// MatchGlobParents is like path.Match, but will match if any component matches
// with optional anchoring.
func MatchGlobParents(pattern string, name string) (matched bool, err error) {
	var anchor bool
	if strings.HasPrefix(pattern, "/") {
		pattern, anchor = pattern[1:], true
	}
	for name != "" {
		// test against the full path
		if m, err := path.Match(pattern, name); m || err != nil {
			return m, err
		}
		// split it
		parent, base := path.Split(name)
		if !anchor {
			// test against the filename
			if m, err := path.Match(pattern, base); m || err != nil {
				return m, err
			}
		}
		// continue with the parent
		name = strings.TrimRight(parent, "/")
	}
	return false, nil
}

// FormatBytesSI formats the provided quantity with SI prefixes.
func FormatBytesSI(b int64) string {
	var neg bool
	if b < 0 {
		neg = true
		b *= -1
	}
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	if neg {
		return fmt.Sprintf("-%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
	} else {
		return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
	}
}
