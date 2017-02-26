all: clean build run

clean:
	rm -fr gdwmstatusbar

build:
	go build

run:
	./gdwmstatusbar

install:
	cp gdwmstatusbar /usr/local/bin
