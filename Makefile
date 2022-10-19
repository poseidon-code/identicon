NAME=identicon
GITHUB_USERNAME=poseidon-code
PROJECT=github.com/$(GITHUB_USERNAME)/$(NAME)
MAIN=main.go
VERSION=1.2.0


compile:
	@printf "\033[37;1m»\033[0m Compiling for linux, windows, macos with x64 & x86 architecture...\n"
	GOOS=linux GOARCH=amd64 go build -o dist/$(NAME)-$(VERSION)-linux-amd64 $(MAIN)
	GOOS=windows GOARCH=amd64 go build -o dist/$(NAME)-$(VERSION)-windows-amd64.exe $(MAIN)
	GOOS=darwin GOARCH=amd64 go build -o dist/$(NAME)-$(VERSION)-darwin-amd64 $(MAIN)
	@printf "\033[32;1m»\033[0m Compiled to 'dist/'\n"

install:
	go build -o /usr/bin/ $(PROJECT)

uninstall:
	rm -rf /usr/bin/identicon

clean:
	go clean -modcache
