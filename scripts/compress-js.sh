#!/bin/bash

find assets/js -type f -name "*.js" | while read -r file; do
  relative_path="${file#assets/js/}"
  dir_name=$(dirname "$relative_path")
  filename=$(basename "$file" .js)

  mkdir -p "public/js/$dir_name"

  pnpm terser "$file" -o "public/js/$dir_name/$filename.min.js" --compress --mangle
done

echo "compressed!"

