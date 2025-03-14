## Project: "demo"

### Generated with
 - Types for the network messaging: true
 - Enabled Observer (http://localhost:9911): false
 - Loggers: colored rotate
 

### Supervision Tree

Applications
 - `MyApp{}` demo/apps/myapp/myapp.go
   - `MySup{}` demo/apps/myapp/mysup.go
     - `Actor1{}` demo/apps/myapp/actor1.go
     - `Actor2{}` demo/apps/myapp/actor2.go

Messages are generated for the networking in demo/types.go
- `MyMsg1{}`
- `MyMsg2{}`


#### Used command

This project has been generated with the `ergo` tool. To install this tool, use the following command:

`$ go install ergo.services/tools/ergo@latest`

Below the command that was used to generate this project:

```$ ergo -path ./ -init demo -with-app MyApp -with-sup MyApp:MySup -with-actor MySup:Actor1 -with-actor MySup:Actor2 -with-msg MyMsg1 -with-msg MyMsg2 -with-logger colored -with-logger rotate ```
