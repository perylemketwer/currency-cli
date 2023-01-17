.PHONY: currency

run:
	@go run main.go $(COIN)

build:
	@go build -o bin/currency-cli .

compile-macos:
	@echo "Compliling to MacOS..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/currency-cli-macos-amd64 .
	@GOOS=darwin GOARCH=arm64 go build -o bin/currency-cli-macos-arm64 .

compile-freebsd:
	@echo "Compliling to FreeBSD..."
	@GOOS=freebsd GOARCH=amd64 go build -o bin/currency-cli-freebsd-amd64 .
	@GOOS=freebsd GOARCH=386 go build -o bin/currency-cli-freebsd-386 .

compile-linux:
	@echo "Compliling to Linux..."
	@GOOS=linux GOARCH=amd64 go build -o bin/currency-cli-linux-amd64 .
	@GOOS=linux GOARCH=386 go build -o bin/currency-cli-linux-386 .

compile-windows:
	@echo "Compliling to Windows..."
	@GOOS=windows GOARCH=amd64 go build -o bin/currency-cli-windows-amd64 .
	@GOOS=windows GOARCH=386 go build -o bin/currency-cli-windows-386 .

compile:
	@echo "Compliling for every OS and Platform..."
	@$(MAKE) compile-macos compile-freebsd compile-linux compile-windows

test:
	@echo "Running tests..."
	@go test -v ./tests

clean:
	@rm -fv bin/*