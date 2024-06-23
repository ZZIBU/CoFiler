## `env.toml` example

[ServerInfo]
port = ":8080"
info = "test"
development = true

[MongoDB]
uri = "mongodb+srv://{ID}:{PASSWORD}@..."
db = "core"

[Storage]
type = "local"
basepath = "./data"

[Logging]
level = -1
encoding = "console"