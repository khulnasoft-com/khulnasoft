#!/bin/bash

git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
blazedock run dev/preview:configure-workspace
blazedock run dev:install-dev-utils
blazedock run dev/preview/previewctl:install
pre-commit install --install-hooks
