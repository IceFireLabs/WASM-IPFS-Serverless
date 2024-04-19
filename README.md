# WASM-IPFS-Serverless

WASM-IPFS-Serverless combines WebAssembly computing, IPFS storage and retrieval, and is a serverless framework for decentralized applications.

## ❤️Thanks for technical support❤️
1. [**Filecoin-Lassie**](https://github.com/filecoin-project/lassie/):Support IPFS file retrieval
2. [**Filecoin-IPLD-Go-Car**](https://github.com/ipld/go-car)：Support IPFS Car file extraction
3. [**Extism**](https://github.com/extism/extism):Support wasm plug-in mechanism
4. [**wazero**](https://github.com/tetratelabs/wazero):Support wasm virtual machine
5. [**Fiber**](https://github.com/gofiber/fiber):Support high-performance HTTP server

## Quick start

### 1. Clone the repository
   ```bash
   git clone https://github.com/BlockCraftsman/WASM-IPFS-Serverless.git
   ```

### 2. Build
   ```bash
   cd WASM-IPFS-Serverless
   make
   ```
### 3. Adjust configuration file
   ```yaml
   app-type: "WASM-IPFS-Serverless-WORKER"

   # The network model of HTTP handle ,NetPoll(gin) RAWEPOLL(fiber)
   net-model: "NETPOLL"

   # Process inflow traffic network configuration
   NetWork:
   bind-network: "TCP" #Network transport layer type: TCP | UDP 
   protocol-type: "HTTP" #Application layer network protocol：HTTP | RESP | QUIC
   bind-address: "127.0.0.1:28080" #Network listening address

   #Runtime debug option
   debug:
   enable: false
   pprof-bind-addr: "127.0.0.1:19090"

   wasm-modules-files:
   enable: false
   path:
      - "hello.wasm"

   wasm-modules-ipfs:
   enable: true
   lassie-net:
      scheme: "http"
      host: "38.45.67.159" #Filecoin Lassie daemon bind IP, This is a test address 
      port: 62156 #Filecoin Lassie daemon bind Port
   cids:
      - "QmeDsaLTc8dAfPrQ5duC4j5KqPdGbcinEo5htDqSgU8u8Z" #wasm IPFS CID
      
   ```
### 4. Load configuration and run
 ```shell
   WASM-IPFS-Serverless -c wis_worker.yaml
 ```

### 5. Testing the IPFS version of WASM serverless

```shell
$ curl -d "WASM-IPFS-ServerLess" "http://localhost:28080"

👋 Hello WASM-IPFS-ServerLess%

```

## Contributing

We welcome contributions from the community. To contribute to this project:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and commit them (`git commit -am 'Add new feature'`).
4. Push your changes to the branch (`git push origin feature/your-feature`).
5. Create a new Pull Request.


## License

This library is dual-licensed under Apache 2.0 and MIT terms.
