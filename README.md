# gg
go get command with url format

# install

```bash
go install github.com/mkusaka/gg@latest
```

# example

```bash
# support browser url format
gg https://github.com/mkusaka/gg

# NOTE: this command just run `go get github.com/mkusaka/gg`

# support git suffix url format
gg https://github.com/mkusaka/gg.git

# support domain/path format
gg github.com/mkusaka/gg

# support go get with flags
gg -u https://github.com/mkusaka/gg
```