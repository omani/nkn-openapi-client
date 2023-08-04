<p align="center"><img src="https://avatars0.githubusercontent.com/u/64492989?s=200&v=4" width="150"></p>

# nkn-openapi-client
A client for the NKN OpenAPI written in go

## Install
`go install github.com/omani/nkn-openapi-client/cmd/nkn-openapi-client@latest`


## Usage
### Clone this repository
`git clone https://github.com/omani/nkn-openapi-client`

### Show transactions of an NKN wallet address:
```
cd cmd
go run main.go transactions --address NKNJ6Tka2rcrqT4FPJTjfoWQLjvahctSiyRF
```

### Show block at a given height
```
cd cmd
go run main.go blocks --height 5648381
```

### Show block with a given hash
```
cd cmd
go run main.go blocks --hash 0d48328a5005e7455c6a5e2a0b5bc346b09fbae129f1775589be83657850656a
```

### Use your own NKN OpenAPI instance by providing an `url` flag
```
cd cmd
go run main.go --url https://my-own-nkn-openapi.tld blocks --hash 0d48328a5005e7455c6a5e2a0b5bc346b09fbae129f1775589be83657850656a
```
