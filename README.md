# Wieserlabs FlexDDS

The present repository contains a command-line interface (CLI) to the network interface of the [Wieserlabs FlexDDS-NG][1].

## Usage

You can download the compiled binaries for your platform from the GitHub release tab.

### Convert

The `convert` tool supports you in the conversion of the output frequency and the amplitude scale to the register values, e.g.,

```shell
./convert freq-out 10e6
# Output: 0x28f5c29

./convert log-ampl-scale 0.0
# Output: 0x4000

./convert lin-ampl-scale 0.0
# Output: 0x0
```

### Control

The `control` tool allows you to configure a slot of the FlexDDS, e.g., to singletone output:

```shell
./control \
    --host=10.163.100.7 \
    --slot=2 \
    --channel=0 \
    --system-clock=1e9 \
    singletone \
        --log-amplitude=0.0 \
        --frequency=60e6

```

[1]: https://www.wieserlabs.com/products/radio-frequency-generators/WL-FlexDDS-NG