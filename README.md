# GoPort-Forward

GoPort-Forward is a simple TCP port forwarding tool written in Go. It reads from a JSON configuration file (`config.json`) to set up multiple source and target connections for port forwarding. If the configuration file does not exist, it will be automatically created with a default configuration.

## Features
- **Configurable Port Forwarding**: You can define multiple source and target addresses for TCP forwarding in the `config.json` file.
- **Automatic JSON Configuration Creation**: If `config.json` is not found, the program automatically creates one with a default configuration.
- **Concurrent Forwarding**: The program handles multiple forwarding tasks concurrently using Goroutines.

## Prerequisites
- Go (version 1.18 or later)
- Network permissions to listen on local ports and connect to the target addresses.

## Installation

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Build the program**:
   ```bash
   go build -o goport-forward
   ```

## Usage

1. **Run the tool**:
   ```bash
   ./goport-forward
   ```

2. On the first run, the program will automatically create a `config.json` file with default configurations:
   ```json
   {
     "source": ["127.0.0.1:445"],
     "target": ["xx.xx.xx.xx:xxx"]
   }
   ```
   You can modify this file to configure your desired source and target addresses.

3. **Modify `config.json`** to set up your own forwarding rules. The format is as follows:

   ```json
   {
     "source": ["<source-ip:source-port>", "..."],
     "target": ["<target-ip:target-port>", "..."]
   }
   ```

    - **Source**: List of IP and port pairs where the tool will listen for incoming connections.
    - **Target**: List of corresponding IP and port pairs where the tool will forward the traffic.

   **Note**: The length of the `source` and `target` arrays must match, as each source will be forwarded to the corresponding target.

4. **Start forwarding**:
   The tool will log which ports it is listening on and to which target it forwards the traffic.

## Example `config.json`

```json
{
    "source": ["127.0.0.1:445", "120.0.0.1:80"],
    "target": ["xx.xx.xx.xx:xxx", "192.168.1.1:30"]
}
```

This configuration forwards:
- Traffic from `127.0.0.1:445` to `xx.xx.xx.xx:xxx`
- Traffic from `120.0.0.1:80` to `192.168.1.1:30`

## Logging
GoPort-Forward logs all important events, including:
- When it starts listening on a source port.
- When it forwards traffic to a target.
- Any errors that occur, such as failures to connect or listen.

## Error Handling
If the number of sources and targets do not match in the configuration file, the program will exit with an error message.

## License
This project is licensed under the MIT License.

---

This README now reflects the updated name **GoPort-Forward** for your project and includes detailed usage instructions.