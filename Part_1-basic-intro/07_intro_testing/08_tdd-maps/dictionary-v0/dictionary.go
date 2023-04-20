package main

// Search returns the value associated to the given word in the dictionary
func Search(d map[string]string, w string) string {
	return d[w]
}