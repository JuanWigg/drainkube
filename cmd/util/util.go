/*
Copyright © 2023 Juan Wiggenhauser <jgwiggenahuser@gmail.com>
*/
package util

func IsMapSubset[K, V comparable](m, sub map[K]V) bool {
	if len(sub) > len(m) {
		return false
	}
	for k, vsub := range sub {
		if vm, found := m[k]; !found || vm != vsub {
			return false
		}
	}
	return true
}
