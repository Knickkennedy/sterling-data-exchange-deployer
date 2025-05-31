# Makefile inside docs/

VENV_DIR := .venv
PYTHON := $(VENV_DIR)/bin/python
PIP := $(VENV_DIR)/bin/pip
MKDOCS := $(VENV_DIR)/bin/mkdocs

.PHONY: help venv activate clean serve build deploy

help:
	@echo "Makefile commands:"
	@echo "  venv    - Create virtual environment and install dependencies"
	@echo "  activate- Print instructions to activate venv"
	@echo "  clean   - Remove site/ and .venv/"
	@echo "  serve   - Serve documentation locally"
	@echo "  build   - Build the documentation site"
	@echo "  deploy  - Deploy to GitHub Pages"

venv:
	@test -d "$(VENV_DIR)" || python3 -m venv $(VENV_DIR)
	@$(PIP) install --upgrade pip setuptools wheel
	@$(PIP) install -r requirements.txt
	@echo "Virtual environment created at $(VENV_DIR)"
	@echo "Activate it with: . $(VENV_DIR)/bin/activate"

activate:
	@echo "To activate the virtual environment, run:"
	@echo ". $(VENV_DIR)/bin/activate"

clean:
	@rm -rf site
	@rm -rf $(VENV_DIR)
	@echo "Cleaned site/ and virtual environment."

serve: venv
	@$(MKDOCS) serve -t material

build: venv
	@$(MKDOCS) build -t material

deploy: venv
	@$(MKDOCS) gh-deploy -t material --force
