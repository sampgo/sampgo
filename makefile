default: build

clean:
	rm -rf build

build:
	CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-shared -o build/libsampgo.so main.go