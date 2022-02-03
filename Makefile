NAME=go-identicons
GITHUB_USERNAME=poseidon-code
PROJECT=github.com/$(GITHUB_USERNAME)/$(NAME)
MAIN=main.go


init:
	@printf "\033[36;1m§\033[0m Implementing \033[36;1mGolang\033[0m template for '\033[36;1m$(NAME)\033[0m' as '\033[36;1m$(PROJECT)\033[0m'.\n"
	@printf "\033[37;1m»\033[0m Generating simple project structure...\n"
	@mkdir bin dist package
	@touch main.go
	@printf "\033[32;1m»\033[0m Generated all files\n"

	@printf "\n\033[37;1m»\033[0m Initializing go.mod for '\033[36;1m$(PROJECT)\033[0m'...\n"
	@go mod init $(PROJECT)
	@printf "\033[32;1m»\033[0m Initialized '$(PROJECT)'\n"

	@printf "\n\033[37;1m»\033[0m Initializing git repository...\n"
	@git init
	@git checkout -b main
	@printf "\033[32;1m»\033[0m Initialized (*main)\n"


run:
	@go run $(MAIN)



build:
	@printf "\033[37;1m»\033[0m Building '$(PROJECT)'...\n"
	go build -o bin/ $(PROJECT)
	@printf "\033[32;1m»\033[0m Built at 'bin/$(NAME)'\n\n"


compile:
	@printf "\033[37;1m»\033[0m Compiling for linux, windows, macos with x64 & x86 architecture...\n"
	GOOS=linux GOARCH=amd64 go build -o dist/$(NAME)-linux-amd64 $(MAIN)
	GOOS=linux GOARCH=386 GO386=softfloat go build -o dist/$(NAME)-linux-386 $(MAIN)
	GOOS=windows GOARCH=amd64 go build -o dist/$(NAME)-windows-amd64.exe $(MAIN)
	GOOS=windows GOARCH=386 GO386=softfloat go build -o dist/$(NAME)-windows-386.exe $(MAIN)
	GOOS=darwin GOARCH=amd64 go build -o dist/$(NAME)-darwin-amd64 $(MAIN)
	@printf "\033[32;1m»\033[0m Compiled to 'dist/'\n"


clean:
	@printf "\033[37;1m»\033[0m Cleaning Golang cached packages...\n"
	@go clean -modcache
	@printf "\033[32;1m»\033[0m Cleaned\n"


tidy:
	@printf "\033[37;1m»\033[0m Tidying Up dependencies...\n"
	@go mod tidy
	@printf "\033[32;1m»\033[0m Finished\n"


publish:
	GOPROXY=proxy.golang.org go list -m $(PROJECT)


purge:
	@printf "\033[37;1m»\033[0m Purging everything (except Makefile)...\n"
	@find * ! -name 'Makefile' -type d -exec rm -rfv {} +
	@rm -rfv .git/
	@printf "\033[32;1m»\033[0m Purged\n"