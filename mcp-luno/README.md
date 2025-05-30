# Luno MCP Server

A [Model Context Protocol](https://modelcontextprotocol.io) (MCP) server that provides access to the Luno cryptocurrency exchange API.

This server enables integration with VS Code's Copilot and other MCP-compatible clients, providing contextual information and functionality related to the Luno cryptocurrency exchange.

## ⚠️ Beta Warning

This project is currently in **beta phase**. While we've made every effort to ensure stability and reliability, you may encounter unexpected behavior or limitations. Please use it with care and consider the following:

- This MCP server config may change without prior notice
- Performance and reliability might not be optimal
- Not all Luno API endpoints are implemented yet

We welcome feedback and bug reports to help improve the project. Please report any issues you encounter via the [GitHub issue tracker](../../issues).

[![Install in VS Code](https://img.shields.io/badge/VS_Code-Install_Luno_MCP-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=mcp-luno&inputs=%5B%7B%22id%22%3A%22luno_api_key_id%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22Luno%20API%20Key%20ID%22%2C%22password%22%3Atrue%7D%2C%7B%22id%22%3A%22luno_api_secret%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22Luno%20API%20Secret%22%2C%22password%22%3Atrue%7D%5D&config=%7B%22command%22%3A%22mcp-luno%22%2C%22args%22%3A%5B%5D%2C%22env%22%3A%7B%22LUNO_API_KEY_ID%22%3A%22%24%7Binput%3Aluno_api_key_id%7D%22%2C%22LUNO_API_SECRET%22%3A%22%24%7Binput%3Aluno_api_secret%7D%22%7D%7D)

_(Requires "building from source" (see below) for now, npm & docker installation coming soon)_

## Features

- **Resources**: Access to account balances, transaction history, and more
- **Tools**: Functionality for creating and managing orders, checking prices, and viewing transaction details
- **Security**: Secure authentication using Luno API keys
- **VS Code Integration**: Easy integration with VS Code's Copilot features

## Installation

### Prerequisites

- Go 1.24 or later
- Luno account with API key and secret

### Building from Source

1. Clone the repository:

```bash
git clone https://github.com/echarrod/mcp-luno
cd mcp-luno
```

2. Build the binary:

```bash
go build -o mcp-luno ./cmd/server
```

3. Make it available system-wide (optional):

```bash
sudo mv mcp-luno /usr/local/bin/
```

## Usage

### Setting up credentials

The server requires your Luno API key and secret. These can be obtained from your Luno account settings, see here for more info: [https://www.luno.com/en/developers](https://www.luno.com/en/developers).

Set these either through:

#### A shell file

Either set this through your shell file or terminal with:
Set the following environment variables:

```bash
export LUNO_API_KEY_ID=your_api_key_id
export LUNO_API_SECRET=your_api_secret
# Optional: Enable debug mode
export LUNO_API_DEBUG=true
```

#### An .env file

Copy the .env.example file and name it .env (this always should be .gitignored), and paste your keys in there.

Depending on your setup, you might need an additional step to load these vars for your application. E.g. https://github.com/joho/godotenv

### Running the server

#### Standard I/O mode (default)

```bash
mcp-luno
```

#### Server-Sent Events (SSE) mode

```bash
mcp-luno --transport sse --sse-address localhost:8080
```

### Command-line options

- `--transport`: Transport type (`stdio` or `sse`, default: `stdio`)
- `--sse-address`: Address for SSE transport (default: `localhost:8080`)
- `--domain`: Luno API domain (default: `api.luno.com`)
- `--log-level`: Log level (`debug`, `info`, `warn`, `error`, default: `info`)

## VS Code Integration

To integrate with VS Code, add the following to your settings.json file (or click on the badge at the top of this README):

### For stdio transport:

```json
"mcp": {
  "servers": {
    "luno": {
      "command": "mcp-luno",
      "args": [],
      "env": {
        "LUNO_API_KEY_ID": "${env:LUNO_API_KEY_ID}",
        "LUNO_API_SECRET": "${env:LUNO_API_SECRET}"
      }
    }
  }
}
```

### For SSE transport:

```json
"mcp": {
  "servers": {
    "luno": {
      "type": "sse",
      "url": "http://localhost:8080/sse"
    }
  }
}
```

## Available Tools

| Tool                | Category            | Description                                       |
| ------------------- | ------------------- | ------------------------------------------------- |
| `get_ticker`        | Market Data         | Get current ticker information for a trading pair |
| `get_order_book`    | Market Data         | Get the order book for a trading pair             |
| `list_trades`       | Market Data         | List recent trades for a currency pair            |
| `get_balances`      | Account Information | Get balances for all accounts                     |
| `create_order`      | Trading             | Create a new buy or sell order                    |
| `cancel_order`      | Trading             | Cancel an existing order                          |
| `list_orders`       | Trading             | List open orders                                  |
| `list_transactions` | Transactions        | List transactions for an account                  |
| `get_transaction`   | Transactions        | Get details of a specific transaction             |

## Examples

### Working with wallets

You can ask Copilot to show your wallet balances:

```text
What are my current wallet balances on Luno?
```

### Trading

You can ask Copilot to help you trade:

```text
Create a limit order to buy 0.001 BTC at 50000 ZAR
```

### Transaction history

You can ask Copilot to show your transaction history:

```text
Show me my recent Bitcoin transactions
```

### Market Data

You can ask Copilot to show market data:

```text
Show me recent trades for XBTZAR
```

```text
What's the latest price for Bitcoin in ZAR?
```

## Security Considerations

This tool requires API credentials that have access to your Luno account. Be cautious when using API keys, especially ones with withdrawal permissions. It's recommended to create API keys with only the permissions needed for your specific use case.

### Best Practices for API Credentials

1. **Create Limited-Permission API Keys**: Only grant the permissions absolutely necessary for your use case
2. **Never Commit Credentials to Version Control**: Ensure `.env` files are always in your `.gitignore`
3. **Rotate API Keys Regularly**: Periodically regenerate your API keys to limit the impact of potential leaks
4. **Monitor API Usage**: Regularly check your Luno account for any unauthorized activity

## License

[MIT License](LICENSE)
