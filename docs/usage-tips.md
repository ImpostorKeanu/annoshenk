# Usage Tips

## Items are just parts of the command.

Contrary to use of the term `flag`, items are just strings that are married
to form a command that is passed to [exec.Command](https://pkg.go.dev/os/exec#Command). This
means that the leading flag itself can be the binary that should be executed.

## Flag Ordering is preserved

The order in which flags are specified in the config file is preserved when
the final command is assembled.

## Modifying command outside of a config file

### Prefix flags to the final command using `--head`

Any values supplied to the `--head` flag will be blindly **prefixed**
to the final command, e.g., `--head sudo` would ensure `sudo` is executed
prior to the prescribed command.

### Suffix flags to the final command using `--tail`

Any values supplied to the `--tail` flag will be blindly **suffixed**
to the final command, e.g., `--tail 192.168.1.5` would ensure that a
target is suffixed to the `examples/nmap.yml` command, giving a clear
target to scan.