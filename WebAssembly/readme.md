to create the WebAssembly code execute the command
GOOS=js GOARCH=wasm go build -o main.wasm toWasm.go

to use the gemerated WebAssembly code
easiest method is to use Node.js, by running this command

GOOS = JS GOARCH=wasm go run .