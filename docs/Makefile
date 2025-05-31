VENV_DIR := ./docs/.venv

.PHONY: help venv activate clean serve build deploy

help:
	@grep "^# help\:" Makefile | grep -v grep | sed 's/# help: //'

# help: venv   - create and install dependencies in virtual environment
venv:
	@test -d "$(VENVS_DIR)" || mkdir -p "$(VENVS_DIR)"
	@rm -Rf "$(VENV_DIR)"
	@python3 -m venv "$(VENV_DIR)"
	@$(VENV_DIR)/bin/pip install --upgrade pip setuptools wheel
	@$(VENV_DIR)/bin/pip install -r docs/requirements.txt
	@$(VENV_DIR)/bin/pip install sphinx_rtd_theme
	@echo -e "Virtual environment created.\nActivate with:\n. $(VENV_DIR)/bin/activate"

# help: activate - show how to activate venv (run manually)
activate:
	@echo -e "Activate virtual environment:\n. $(VENV_DIR)/bin/activate"

# help: clean - remove build artifacts
clean:
	@rm -rf ./site

# help: serve - serve docs locally
serve:
	@$(VENV_DIR)/bin/mkdocs serve

# help: build - build static site
build:
	@$(VENV_DIR)/bin/mkdocs build

# help: deploy - deploy site to GitHub pages
deploy:
	@$(VENV_DIR)/bin/mkdocs gh-deploy --force
