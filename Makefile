.PHONY: run build clean

run:
	air run

build:
	go build -o ./tmp/main.exe ./cmd

clean:
	rm -rf tmp