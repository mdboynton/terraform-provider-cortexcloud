# Provider values
CC_PROVIDER_HOSTNAME = registry.terraform.io
CC_PROVIDER_NAMESPACE = PaloAltoNetworks
CC_PROVIDER_NAME = cortexcloud
CC_PROVIDER_BINARY = terraform-provider-${CC_PROVIDER_NAME}
CC_PROVIDER_VERSION ?= 0.0.0-dev

# OS and architecture of the system that will run the provider
# Must follow the schema "os_architecture"
TARGET_OS_ARCH ?= darwin_arm64

IS_CI_EXECUTION=0

# Retrieve operating system name and architecture 
os := $(shell uname -s | awk '{print tolower($0)}')
arch := $(shell uname -m)

# Strip whitespace from TARGET_OS_ARCH value
target_os_arch_stripped=$(shell echo "$(TARGET_OS_ARCH)" | xargs)

default: install

format:
	gofmt -l -w .

# Print warning message if target operating system architecture does not
# match the values returned by the system, or error message if this is
# being executed in a CI pipeline (dictated by the IS_CI_EXECTION value)
checkos:
ifneq ($(os)_$(arch), $(target_os_arch_stripped))
ifeq ($(IS_CI_EXECUTION), 0)
	$(info WARNING: Provided TARGET_OS_ARCH value "$(target_os_arch_stripped)" does not match the expected value for the detected operating system and architecture "$(os)_$(arch)". This may result in Terraform being unable to find the provider binary.)
else ifeq ($(IS_CI_EXECUTION), 1)
	$(error Provided OS_ARCH value "$(target_os_arch_stripped)" does not match the expected value for the detected operating system and architecture "$(os)_$(arch)".)
endif
endif

# Build provider binary
build: checkos
	@echo "Building provider ${CC_PROVIDER_BINARY}"
	@go build -o ${CC_PROVIDER_BINARY}

# Create plugin directory and move binary
install: build
	@echo "Creating plugin directory ~/.terraform.d/plugins/${CC_PROVIDER_HOSTNAME}/${CC_PROVIDER_NAMESPACE}/${CC_PROVIDER_NAME}/${CC_PROVIDER_VERSION}/${TARGET_OS_ARCH}"
	@mkdir -p ~/.terraform.d/plugins/${CC_PROVIDER_HOSTNAME}/${CC_PROVIDER_NAMESPACE}/${CC_PROVIDER_NAME}/${CC_PROVIDER_VERSION}/${TARGET_OS_ARCH}
	@echo "Moving binary to plugin directory..."
	@mv ${CC_PROVIDER_BINARY} ~/.terraform.d/plugins/${CC_PROVIDER_HOSTNAME}/${CC_PROVIDER_NAMESPACE}/${CC_PROVIDER_NAME}/${CC_PROVIDER_VERSION}/${TARGET_OS_ARCH}
	@echo "Done!"

# Delete provider binary from plugin directory
clean:
	rm -rf ~/.terraform.d/plugins/${CC_PROVIDER_HOSTNAME}/${CC_PROVIDER_NAMESPACE}/${CC_PROVIDER_NAME}

# Run acceptance test suite
acctest: build
	go test -v ./internal/acceptance/ -count=1
