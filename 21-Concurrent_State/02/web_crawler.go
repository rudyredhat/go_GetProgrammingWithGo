// web crawler
// event loops explained further with a rover eg
package main

import "sync"

// Visited tracks whether web pages have been visited.
// Its methods may be used concurrently from multiple goroutines.
// We’ll store a map holding the URL of the web page and guard it with a mutex.
type Visited struct {
	// mu guards the visited map.
	mu      sync.Mutex
	visited map[string]int
}

//The code in the next listing defines a VisitLink method to be called when a link has been encountered;
// it returns the number of times that link has been encountered before.
// VisitLink tracks that the page with the given URL has
// been visited, and returns the updated link count.
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

var mu sync.Mutex // create a var
func main() {
	mu.Lock()         // lock mutes
	defer mu.Unlock() // unlock mutex from returning
}
