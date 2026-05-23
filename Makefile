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

install:
	install -Dm755 $(BUILDDIR)/$(EXEC) $(INSTALL_DIR)/$(EXEC)

uninstall:
	rm $(INSTALL_DIR)/$(EXEC)