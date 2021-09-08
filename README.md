An example calling from C to Go to C, tested on Windows with MinGW.

First build the Go archive:
`go build -buildmode=c-archive -o hello.a .\main.go`

In addition to a static library, `hello.a`, there should also be a Go-generated C header file, `hello.h` (included here for reference), which is imported in `main.c`.

Next, build the C executable:
`gcc .\main.c .\hello.a -o hello.exe`

Expected output:
```
Hello from C!
Hello from Go!
Hello from C from Go!
```
