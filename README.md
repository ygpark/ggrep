# ggrep - A Grep Alternative in Golang

`ggrep` is a command-line tool written in Golang that functions as an alternative to Linux `grep`. It supports **regular expressions**, **recursive search**, **line numbers**, **color highlighting**, and **various options** similar to GNU `grep`.

## Features

- ✅ Supports **regular expressions** (default behavior)
- ✅ **Recursive search** (`-r`) for searching within directories
- ✅ **Invert match** (`-v`) to exclude matching lines
- ✅ **Count matches** (`-c`) to display the number of occurrences
- ✅ **Ignore case** (`-i`) for case-insensitive search
- ✅ **Show line numbers** (`-n`) for better readability
- ✅ **Color highlighting** for matched patterns

## Installation

To install `ggrep`, make sure you have Go installed and run:

```sh
# Clone the repository
git clone https://github.com/your-repo/ggrep.git
cd ggrep

# Build the binary
go build -o ggrep

# Move to a system-wide directory (optional)
sudo mv ggrep /usr/local/bin/
```

## Usage

### Basic Search

Search for a pattern in a file:

```sh
ggrep "error" example.txt
```

### Recursive Search (`-r`)

Search inside all files in a directory:

```sh
ggrep -r "error" logs/
```

### Case-Insensitive Search (`-i`)

```sh
ggrep -i "error" logs.txt
```

### Show Line Numbers (`-n`)

```sh
ggrep -n "warning" logs.txt
```

### Invert Match (`-v`)

Find lines **not** containing the pattern:

```sh
ggrep -v "DEBUG" logs.txt
```

### Count Matches (`-c`)

Count occurrences of the pattern:

```sh
ggrep -c "ERROR" logs.txt
```

### Highlight Matches (Default)

Matches are automatically highlighted in **red**.

### Search Using Regular Expressions

```sh
ggrep "[0-9]+" data.txt  # Find numbers
```

### Search for Exact Words

Use `\b` for word boundaries:

```sh
ggrep "\berror\b" logs.txt
```

## Differences Between `ggrep` and `grep`

| Feature                   | `ggrep` (Golang)          | `grep` (GNU)          |
| ------------------------- | ------------------------- | --------------------- |
| **Default regex support** | ✅ Yes (always enabled)   | ❌ No (`-E` required) |
| **Recursive search**      | ✅ `-r`                   | ✅ `-r`               |
| **Invert match**          | ✅ `-v`                   | ✅ `-v`               |
| **Count matches**         | ✅ `-c`                   | ✅ `-c`               |
| **Color highlighting**    | ✅ ANSI-coded             | ✅ `--color=auto`     |
| **Word boundary search**  | ✅ `\bword\b`             | ✅ `\bword\b`         |
| **File pattern support**  | ✅ Uses `filepath.Glob()` | ❌ Relies on shell    |

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
