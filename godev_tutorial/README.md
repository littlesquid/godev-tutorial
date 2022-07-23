## Tutorial: Getting start with Go ##

- A main function excutes by default when run the main package.
- Go execute init functions automatically at program startup, after global variable have been initialized, a common use of init function is to verify or repair correctness of the program state before real execution begins.

### Reference ###
- Standard Library : https://pkg.go.dev/std 

    the standard library packages got when installed Go.

- Go Commands : https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program

- External Package : https://pkg.go.dev/

- Authenticating Modules : https://go.dev/ref/mod#authenticating

- go.sum files : contains cryptographic hashes of the module's direct and indirect dependencies.
    Fields in a go.sum's line
        - a module path : the name of module the hash belongs to
        - a version : the version of module the hash belongs to. if end with /go.mod, the hash is for the go.mod file only; otherwise, the hash is for the file within the .zip file.
        - hash : consist of algorithm name and base64-encoded cryptographic hash, separated by a colon (:). currently support only SHA-256 (h1) hash algorithm