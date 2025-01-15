# Program name
BINARY_NAME=notes-cli

# Installation directory
INSTALL_DIR=/usr/local/bin

# Source directory
SRC_DIR=.

# GO compiler
GO=go

# Install parameters
INSTALL_CMD=install

# Compilation command
build:
	$(GO) build
	@echo "Build complete: $(BINARY_NAME)"

# Installing
install: build
	$(INSTALL_CMD) -m 755 $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	chown root:root $(INSTALL_DIR)/$(BINARY_NAME)
	chmod u+s $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Program installed and setuid root flag added."

# Deinstalling
uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Program uninstalled."
	@echo "/etc/tasks.json is left. You can delete it with root permissions, if you want."

# Clean compiled binaries
clean:
	rm -f $(BINARY_NAME)
	@echo "Clean complete."

# All : build, install and clean
all: build install clean