// mutual exclusion
// mutex has two methods - lock and unlock
package main

import "sync"

var mu sync.Mutex // create a var
func main() {
	mu.Lock()         // lock mutes
	defer mu.Unlock() // unlock mutex from returning
}
