Day 9 - Select, Mutex, and WaitGroups

    Use select{} to handle multiple channels.
    Learn about sync.Mutex to prevent race conditions.
    Use WaitGroups to synchronize goroutines.

🛠 Project: Concurrent Web Scraper

Fetch page titles from multiple websites in parallel.

### Select{} - Handling Multiple Channels

package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    // Simulate async operations
    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "result from ch1"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "result from ch2"
    }()

    select {
    case res := <-ch1:
        fmt.Println("Received:", res)
    case res := <-ch2:
        fmt.Println("Received:", res)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout!")
    }
}
