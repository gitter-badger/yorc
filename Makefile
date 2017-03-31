GOTOOLS = golang.org/x/tools/cmd/stringer github.com/tools/godep github.com/jteeuwen/go-bindata/...

PACKAGES=$(shell go list ./... | grep -v '/vendor/')

VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods \
         -nilfunc -printf -rangeloops -shift -structtags -unsafeptr

buildnformat: build format

build: test
	@echo "--> Running go build"
	@go generate $(PACKAGES)
	@CGO_ENABLED=0 go build

checks:
	@./build/checks.sh $(GOTOOLS)

dist: build
	@echo "--> Creating an archive"
	@tar czvf janus.tgz janus

test: checks
ifndef SKIP_TESTS
	@echo "--> Running go test"
	@export PATH=$$PWD/build:$$PATH; go test $(PACKAGES) $(TESTARGS) -p 1
endif


cover: 
	@go test $(PACKAGES) -p 1 --cover $(COVERARGS)  

format:
	@echo "--> Running go fmt"
	@go fmt $(PACKAGES)

vet:
	@echo "--> Running go tool vet $(VETARGS) ."
	@go list ./... \
		| grep -v '/vendor/' \
		| cut -d '/' -f 4- \
		| xargs -n1 \
			go tool vet $(VETARGS) ;\
	if [ $$? -ne 0 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for reviewal."; \
	fi

tools:
	go get -u -v $(GOTOOLS)

savedeps: checks
	@godep save -v ./...

restoredeps: checks
	@godep restore -v


.PHONY: buildnformat build cov checks test cover format vet tools
