.PHONY: build run

build:
	go build -o Cooking-Fun

run:
	go run .

clean:
	rm -f Cooking-Fun