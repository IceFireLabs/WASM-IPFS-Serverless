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
      host: "38.45.67.159" #Filecoin Lassie daemon bind IP, This is a temporarily available address. When it is unavailable, please visit to install and run daemon. (https://github.com/filecoin-project/lassie?tab=readme-ov-file#http-api)
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

## Performance Testing

```shell
$ hey -n 1000000 -c 50 -m POST \ 
-d 'wasm-ipfs-serverless' \
"http://127.0.0.1:28080"

Summary:
  Total:        25.5373 secs
  Slowest:      0.0360 secs
  Fastest:      0.0001 secs
  Average:      0.0013 secs
  Requests/sec: 39158.4411
  
  Total data:   31000000 bytes
  Size/request: 31 bytes

Response time histogram:
  0.000 [1]     |
  0.004 [995857]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.007 [3829]  |
  0.011 [205]   |
  0.014 [22]    |
  0.018 [14]    |
  0.022 [18]    |
  0.025 [4]     |
  0.029 [0]     |
  0.032 [0]     |
  0.036 [50]    |


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0004 secs
  50% in 0.0014 secs
  75% in 0.0017 secs
  90% in 0.0021 secs
  95% in 0.0024 secs
  99% in 0.0031 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0001 secs, 0.0360 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0000 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0135 secs
  resp wait:    0.0012 secs, 0.0000 secs, 0.0349 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0049 secs

Status code distribution:
  [200] 1000000 responses
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
