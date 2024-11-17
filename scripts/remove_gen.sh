#!/bin/bash

# Проверяем, передан ли аргумент
if [ -z "$1" ]; then
  echo "Usage: $0 <directory>"
  exit 1
fi

# Директория, переданная пользователем
TARGET_DIR="$1"

# Проверяем, существует ли директория
if [ ! -d "$TARGET_DIR" ]; then
  echo "Error: Directory '$TARGET_DIR' does not exist."
  exit 1
fi

# Удаляем .gen из имен файлов в указанной директории
find "$TARGET_DIR" -type f -name "*.gen.*" | while read -r file; do
  # Новый путь к файлу без .gen
  new_name=$(echo "$file" | sed 's/\.gen//')

  # Переименование файла
  mv "$file" "$new_name"
  echo "Renamed: $file -> $new_name"
done

echo "Done renaming files in directory: $TARGET_DIR"
