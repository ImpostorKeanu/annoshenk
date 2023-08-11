# Config File Format

**Note:** See the `examples/` directory to get a better understanding
of how config files are structured.

A config file is an array of "item" objects:

```yaml
- flag: --some-flag
  delimiter: null
  default: some-default
  value: current-value
  comments:
    - This flag does a thing.
  example_values:
    - an-example-value
    - another-example-value
  exclude: false
```

## Explanation of YAML Fields

- `flag` - Is the base of the flag as it will be appended
  to the command that will be executed.
- `delimiter` - Default: `=` - A value that will bind `value` or `default` to
  `flag` in the format of `<FLAG><DELIMITER><VALUE>`, e.g.,
  `--some-flag=current-value`.
- `default` - A safe default for the flag's value.
    - Overridden by `value`.
- `value` - A value that overrides the default.
- `comments` - An array of comments about the flag.
- `example_values` - Any extra example values.
- `exclude` - Excludes the flag from current execution.

## Field Omission

**Reminder:** See the `examples/` directory.

Any flag can be omitted from a given flag, e.g., these are valid flag configs:

```yaml
- flag: -sS
  comments:
    - Nmap SYN scan flag.
- flag: -Pn
  comments:
    - Disable ping.
```