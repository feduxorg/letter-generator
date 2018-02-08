package test

import "fmt"

func truncate(t testHelper, s string, count int) string {
	t.Helper()
	var shortenedS, suffix string

	if len(s) > count {
		shortenedS = s[:count-1]
		suffix = fmt.Sprintf(" ... (%d chars)", len(s))
	} else {
		shortenedS = s
		suffix = ""
	}
	return shortenedS + suffix
}
