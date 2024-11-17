# gRPC

1. First step is to install the protobuf-compiler. Steps available on the official website `https://grpc.io/docs/languages/go/quickstart/`. On mac, `brew install protobuf`. 
2. Second step would be to install plugins to generate go code. Links will be found on the above link itself. 
3.  In the `rpc_create_user.proto`, using the plugin for VSCode, `vscode-proto3` it will not be able to initially find `user` This is cause of the settings issue where the extension expects all the protos to be in the projects root dir. To fix this, we can copy the settings from the extension main page, and then put in the vscode's settings for proto3. 
4. After generating the proto files, there will be some error due to missing dependencies. `go mod tidy` will fix em. 
5. To test out grpc apis, we use a tool called evans `https://github.com/ktr0731/evans`. We can use the command `evans --host localhost --port 9091 -r repl` to start the evans console connecting to your running server.