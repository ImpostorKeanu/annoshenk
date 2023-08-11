AnnoShenk (annotated shenks) is a command wrapper that facilitates
configuration files for any command line utility. It was inspired
by [RobotClone](https://github.com/Wh1t3Rh1n0/RobotClone/tree/main),
but sacrifices comment-based simplicity for a YAML file format.

# Usage

```bash
shenk nmap.yml --tail 192.168.1.5
```

The above example would run the `nmap` command specified in [the nmap example](examples/nmap.yml)
while suffixing `192.168.1.5` to the command as a target.

See `--help` for to better understand the `--head` and `--tail` flags.

# Further Reading

- [Examples](examples)
- [Config File Format](docs/config-file-format.md)
- [Usage Tips](docs/usage-tips.md)