all: $GOPATH/bin/gopher-count

$GOPATH/bin/gopher-count: gopher-count.go
	go install
