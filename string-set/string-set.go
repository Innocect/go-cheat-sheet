package main

type stringSet map[string]struct{}

// 1. Add to set
func (s stringSet) Add(val string) {
	s[val] = struct{}{}
}

// 2. Remove from set
func (s stringSet) Remove(val string) {
	delete(s, val)
}

// 3. Copy a set into Set
func (s stringSet) Copy(val stringSet) {
	for k := range val {
		s[k] = struct{}{}
	}
}

// 4. Print i.e return in array

func (s stringSet) ConvertToSlice() (slice []string) {
	for k := range s {
		slice = append(slice, k)
	}
	return slice
}

// 5. If Key Exists
func (s stringSet) IsExists(val string) bool {
	if _, ok := s[val]; ok {
		return ok
	}
	return false
}

// 6. Convert slice to set\
func (s stringSet) ConvertToStringSet(val []string) {
	for _, str := range val {
		s[str] = struct{}{}
	}
}
