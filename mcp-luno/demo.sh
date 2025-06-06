#!/bin/bash
# Demo script to showcase Luno MCP Server functionality

# Check if the API credentials are set
if [ -z "$LUNO_API_KEY_ID" ] || [ -z "$LUNO_API_SECRET" ]; then
  echo "Error: Luno API credentials not set"
  echo "Please set LUNO_API_KEY_ID and LUNO_API_SECRET environment variables"
  echo "Example:"
  echo "  export LUNO_API_KEY_ID=your_api_key_id"
  echo "  export LUNO_API_SECRET=your_api_secret"
  exit 1
fi

# Print header
echo "================================================"
echo "Luno MCP Server Demo"
echo "================================================"
echo ""

# Build the server if not already built
echo "Building Luno MCP Server..."
go build -o mcp-luno ./cmd/server
echo "Build complete!"
echo ""

# Start the server in the background (SSE mode)
# echo "Starting Luno MCP Server in SSE mode..."
# ./mcp-luno --transport sse --sse-address localhost:8080 --log-level debug &

# Start the server in the background
echo "Starting Luno MCP Server in stdio mode..."
./mcp-luno --log-level debug &

SERVER_PID=$!

# Wait for server to start
sleep 2
echo "Server started with PID: $SERVER_PID"
echo ""

# Print instructions
echo "The server is now running in SSE mode on localhost:8080"

# Test the list_trades tool if curl is available
if command -v curl &> /dev/null; then
  echo "Testing list_trades tool for XBTZAR pair..."
  echo ""

  curl -s -X POST -H "Content-Type: application/json" -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "callTool",
    "params": {
      "name": "list_trades",
      "arguments": {
        "pair": "XBTZAR"
      }
    }
  }' http://localhost:8080/jsonrpc | json_pp

  echo ""
  echo "Tool test complete!"
  echo ""
fi

# Wait for user to press a key
echo "Press any key to stop the server and exit..."
read -n 1 -s

# Kill the server
echo "Stopping server..."
kill $SERVER_PID
echo "Server stopped"
echo ""

echo "Thank you for trying the Luno MCP Server!"
