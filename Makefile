.PHONY: build pack release clean

install: demo.go
		    @CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags "-s -w"  -o bin/demo demo.go

clean: bin/
		    @rm -rf bin/*