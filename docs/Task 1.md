# Deploy the Engine: Clone the open-source repository for Chronos - A static race detector for the go language. Follow its setup guide and run the analyzer against a simple multi-threaded Go file that uses basic mutex locking.

## Issues I had:
- Chronos only runs using go 1.15. 
- Works only correctly if the {VCS}/{Company}/{module} folder structure is present.

## Tests:
1. I tested a program using mutex locking:
[mutex/main.go](../src/github.com/seminar/mutex/main.go)
    - Chronos outputs this:
    ```
    No data races found
    ```

2. I tested the race example from the Chronos docs to make sure Chronos works properly
    - Chronos outputs this:
    [race/main.go](../src/github.com/seminar/race/main.go)
    ```
    Potential race condition:
    Access1:
            a.X = 2
            ^
    /workspaces/Seminar/src/github.com/seminar/race/main.go:7:6 ->
    /workspaces/Seminar/src/github.com/seminar/race/main.go:12:4
    
    Access2:
            a.X = 1
            ^
    /workspaces/Seminar/src/github.com/seminar/race/main.go:7:6 ->
    /workspaces/Seminar/src/github.com/seminar/race/main.go:11:3 ->
    /workspaces/Seminar/src/github.com/seminar/race/main.go:10:5 
    ```