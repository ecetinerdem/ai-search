package main

func inExplored(needle Point, haystack []Point) bool {

	for _, x := range haystack {
		if needle.Row == x.Row && needle.Col == x.Col {
			return true
		}
	}
	return false
}
