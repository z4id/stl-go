package main

func Contains[T comparable](slice_ []T, elem_ T) bool {
	/*
		Utility function to return true if an element is in slice.
	*/
	for _, item_ := range slice_ {
		if item_ == elem_ {
			return true
		}
	}
	return false
}
