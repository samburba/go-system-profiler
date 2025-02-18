#!/usr/bin/env bash

input=$(system_profiler -listDataTypes)

echo "$input" | while read line; do
    # Remove "SP" prefix and "DataType" suffix, then convert to lowercase and replace spaces with underscores
    formatted=$(echo "$line" | sed 's/^SP//; s/DataType$//' | tr '[:upper:]' '[:lower:]' | sed 's/[A-Z]/_\0/g' | sed 's/^_//')
    echo "$formatted"
done

