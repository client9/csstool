package csstool

// TagMatch determines if a given CSS identifier should be kept or removed
type TagMatch struct {
	tags map[string]bool
}

// NewTagMatch creates an initialized TagMatch object
func NewTagMatch(tags []string) *TagMatch {
	tagmap := make(map[string]bool, len(tags))
	for _, tag := range tags {
		tagmap[tag] = true
	}
	delete(tagmap, "")
	return &TagMatch{tags: tagmap}
}

// Keep returns true if a tag is to be preserved
func (tm *TagMatch) Keep(val string) bool {
	if len(tm.tags) == 0 || len(val) == 0 {
		return true
	}

	// now we know len(val) > 0
	//  special ones
	if val[0] == '*' || val[0] == ':' || val[0] == '[' {
		return true
	}
	return tm.tags[val]
}

// Remove returns true if tag is to be dropped
func (tm *TagMatch) Remove(val []byte) bool {
	if len(tm.tags) == 0 {
		return false
	}
	return !tm.Keep(string(val))
}
