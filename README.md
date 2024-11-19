# PcapRunner

**PcapRunner** is a sleek and efficient Go application designed to process and replay packet capture (pcap) files. It offers configurable options through YAML files and supports internationalized messaging for a streamlined experience.

## Features

- **Packet Capture Replay**: Reads and replays packets from `.pcap` files.
- **Customizable Configuration**: Configure the application using YAML files to specify:
  - Pcap file path.
  - Network interface for replay.
  - Language file for internationalized messages.
- **Error and Message Management**: Uses a YAML-based language file for consistent error and message handling.
- **Flexible Execution**: Command-line flags enable easy overrides of default settings.

## Prerequisites

- Go (version 1.23.2 or later)
- Installed libraries:
  - [`github.com/google/gopacket`](https://github.com/google/gopacket) for packet processing.
  - [`gopkg.in/yaml.v2`](https://pkg.go.dev/gopkg.in/yaml.v2) for YAML parsing.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/pcaprunner.git
   cd pcaprunner
   ```

2. Build the application:
   ```bash
   go build -o pcaprunner main.go
   ```

3. Ensure the required configuration files are available:
   - A `.yaml` file for application configuration.
   - A `.yaml` file for language-based messages.

## Usage

Run the application with the following command-line flags:

```bash
./pcaprunner -config <path-to-config-file> -iface <network-interface>
```

### Flags

- `-config`: Path to the YAML configuration file (required).
- `-iface`: Name of the network interface to use for replaying packets (optional, overrides config).
- `-lang`: Path to the YAML language file (optional).

### Example Configuration File

```yaml
pcapFile: "/path/to/capture.pcap"
ifaceName: "eth0"
languageFile: "/path/to/language.yaml"
```

### Example Language File

```yaml
errors:
  missingFile: "The specified file could not be found."
  invalidConfig: "The configuration file is invalid."

messages:
  startReplay: "Starting packet replay on interface {{ .Interface }}..."
  endReplay: "Packet replay completed successfully."
```

## Logging

**PcapRunner** logs all key events and errors to the console for transparency. Redirect logs as needed in production environments.

## Future Enhancements

- Real-time packet filtering and analysis.
- Web-based dashboard for monitoring packet replay.

## Contributing

We welcome contributions! Feel free to submit issues or pull requests to enhance **PcapRunner**.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.