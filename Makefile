.PHONY: currency

run:
	@go run main.go $(COIN)

build:
	@go build -o bin/currency .

compile-macos:
	@echo "Compliling to MacOS..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/currency-cli-macos-amd64/currency .
	@GOOS=darwin GOARCH=arm64 go build -o bin/currency-cli-macos-arm64/currency .

compile-freebsd:
	@echo "Compliling to FreeBSD..."
	@GOOS=freebsd GOARCH=amd64 go build -o bin/currency-cli-freebsd-amd64/currency .
	@GOOS=freebsd GOARCH=386 go build -o bin/currency-cli-freebsd-386/currency .

compile-openbsd:
	@echo "Compliling to OpenBSD..."
	@GOOS=openbsd GOARCH=amd64 go build -o bin/currency-cli-openbsd-amd64/currency .
	@GOOS=openbsd GOARCH=386 go build -o bin/currency-cli-openbsd-386/currency .

compile-linux:
	@echo "Compliling to Linux..."
	@GOOS=linux GOARCH=amd64 go build -o bin/currency-cli-linux-amd64/currency .
	@GOOS=linux GOARCH=386 go build -o bin/currency-cli-linux-386/currency .

compile-windows:
	@echo "Compliling to Windows..."
	@GOOS=windows GOARCH=amd64 go build -o bin/currency-cli-windows-amd64/currency.exe .
	@GOOS=windows GOARCH=386 go build -o bin/currency-cli-windows-386/currency.exe .

compile:
	@echo "Compliling for every OS and Platform..."
	@$(MAKE) compile-macos compile-freebsd compile-openbsd compile-linux compile-windows

test:
	@echo "Running tests..."
	@go test -v ./tests

clean:
	@rm -rf ./bin/
