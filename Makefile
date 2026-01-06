.PHONY: run build clean install

run:
	air run

build:
	go build -o ./tmp/task-cli.exe ./cmd

clean:
	@if exist ".\tmp" rmdir /s /q ".\tmp"

install: build
	@echo "moving the executable to a temporary bin folder"
	@if not exist ".\bin" mkdir ".\bin"
	@move /Y ".\tmp\task-cli.exe" ".\bin\task-cli.exe"
	@echo "Installation complete. Add the bin directory to your PATH to run the application from anywhere."