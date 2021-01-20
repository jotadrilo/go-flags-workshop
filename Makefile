EXAMPLES=\
	commandline-1\
	commandline-2\
	commandline-3\
	commandline-4\
	custom-1\
	custom-2\
	custom-3\
	custom-4\
	cobra-1\
	cobra-2\
	cobra-3

CAT := $(shell command -v bat 2>/dev/null || command -v cat)

all: $(EXAMPLES)

clean:
	@rm -rf bin

%: examples/%/main.go
	@echo "go build -o bin/$@ $<"
	@go build -o bin/$@ $<

show-%: examples/%/main.go
	@$(CAT) $<
