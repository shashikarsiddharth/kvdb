## Simple in-memory kvdb

### Assumptions:
- Commands are given in uppercase ex: `SET q 10`

### How to run binary:

```shell
mkdir kvdb
cd kvdb
git clone git@github.com:shashikarsiddharth/kvdb.git
go build .
./kvdb
```

### How to run test:
```shell
go test -v
```
