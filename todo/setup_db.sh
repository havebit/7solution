#!/bin/bash
sqlite3 todo.db <<EOF
CREATE TABLE IF NOT EXISTS todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT 0
);
EOF
echo "Database todo.db created with table 'todos'."
