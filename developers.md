## Run tests and code coverage like this:

1. Set GOPATH environment variable:

       $ export GOPATH=/path/to/your/goworkspace

2. Clone the repository to it's correct location

       $ git clone http://github.com/gima/jsonv "$GOPATH/src/github.com/gima/jsonv"

3. Navigate to the cloned jsonv repository.

       $ cd "$GOPATH/src/github.com/gima/jsonv"

4. Run the tests:
       
       $ go test ./src/...

5. Install code coverage tool

       $ go get github.com/axw/gocov

6. Run the code coverage tool
       
       $ "$GOPATH/bin/gocov" test ./src/... | "$GOPATH/bin/gocov" report
