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

## Run code:

1. Build and run the source using `make run`
2. On the other side a json payload can be sent to the running tcp server as:
    * Using ncat on linux as :
        * `ncat 127.0.0.1 9090 < $path_to_json_file` or
        * `ncat 127.0.0.1 9090` and then type a valid json string on the cursor prompt
    * Using the test code in `test` package as:
        * `go build test/main.go`
        * `cd test;./main`
3. To see data buffered in-memory + data in file:
    * `ncat 127.0.0.1 8080 < $path_to_file_with_any_string` or `ncat 127.0.0.1 8080` and then type any string on the cursor prompt
    * Or modify the test script described above in 3.ii as per needs