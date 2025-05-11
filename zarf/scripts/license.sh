#!/bin/bash


LICENSE_HEADER='/*
Copyright 2025 Alexis Bouchez <alexbcz@proton.me>

This file is part of Vauban.

Vauban is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Vauban is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with Vauban. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing options, please contact Alexis Bouchez at alexbcz@proton.me
*/'

has_license() {
    local file="$1"
    grep -q "Copyright 2025 Alexis Bouchez" "$file"
    return $?
}

add_license() {
    local file="$1"
    if [[ -f "$file" ]]; then
        echo "Adding license header to $file"
        local tmp_file=$(mktemp)
        echo "$LICENSE_HEADER" > "$tmp_file"
        cat "$file" >> "$tmp_file"
        mv "$tmp_file" "$file"
    else
        echo "Error: File $file not found"
    fi
}

    echo "Usage: $0 [file1.go file2.go ...] OR $0 --dir directory"
    echo "Adds the AGPL license header to Go files if not already present"
    exit 1
fi

if [[ "$1" == "--dir" ]]; then
        echo "Error: Directory path not provided with --dir option"
        exit 1
    fi
    
    directory="$2"
    
    if [[ ! -d "$directory" ]]; then
        echo "Error: Directory $directory not found"
        exit 1
    fi
    
    find "$directory" -type f -name "*.go" | while read -r file; do
        if ! has_license "$file"; then
            add_license "$file"
        else
            echo "License already present in $file, skipping"
        fi
    done
else
    for file in "$@"; do
        if [[ "$file" == *.go ]]; then
            if ! has_license "$file"; then
                add_license "$file"
            else
                echo "License already present in $file, skipping"
            fi
        else
            echo "Skipping $file (not a Go file)"
        fi
    done
fi

echo "License header processing complete"