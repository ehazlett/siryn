REPO=ehazlett/siryn
COMMIT=`git rev-parse --short HEAD`

all: image

build: build-binaries build-cleanup

build-binaries: 
	@docker build -t siryn-build -f Dockerfile.build .
	@docker run -d -ti --name siryn-build siryn-build sh
	@docker cp siryn-build:/go/bin/prometheus ./
	@docker cp siryn-build:/go/bin/pushgateway ./
	@cd cmd/siryn ; go build -a -tags 'netgo' -ldflags "-w -X github.com/ehazlett/siryn/version.GitCommit $(COMMIT) -linkmode external -extldflags -static" .
	@cp cmd/siryn/siryn .

build-cleanup:
	@docker rm -fv siryn-build > /dev/null 2>&1 || true

image: build build-cleanup
	@docker build -t $(REPO) .
	@echo "Image created: $(REPO)"

clean: build-cleanup
	@rm -f pushgateway prometheus siryn cmd/siryn/siryn

.PHONY: build build-cleanup image clean
