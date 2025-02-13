#!/usr/bin/env bash

# Traverse through all directories (recursively)
find . -type d | while read dir; do
    # Check if the directory is empty
    if [ "$(ls -A "$dir")" ]; then
        continue  # Skip non-empty directories
    fi

    # Extract the directory name and create a corresponding Go file
    dir_name=$(basename "$dir")
    go_file="$dir/$dir_name.go"

    # Create the Go file with a message
    echo "// This file needs to be implemented" > "$go_file"
    echo "Created: $go_file"
done

