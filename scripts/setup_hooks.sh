#!/bin/bash

for file in "./hooks"/*; do
  if [ -f "$file" ]; then
    filename=$(basename "$file")
    if [[ "$filename" == *.* ]]; then
      continue
    fi

    chmod +x "$file"

    echo "Linking $filename to .git/hooks/$filename"
    ln -sfn "../../hooks/$filename" "../.git/hooks/$filename"
  fi
done