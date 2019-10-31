# Word Counter (Go Version)
Counts the number of occurrences of every word in all text files within a folder.

Goals of this project are:
* Learn the Go language.
* Compare with implementations of the same project in C++, Node, Java, Python.

# Building
```
go build cmd/main.go
```

# Running
```
go run cmd/main.go <folder>
```
or with previous build step:
```
./main <folder>
```
For example, count the words of all files in folder `./data` and write the results in file `wordcounts.txt`:
```
./main ./data > wordcounts.txt
```

# Testing
```
go test ./test/...
```

# License
MIT
