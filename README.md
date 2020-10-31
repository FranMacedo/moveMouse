### Run program

```go run main.go```

### Compile into exe that does not open cmd

```go build -ldflags "-H windowsgui"```

### Compile into exe with icon (no cmd)

```fyne package -icon icon.png```