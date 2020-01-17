run:
	env GOOS=js GOARCH=wasm go build -o static/main.wasm main.go
	# if bash then GOOS=js GOARCH=wasm go build -o static/main.wasm main.go
	go run server.go
