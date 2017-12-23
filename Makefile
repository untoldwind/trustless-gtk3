
VERSION ?= $(shell date -u +%Y%m%d.%H%M%S)

#all: export GOPATH=${PWD}/../../../..
all: format
	@mkdir -p bin
	@echo "--> Running go build ${VERSION}"
	@go build -ldflags '-s -w' -v -i -o bin/trustless-gtk3 github.com/untoldwind/trustless-gtk3

#format: export GOPATH=${PWD}/../../../..
format:
	@echo "--> Running go fmt"
	@go fmt ./...

install.local: export GOPATH=${PWD}/../../../..
install.local: all
	@cp bin/trustless-gtk3 ${HOME}/bin
	@sed 's:@@@HOME@@@:'"${HOME}"':g' scripts/trustless-gtk3.desktop > ${HOME}/.local/share/applications/trustless-gtk3.desktop

cross: bin.linux64 dist.windows64

bin.linux64: export GOPATH=${PWD}/../../../..
bin.linux64: export GOOS=linux
bin.linux64: export GOARCH=amd64
bin.linux64: export CGO_ENABLED=1
bin.linux64:
	@mkdir -p bin
	@echo "--> Running go build ${VERSION}"
	@go build -ldflags '-s -w' -v -i -o bin/trustless-gtk3 github.com/untoldwind/trustless-gtk3

bin.windows64: export GOPATH=${PWD}/../../../..
bin.windows64: export GOOS=windows
bin.windows64: export GOARCH=amd64
bin.windows64: export CGO_ENABLED=1
bin.windows64: export CC=x86_64-w64-mingw32-gcc
bin.windows64: export PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/lib/pkgconfig
bin.windows64:
	@mkdir -p bin
	@echo "--> Running go build ${VERSION}"
	@go build -ldflags '-s -w -H=windowsgui' -v -o bin/trustless-gtk3-windows-amd64.exe github.com/untoldwind/trustless-gtk3

dist.windows64: bin.windows64
dist.windows64: dist/windows64/libatk-1.0-0.dll
dist.windows64: dist/windows64/libbz2-1.dll
dist.windows64: dist/windows64/libcairo-2.dll
dist.windows64: dist/windows64/libcairo-gobject-2.dll
dist.windows64: dist/windows64/libepoxy-0.dll
dist.windows64: dist/windows64/libexpat-1.dll
dist.windows64: dist/windows64/libffi-6.dll
dist.windows64: dist/windows64/libfontconfig-1.dll
dist.windows64: dist/windows64/libfreetype-6.dll
dist.windows64: dist/windows64/libgdk-3-0.dll
dist.windows64: dist/windows64/libgdk_pixbuf-2.0-0.dll
dist.windows64: dist/windows64/libgcc_s_seh-1.dll
dist.windows64: dist/windows64/libgio-2.0-0.dll
dist.windows64: dist/windows64/libglib-2.0-0.dll
dist.windows64: dist/windows64/libgmodule-2.0-0.dll
dist.windows64: dist/windows64/libgtk-3-0.dll
dist.windows64: dist/windows64/libgobject-2.0-0.dll
dist.windows64: dist/windows64/libgraphite2.dll
dist.windows64: dist/windows64/libharfbuzz-0.dll
dist.windows64: dist/windows64/libiconv-2.dll
dist.windows64: dist/windows64/libintl-8.dll
dist.windows64: dist/windows64/libjasper.dll
dist.windows64: dist/windows64/libjpeg-8.dll
dist.windows64: dist/windows64/libpango-1.0-0.dll
dist.windows64: dist/windows64/libpangocairo-1.0-0.dll
dist.windows64: dist/windows64/libpangoft2-1.0-0.dll
dist.windows64: dist/windows64/libpangowin32-1.0-0.dll
dist.windows64: dist/windows64/libpixman-1-0.dll
dist.windows64: dist/windows64/libpng16-16.dll
dist.windows64: dist/windows64/libpcre-1.dll
dist.windows64: dist/windows64/libstdc++-6.dll
dist.windows64: dist/windows64/libwinpthread-1.dll
dist.windows64: dist/windows64/zlib1.dll
	@mkdir -p dist/windows64/share
	@mkdir -p dist/windows64/etc/gtk-3.0
	@cp bin/trustless-gtk3-windows-amd64.exe dist/windows64/trustless-gtk3.exe
	@cp -r /usr/x86_64-w64-mingw32/share/icons dist/windows64/share
	@cp scripts/windows/settings.ini dist/windows64/etc/gtk-3.0
	@cd dist/windows64; zip -r ../../bin/trustless-gtk3-windows.zip *

dist/windows64/%.dll: /usr/x86_64-w64-mingw32/bin/%.dll
	@mkdir -p dist/windows64
	@cp $< $@

dep.install:
	@echo "-> dep install"
	@go get github.com/golang/dep/cmd/dep
	@go build -v -o bin/dep github.com/golang/dep/cmd/dep

dep.ensure: dep.install
	@bin/dep ensure
	@bin/dep prune

release:
	@echo "--> github-release"
	@go get github.com/c4milo/github-release
	@go build -v -o bin/github-release github.com/c4milo/github-release
	@bin/github-release untoldwind/trustless-gtk3 ${VERSION} master ${VERSION} 'bin/trustless-*'
