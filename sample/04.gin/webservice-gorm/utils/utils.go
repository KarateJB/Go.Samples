package utils

// Map: map A array to B array
func Map[S, D any](src []S, f func(S) D) []D {
	us := make([]D, len(src))
	for i := range src {
		us[i] = f(src[i])
	}
	return us
}
