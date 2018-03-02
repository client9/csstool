package csstool

// TagMatch determines if a given CSS identifier should be kept or removed
type TagMatch struct {
	tags map[string]bool
	keep bool
}

// NewTagMatch creates an initialized TagMatch object
func NewTagMatch(tags []string, keep bool) *TagMatch {
	tagmap := make(map[string]bool, len(tags))
	for _, tag := range tags {
		tagmap[tag] = true
	}
	delete(tagmap, "")
	if keep && len(tagmap) > 0 {
		tagmap["*"] = true
		tagmap[":"] = true
		tagmap["["] = true
	}
	return &TagMatch{tags: tagmap, keep: keep}
}

// Keep returns true if a tag is to be preserved
func (tm *TagMatch) Keep(val string) bool {
	if len(tm.tags) == 0 {
		return true
	}
	inmap := tm.tags[val]
	if tm.keep {
		return inmap
	}

	// !keep --> remove
	// if in map, then remove
	return !inmap
}

// Remove returns true if tag is to be dropped
func (tm *TagMatch) Remove(val []byte) bool {
	if len(tm.tags) == 0 {
		return false
	}
	return !tm.Keep(string(val))
}
