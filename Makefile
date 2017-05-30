# Go parameters
BUILDPATH=$(CURDIR)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOBENCH=$(GOCMD) test -check.b
GODEP=$(GOTEST) -i
GOFMT=gofmt -w
GOGET=$(GOCMD) get
GOTESTFLAGS=-check.vv

# Package lists
TOPLEVEL_PKG := github.com/theodesp/go-shuffled-queue

get:
	@echo "getting all dependencies"
	$(GOGET)

build:
	$(GOINSTALL) $(TOPLEVEL_PKG)
	@echo "build sucessfully"

clean:
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg

test:
	@echo "Running Tests..."
	$(GOTEST) $(GOTESTFLAGS) $(TOPLEVEL_PKG)


bench:
	@echo "Running Benchmarks..."
	$(GOTEST) $(TOPLEVEL_PKG)


format:
	$(GOFMT) $(TOPLEVEL_PKG)

all: get build