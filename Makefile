NAME=ttoys

ifeq ($(shell uname -m),x86_64)
  ARCH?=x86_64
else ifeq ($(shell uname -m),arm64)
  ARCH?=aarch64
endif

ifeq ($(shell uname -o),Darwin)
  OS?=darwin
else
  OS?=linux
endif

.PHONY: check
check:  ## Run nix flake check
	nix build \
		.\#checks.$(ARCH)-$(OS).$(NAME) \
		--print-build-logs
