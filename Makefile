HOSTNAME = registry.terraform.io
NAMESPACE = PaloAltoNetworks
NAME = cortexcloud
BINARY = terraform-provider-${NAME}

VERSION ?= 0.0.0-dev
#OS_ARCH ?= darwin_amd64

# For Apple silicon machines:
OS_ARCH ?= darwin_arm64 

default: install

format:
	gofmt -l -w .

build:
	@echo "Building provider..."
	@go build -o ${BINARY}

install: build
	@echo "Creating plugin directory..."
	@mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	@echo "Moving binary to plugin directory..."
	@mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	@echo "Done!"

clean:
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}

acctest: build
	go test -v ./internal/acceptance/ -count=1
