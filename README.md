<div align="center">
  <h1>Go HLS</h1>
  <h5>todo ^_^</h5>
</div>

### How to run?

- `$ git clone https://github.com/zikwall/go-hls`
- `$ go run . --root-file-directory <path | default ./tmp>`

```shell
$ docker run -d -p 1338:1338 \
    -e ROOT_FILE_DIRECTORY='<path | default ./tmp>' \
    --name go-hls-example qwx1337/go-hls:latest
```