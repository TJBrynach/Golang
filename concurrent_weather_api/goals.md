Day 9 - Select, Mutex, and WaitGroups

    Use select{} to handle multiple channels.
    Learn about sync.Mutex to prevent race conditions.
    Use WaitGroups to synchronize goroutines.

🛠 Project: Concurrent Web Scraper

Fetch page titles from multiple websites in parallel.

### Select{} - Handling Multiple Channels

```
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
```

### sync.Mutex - Preventing Race Conditions

```
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    counter := 0

    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++ // critical section
            mu.Unlock()
        }()
    }

    wg.Wait()
    fmt.Println("Final counter:", counter)
}
```

1. initiate mutex object mu
    -   this ensures only one goroutine can access a shared resource at a time, like a locked door. 1 goroutine can go in the room at a time.
2. initialising our shared variable
3. initiate waitgroup wg
    -   tracks a set of goroutines and waits until theyre all done. - a counter
4. loop to 5, launching 5 goroutines adding to the waitgroup
5. wg.Add(1) - informs waitgroup a goroutine has started
6. wg.Done() schedules to signal the goroutine is done when finished
7. mu.lock - locks the mutex before entering counter
8. wg.Wait() waits for all the goroutines to complete.