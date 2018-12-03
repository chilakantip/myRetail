setup:
	@go get -t -d -v 
	@go get -u golang.org/x/tools/cmd/...
	@go get -u github.com/alecthomas/gometalinter
	@go get -u github.com/pierrre/gotestcover
	@gometalinter -i -u

.PHONY: mcc setup
