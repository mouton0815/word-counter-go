# Word Counter (Go Version)
Counts the number of occurrences of every word in all text files within a folder.

Goals of this project are:
* Learn the Go language.
* Compare with implementations of the same project in C++, Node, [Java](https://github.com/mouton0815/word-counter-java), Python.

The project consists of
* A [path collector](internal/path-collector.go) that retrieves the path names of all `*.txt` files in a given folder and its subdirectories
and passes them to a channel named `pathQueue`.
* A [file reader](internal/file-reader.go) that reads the files and passes the content text to a [tokenizer](internal/tokenizer.go),
which splits the text into words and passes them to a channel name `wordQueue`.
* A number of [workers](internal/worker.go) that receive path names from a `pathQueue` and hands them over to the file reader.
* A [worker pool](internal/worker-pool.go) that spawns a worker for every available CPU and waits for their terminations.
* A [word counter](internal/word-counter.go) that listens to `wordQueue` and counts the number of occurrences for every word.
* A [main](cmd/main.go) program that wires the classes, starts the path collector, the worker pool, and the word counter.
Finally, it outputs the word lists ordered by decreasing number of occurences. 

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
