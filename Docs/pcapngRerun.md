# Working Document: PCap Rerun Application Development

## Table of Contents

1. [Introduction](#introduction)
2. [Objectives](#objectives)
3. [Technology Stack](#technology-stack)
4. [Functional Requirements](#functional-requirements)
   - [Basic Features](#basic-features)
   - [Advanced Replay Features](#advanced-replay-features)
   - [Logging & Analysis Features](#logging--analysis-features)
   - [User-Focused Features](#user-focused-features)
   - [Additional User Experience Features](#additional-user-experience-features)
5. [Non-Functional Requirements](#non-functional-requirements)
6. [Design Considerations](#design-considerations)
7. [Development Plan](#development-plan)
8. [References](#references)

---

## Introduction

The **PCap Rerun** application is a Go-based tool designed to replay `.pcapng` files from a workstation. This document compiles all the necessary information required to build the tool, outlining its objectives, features, technology stack, and development considerations.

## Objectives

- **Primary Objective**: Develop an application capable of replaying network traffic from `.pcapng` files.
- **Secondary Objectives**:
  - Provide advanced replay and analysis features.
  - Ensure cross-platform compatibility (Linux, Windows, macOS).
  - Offer both CLI and web-based interfaces for user interaction.

## Technology Stack

- **Programming Language**: Go (Golang)
- **Libraries/Frameworks**:
  - `gopacket` for packet parsing and manipulation.
  - `pcapgo` for reading `.pcapng` files.
  - Web frameworks like `Gin` or `Echo` for the web interface.
  - TUI libraries like `tview` for interactive CLI.

## Functional Requirements

### Basic Features

1. **Load pcapng Files**
   - Ability to load and parse `.pcapng` files.
   - Validate file integrity before processing.

2. **Replaying pcapng Files**
   - Replay packets according to original timestamps.
   - Option to set custom intervals between packets.

3. **Filtering Packets**
   - Apply filters by IP address, port number, protocol type.
   - Support for complex filter expressions.

4. **Real-time Replay**
   - Minimal delay to simulate real-time network traffic.
   - Synchronize replay with actual packet capture times.

5. **Loop Replay**
   - Option to replay packets indefinitely.
   - Set a specific number of replay iterations.

6. **CLI Interface**
   - Command-line interface for configuration and execution.
   - Provide help and usage documentation.

7. **Custom Output Interfaces**
   - Send packets through a specified network interface.
   - List available interfaces for user selection.

8. **Custom Replay Speed**
   - Adjust replay speed (faster, slower, real-time).
   - Speed modifiers (e.g., 2x, 0.5x).

### Advanced Replay Features

1. **Protocol-specific Parsing**
   - Deep parsing for protocols like HTTP, DNS, FTP.
   - Display protocol-specific information during replay.

2. **Interactive Mode**
   - Text-based or graphical UI for monitoring replay.
   - Control replay operations in real-time.

3. **Multi-threading Support**
   - Utilize multiple CPU cores for performance.
   - Efficient handling of large `.pcapng` files.

4. **Packet Modification**
   - Modify packet headers (IP, TCP, UDP) on-the-fly.
   - Recalculate checksums after modifications.

### Logging & Analysis Features

1. **Replay Statistics**
   - Real-time statistics: packets sent, bandwidth usage, errors.
   - Display retransmissions and dropped packets.

2. **Log All Replays**
   - Detailed logging of replayed traffic.
   - Support for different log formats.

3. **Replay Results Comparison**
   - Compare replayed traffic with original capture.
   - Highlight discrepancies and anomalies.

4. **Packet Summary Output**
   - Export packet summaries in JSON, CSV formats.
   - Include key metadata for analysis.

5. **Packet Loss Detection**
   - Detect and report any packet loss during replay.
   - Provide insights into potential causes.

6. **Replay Visualization**
   - Graphical charts for bandwidth and packet flow over time.
   - Tools like `gocui` or external visualization libraries.

### User-Focused Features

1. **Configuration File Support**
   - Load and save replay settings in YAML or JSON.
   - Facilitate repeatable replay sessions.

2. **Cross-platform Support**
   - Ensure compatibility with major operating systems.
   - Platform-specific instructions and optimizations.

3. **Web Interface**
   - Lightweight web UI for remote control.
   - Secure access with authentication mechanisms.

4. **Packet Marking/Tagging**
   - Mark specific packets for emphasis.
   - Custom tags for grouping and identification.

5. **Integration with Third-party Tools**
   - APIs for interfacing with network analysis tools.
   - Export data in formats compatible with tools like Wireshark.

### Additional User Experience Features

1. **Error Handling and Reporting**
   - Comprehensive error messages.
   - Log files for debugging and support.

2. **Remote Replay Control**
   - Control replay operations over the network.
   - RESTful API endpoints for automation.

3. **Replay Scheduling**
   - Schedule replays at specific times.
   - Recurring replay tasks with cron-like syntax.

4. **Replay Pause/Resume**
   - Pause replay without losing state.
   - Resume from the exact point of interruption.

## Non-Functional Requirements

- **Performance**
  - Efficient memory and CPU utilization.
  - Capable of handling high-throughput scenarios.

- **Security**
  - Secure handling of network interfaces.
  - Authentication for web and remote features.

- **Usability**
  - Intuitive interfaces (CLI, TUI, Web).
  - Comprehensive documentation and help guides.

- **Reliability**
  - Robust error recovery mechanisms.
  - Consistent performance under varying conditions.

- **Scalability**
  - Ability to handle large `.pcapng` files.
  - Modular design for future feature expansions.

## Design Considerations

- **Modular Architecture**
  - Separate components for parsing, replaying, and UI.
  - Ease of maintenance and testing.

- **Thread Safety**
  - Proper synchronization when using concurrency.
  - Avoid race conditions and deadlocks.

- **Cross-platform Network Access**
  - Abstract network interface interactions.
  - Handle OS-specific network API differences.

- **User Permissions**
  - Address the need for elevated permissions.
  - Provide guidelines for running with necessary rights.

## Development Plan

1. **Planning and Requirements Gathering**
   - Finalize feature set and priorities.
   - Identify potential risks and mitigation strategies.

2. **Prototype Development**
   - Create a basic version to load and replay `.pcapng` files.
   - Validate core functionalities.

3. **Incremental Feature Implementation**
   - Implement basic features first.
   - Gradually add advanced features and enhancements.

4. **Testing Strategy**
   - Unit tests for individual components.
   - Integration tests for combined features.
   - Performance testing with large data sets.

5. **User Interface Development**
   - Develop CLI commands and options.
   - Build the web interface with essential controls.

6. **Documentation**
   - Write user manuals and API documentation.
   - Include examples and best practices.

7. **Deployment and Distribution**
   - Prepare build scripts for different platforms.
   - Package the application for easy installation.

8. **Maintenance and Support**
   - Establish a process for handling bug reports.
   - Plan for regular updates and feature additions.

## References

- **Go Programming Language Documentation**
  - [https://golang.org/doc/](https://golang.org/doc/)
- **Gopacket Library**
  - [https://github.com/google/gopacket](https://github.com/google/gopacket)
- **Pcapgo Library**
  - [https://pkg.go.dev/github.com/google/gopacket/pcapgo](https://pkg.go.dev/github.com/google/gopacket/pcapgo)
- **Go Concurrency Patterns**
  - [https://golang.org/doc/effective_go.html#concurrency](https://golang.org/doc/effective_go.html#concurrency)
- **Network Programming with Go**
  - Relevant articles and books for in-depth understanding.

---

*This document serves as a comprehensive guide to developing the PCap Rerun application. It should be reviewed and updated regularly throughout the development process to ensure alignment with project goals and user needs.*