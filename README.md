# Go_grep
![GoGrepLogo.png](assets/GoGrepLogo.png)

`gogrep` is a recursive CLI search tool for regular expressions.

## Usage

```bash
gogrep [flags] <pattern> [path]
```

- `pattern`: Go regular expression
- `path`: file or directory to search (default: `.`)

### Flags

- `-i` case-insensitive match
- `-v`, `--version` print version
- `-u`, `--update` check latest release
- `-h`, `--help` show help

## Examples

```bash
gogrep "TODO"
gogrep -i "error|warning" ./internal
gogrep "^func\\s+Search" main.go
```

Output is `path:line:content` (similar to ripgrep), with ANSI highlighting unless `NO_COLOR` is set.
