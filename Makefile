#PKGS_WITH_OUT_EXAMPLES := $(shell go list ./... | grep -v 'cicomm/model')
PKGS_WITH_OUT_EXAMPLES := $(shell go list ./...)
PKGS_WITH_OUT_EXAMPLES_AND_UTILS := $(shell go list ./... | grep -v 'examples/\|utils/')
GO_FILES := $(shell find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" -print0 | xargs -0)

export GOPROXY=https://goproxy.io
export GO111MODULE=on

checkLocal: deps vet lint misspell staticcheck cyclo const

deps:
	go get golang.org/x/lint/golint
	go get github.com/fzipp/gocyclo
	go get github.com/client9/misspell/cmd/misspell
	go get honnef.co/go/tools/cmd/staticcheck
	go get github.com/jgautheron/goconst/cmd/goconst

vet:
	@echo "vet"
	go vet $(PKGS_WITH_OUT_EXAMPLES)

lint:
	@echo "golint"
	golint -set_exit_status $(PKGS_WITH_OUT_EXAMPLES_AND_UTILS)

misspell:
	@echo "misspell"
	misspell -source=text -error $(GO_FILES)

staticcheck:
	@echo "staticcheck"
	staticcheck $(PKGS_WITH_OUT_EXAMPLES)

cyclo:
	@echo "gocyclo"
#	gocyclo -over 15 $(GO_FILES)
	gocyclo -top 10 $(GO_FILES)

const:
	@echo "goconst"
	goconst $(PKGS_WITH_OUT_EXAMPLES)

test:
	@echo "test"
	go test -v -cover $(PKGS_WITH_OUT_EXAMPLES)

