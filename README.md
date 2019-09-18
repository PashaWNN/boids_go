# Boids simulation

Boids algorithm implementation in Go language, compiled to WebAssembly

[Live demo](pashawnn.github.io/boids_go/)


## Compilation

```
GOOS=js GOARCH=wasm go build -o main.wasm 
```

Compiled binary can be found in `Releases` section of GitHub repository. Note that you can't just open index.html from local filesystem. You need web-server which sets correct mime type (`mime .wasm application/wasm`) to run WebAssembly. Simpliest solution is to download Caddy server and just run:
```
caddy
```
from project root.