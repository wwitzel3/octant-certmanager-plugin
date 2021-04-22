PLUGIN_NAME=octant-certificates.certmanager.k8s.io

ifdef XDG_CONFIG_HOME
	OCTANT_PLUGINSTUB_DIR ?= ${XDG_CONFIG_HOME}/octant/plugins
# Determine in on windows
else ifeq ($(OS),Windows_NT)
	OCTANT_PLUGINSTUB_DIR ?= ${LOCALAPPDATA}/octant/plugins
else
	OCTANT_PLUGINSTUB_DIR ?= ${HOME}/.config/octant/plugins
endif

build:
	@go build -o $(PLUGIN_NAME)

install: build
	@echo Installing to $(OCTANT_PLUGINSTUB_DIR)/$(PLUGIN_NAME)
	@mkdir -p $(OCTANT_PLUGINSTUB_DIR)
	@cp $(PLUGIN_NAME) $(OCTANT_PLUGINSTUB_DIR)/$(PLUGIN_NAME)
