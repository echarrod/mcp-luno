module github.com/echarrod/mcp-luno

go 1.24

replace github.com/luno/luno-go => ../../../luno-go

require (
	github.com/joho/godotenv v1.5.1
	github.com/luno/luno-go v0.0.33
	github.com/mark3labs/mcp-go v0.27.1
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
	golang.org/x/time v0.10.0 // indirect
)
