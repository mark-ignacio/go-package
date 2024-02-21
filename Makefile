ifeq ($(ARCH),)
ARCH := $(shell go env GOARCH)
endif

.PHONY: all
all: clean
	go run ./prep/main.go -version $(VERSION) -arch $(ARCH)
	$(MAKE) nfpm VERSION=$(VERSION) PKG=rpm
	$(MAKE) nfpm VERSION=$(VERSION) PKG=deb

.PHONY: nfpm
nfpm:
	VERSION=$(VERSION) \
	ARCH=$(ARCH) \
		nfpm package -p $(PKG) -t dist

.PHONY: clean
clean:
	rm -rf package/*