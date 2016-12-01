all: test install run
install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/bot_agent_bamboo/*.go
test:
	GO15VENDOREXPERIMENT=1 go test -cover `glide novendor`
vet:
	go tool vet .
	go tool vet --shadow .
lint:
	golint -min_confidence 1 ./...
errcheck:
	errcheck -ignore '(Close|Write)' ./...
check: lint vet errcheck
runnsqlookupd:
	nsqlookupd \
	-http-address=127.0.0.1:4161 \
	-tcp-address=127.0.0.1:4160
runnsqd:
	mkdir -p target/nsqd
	nsqd \
	-lookupd-tcp-address=127.0.0.1:4160 \
	-data-path=target/nsqd
run:
	bot_agent_bamboo \
	-logtostderr \
	-v=2 \
	-nsqd-address=localhost:4150 \
	-nsq-lookupd-address=localhost:4161
format:
	find . -name "*.go" -exec gofmt -w "{}" \;
	goimports -w=true .
prepare:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/Masterminds/glide
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	glide install
update:
	glide up
clean:
	rm -rf vendor
