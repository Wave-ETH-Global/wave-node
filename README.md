# Wave Node

Wave Node is implementation of backend server for Wave network.

# Building

```
go build -o ./build/wave ./cmd/wave
```

# Running

1. Generate the config of node
```
$ ./build/wave -genconfig > config.toml
```

2. Set values in the newly created config file called `config.toml`

3. Run the node
```
$ ./build/wave -config ./config.toml
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.