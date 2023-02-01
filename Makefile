.PHONY: currency tests

run:
	@go run main.go $(COIN)

build:
	@go build -o bin/currency .

compile-dragonfly:
	@echo "Compliling to DragonFlyBSD..."
	@GOOS=dragonfly GOARCH=amd64 go build -o bin/currency-cli-dragonfly-amd64/currency .

compile-freebsd:
	@echo "Compliling to FreeBSD..."
	@GOOS=freebsd GOARCH=amd64 go build -o bin/currency-cli-freebsd-amd64/currency .
	@GOOS=freebsd GOARCH=arm64 go build -o bin/currency-cli-freebsd-arm64/currency .
	@GOOS=freebsd GOARCH=386 go build -o bin/currency-cli-freebsd-386/currency .

compile-linux:
	@echo "Compliling to Linux..."
	@GOOS=linux GOARCH=amd64 go build -o bin/currency-cli-linux-amd64/currency .
	@GOOS=linux GOARCH=arm64 go build -o bin/currency-cli-linux-arm64/currency .
	@GOOS=linux GOARCH=386 go build -o bin/currency-cli-linux-386/currency .

compile-netbsd:
	@echo "Compliling to NetBSD..."
	@GOOS=freebsd GOARCH=amd64 go build -o bin/currency-cli-netbsd-amd64/currency .
	@GOOS=freebsd GOARCH=arm64 go build -o bin/currency-cli-netbsd-arm64/currency .
	@GOOS=freebsd GOARCH=386 go build -o bin/currency-cli-netbsd-386/currency .

compile-openbsd:
	@echo "Compliling to OpenBSD..."
	@GOOS=openbsd GOARCH=amd64 go build -o bin/currency-cli-openbsd-amd64/currency .
	@GOOS=openbsd GOARCH=arm64 go build -o bin/currency-cli-openbsd-arm64/currency .
	@GOOS=openbsd GOARCH=386 go build -o bin/currency-cli-openbsd-386/currency .

compile-macos:
	@echo "Compliling to MacOS..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/currency-cli-darwin-amd64/currency .
	@GOOS=darwin GOARCH=arm64 go build -o bin/currency-cli-darwin-arm64/currency .

compile-solaris:
	@echo "Compliling to Solaris..."
	@GOOS=solaris GOARCH=amd64 go build -o bin/currency-cli-solaris-amd64/currency .

compile-windows:
	@echo "Compliling to Windows..."
	@GOOS=windows GOARCH=amd64 go build -o bin/currency-cli-windows-amd64/currency.exe .
	@GOOS=windows GOARCH=arm64 go build -o bin/currency-cli-windows-arm64/currency.exe .
	@GOOS=windows GOARCH=386 go build -o bin/currency-cli-windows-386/currency.exe .

compile:
	@echo "Compliling for every OS and Platform..."
	@$(MAKE) compile-dragonfly compile-freebsd compile-linux compile-netbsd compile-openbsd compile-macos compile-solaris compile-windows
	@echo "Done!"

tests:
	@echo "Running tests..."
	@go test -v ./tests

clean:
	@echo "Cleaning environment..."
	@rm -rfv ./bin/
	@rm -rfv ./dist/
	@echo "Done!"
