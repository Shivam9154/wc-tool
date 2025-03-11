##  WC - Word Count

A clone of the Unix wc utility written in Go, implementing the core functionality of counting bytes, lines, words, and characters in text files.

## Features

1.Count bytes (-c)
2.Count lines (-l)
3.Count words (-w)
4.Count characters (-m)
5.Default mode (equivalent to -c -l -w)
6.Support for file input and standard input


## Building
go build -o wc ./cmd/wc


## Usage
```
# Count bytes
./wc -c test.txt

# Count lines
./wc -l test.txt

# Count words
./wc -w test.txt

# Count characters
./wc -m test.txt

# Default mode (lines, words, bytes)
./wc test.txt

# Read from standard input
cat test.txt | ./wc -l
```   
