BUILD_ENV := CGO_ENABLED=0
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

TARGET_EXEC := pp

.PHONY: all clean setup build-linux build-osx build-windows

all: clean setup build-linux build-osx build-windows

clean:
	rm -rf build

setup:
	mkdir -p build/linux
	mkdir -p build/osx
	mkdir -p build/windows

build-linux: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o build/linux/${TARGET_EXEC}
	./build/linux/pp install
	cp ./build/linux/pp /usr/bin/pp

build-osx: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o build/osx/${TARGET_EXEC}
	./build/osx/pp install
	cp ./build/osx/pp /usr/bin/pp

build-windows: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=windows go build ${LDFLAGS} -o build/windows/${TARGET_EXEC}.exe
	./build/windows/pp install
	cp ./build/windows/pp.exe /usr/bin/pp.exe