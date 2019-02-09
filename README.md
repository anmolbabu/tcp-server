# tcp-server

## Introduction

This is a simple tcp server written in golang. It does the following
at the moment:

Exposes tcp servers at 2 ports:
a. 9090:
    Any valid json message on this port, stores the passed message into a in-memory buffer until 20 seconds
    after which, asynchrously, a go-routine pulls out the data from the buffered storage and writes it to a
    file.
b. 8080:
    Any request on this port, prints all the data from the file + the in-memory buffered storage

## Possible Enhancements:

The ports and the time interval of asynchronous persistence of data in-memory to file, can be made configurable
via an entry in config file and the same can be used.

## Run code:

1. Build the source using `make`
2. Run the binary as `./tcp-server`
3. On the other side a json payload can be sent to the running tcp server as:
    i. Using ncat on linux as :
        a. `ncat 127.0.0.1 9090 < $path_to_json_file` or
        b. `ncat 127.0.0.1 9090` and then type a valid json string on the cursor prompt
   ii. Using the test code in `test` package as:
        a. `go build test/main.go`
        b. `cd test;./main`
4. To see data buffered in-memory + data in file:
    a. `ncat 127.0.0.1 8080 < $path_to_file_with_any_string` or `ncat 127.0.0.1 8080` and then type any string on the cursor prompt
    b. Or modify the test script described above in 3.ii as per needs