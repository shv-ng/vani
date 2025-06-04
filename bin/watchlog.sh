#!/bin/bash

# ANSI color codes
COLOR_TAG='\033[1;35m'     # [Edicational LSP]
COLOR_DATE='\033[1;33m'    # 2025/05/31
COLOR_TIME='\033[0;32m'    # 20:33:25
COLOR_FILE='\033[1;37m'    # main.go
COLOR_LINE='\033[38;5;208m' # :15:
COLOR_TEXT='\033[0;36m'    # remaining text
RESET='\033[0m'

tail -F vani.log 2>&1 | while read -r line; do
    if [[ "$line" == *"file truncated"* ]]; then
        echo -e "\n\n------------------------------------------------"
        echo "File cleared"
        echo -e "------------------------------------------------\n\n"
    else
        # Extract structured components with regex
        if [[ "$line" =~ ^(\[[^]]+\])\ ([0-9]{4}/[0-9]{2}/[0-9]{2})\ ([0-9]{2}:[0-9]{2}:[0-9]{2})\ ([^:]+)(:[0-9]+:)\ (.*)$ ]]; then
    echo -e "${COLOR_TAG}${BASH_REMATCH[1]}${RESET} \
      ${COLOR_DATE}${BASH_REMATCH[2]}${RESET} \
      ${COLOR_TIME}${BASH_REMATCH[3]}${RESET} \
      ${COLOR_FILE}${BASH_REMATCH[4]}${RESET}\
      ${COLOR_LINE}${BASH_REMATCH[5]}${RESET} \
      ${COLOR_TEXT}${BASH_REMATCH[6]}${RESET}"
  else
    # fallback for non-matching lines
    echo -e "$line"
  fi
fi
done
