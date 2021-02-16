<div align="center">
  <h1>Go HLS</h1>
  <h5>todo ^_^</h5>
</div>

### How to run?

- [x] `$ git clone https://github.com/zikwall/go-hls`

#### Native

```shell
$ go run . \
	--root-file-directory <path | default ./tmp> \
	--input-type=<[0, 1] | default 1> \
	--tcp-port=<int | default 1339>
```

#### Docker

```shell
$ docker run -d -p 1338:1338 \
    -e ROOT_FILE_DIRECTORY='<path | default ./tmp>' \
	-e INPUT_TYPE='<[0, 1] | default 1>' \
	-e TCP_PORT=<int | default 1339> \
    --name go-hls-example qwx1337/go-hls:latest
```