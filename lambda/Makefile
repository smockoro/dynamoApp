all: build zip

.PHONY: build
build:
	mkdir bin
	GOOS=linux go build -o bin/put_message put_message/main.go
	GOOS=linux go build -o bin/recipient_select recipient_select/main.go

.PHONY: zip
zip:
	mkdir zip
	zip -rj zip/put_message.zip bin/put_message
	zip -rj zip/recipient_select.zip bin/recipient_select

.PHONY: clean
clean:
	rm -rf bin
	rm -rf zip
