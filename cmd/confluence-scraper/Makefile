.PHONY: build
build:
	@ARCH=`uname -m`; \
    echo "アーキテクチャ: $$ARCH"; \
	if [ "$$ARCH" = "x86_64" ]; then \
		GOARCH="amd64"; \
	elif [ "$$ARCH" = "aarch64" ]; then \
		GOARCH="arm64"; \
	else \
		echo "Unsupported architecture: $$ARCH"; \
		exit 1; \
	fi; \
	echo "GOARCH: $$GOARCH"; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-ldflags="-w -s" \
		-trimpath \
		-o confluence-scraper \
		main.go
	@echo "Build complete"
