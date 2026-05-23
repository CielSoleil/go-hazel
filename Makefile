GOFILE      = hazel.go
EXEC        = hazel
BUILDDIR    = build
PREFIX     ?= $(HOME)/.local
INSTALL_DIR = $(PREFIX)/bin

.PHONY: install uninstall clean build

build:
	go build -o $(BUILDDIR)/$(EXEC) $(GOFILE)

clean:
	rm $(BUILDDIR)/$(EXEC)

