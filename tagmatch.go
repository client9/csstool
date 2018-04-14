package csstool

type matcher interface {
	Remove([]byte) bool
}

// EmptyMatcher this keeps all elements, or rather, doesn't remove anything
type EmptyMatcher struct{}

// Remove always returns false (i.e. keep everything)
func (em *EmptyMatcher) Remove(val []byte) bool {
	return false
}

// TagMatcher determines if a given CSS identifier should be kept or removed
// doesn't need to be public
type TagMatcher struct {
	tags map[string]bool
}

// NewTagMatcher creates an initialized TagMatch object
func NewTagMatcher(tags []string) *TagMatcher {
	tagmap := make(map[string]bool, len(tags))

	tagmap[""] = true
	tagmap["*"] = true
	tagmap[":root"] = true
	tagmap["::after"] = true
	tagmap["::before"] = true

	for _, tag := range tags {
		tagmap[tag] = true
	}

	return &TagMatcher{tags: tagmap}
}

// Remove returns true if tag is to be dropped
func (tm *TagMatcher) Remove(val []byte) bool {
	return !tm.tags[string(val)]
}

// AddSelector adds a selector to save
func (tm *TagMatcher) AddSelector(key string) {
	tm.tags[key] = true
}

// RemoveSelector deletes a selector to save
func (tm *TagMatcher) RemoveSelector(key string) {
	delete(tm.tags, key)
}
