package smg

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var sitemapIndexNumberRegex = regexp.MustCompile(`-(\d+)\.xml$`)

func extractSortableKey(loc string) string {
	// Trim full URL to just the filename
	_, filename := path.Split(loc)

	// Extract numeric suffix
	matches := sitemapIndexNumberRegex.FindStringSubmatch(filename)
	if len(matches) == 2 {
		num, err := strconv.Atoi(matches[1])
		if err == nil {
			// Pad number to make it sortable as string (natural order)
			return fmt.Sprintf("%s-%06d", baseNameWithoutNumber(filename), num)
		}
	}

	// fallback to raw name
	return filename
}

func baseNameWithoutNumber(filename string) string {
	idx := strings.LastIndex(filename, "-")
	if idx > 0 {
		return filename[:idx]
	}
	return filename
}
