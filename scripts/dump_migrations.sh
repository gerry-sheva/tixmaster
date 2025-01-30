#!/bin/bash

# Directory containing SQL files
MIGRATIONS_DIR="../pkg/database/migrations"

# Output file name
OUTPUT_FILE="../pkg/database/test/migrations.sql.gz"

# Create a temporary file to store combined SQL
TMP_FILE="temp_combined.sql"

# Clear or create temporary file
> "$TMP_FILE"

# Loop through all .sql files in the migrations directory
for sql_file in "$MIGRATIONS_DIR"/*.sql
do
    if [ -f "$sql_file" ]
    then
        echo "Processing $sql_file..."

        # Add a delimiter comment to separate different files
        echo -e "\n-- File: $(basename "$sql_file")\n" >> "$TMP_FILE"

        # Append the contents of the SQL file
        cat "$sql_file" >> "$TMP_FILE"
    fi
done

# Compress the combined file
echo "Compressing files into $OUTPUT_FILE..."
gzip -c "$TMP_FILE" > "$OUTPUT_FILE"

# Clean up temporary file
rm "$TMP_FILE"

echo "Complete! All SQL files have been combined into $OUTPUT_FILE"
