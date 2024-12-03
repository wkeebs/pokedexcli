package main

func FilterEmpty(s []string) []string {
	var newSlice []string
	for _, s := range s {
		if s != "" {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}
