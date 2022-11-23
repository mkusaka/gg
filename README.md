# gog
go get command with url format

# install

```bash
go install github.com/mkusaka/gog@latest
```

# example

```bash
# support browser url format
gog https://github.com/mkusaka/gog

# NOTE: this command just run `go get github.com/mkusaka/gog`

# support git suffix url format
gog https://github.com/mkusaka/gog.git

# support domain/path format
gog github.com/mkusaka/gg

# support go get with flags
gog -u https://github.com/mkusaka/gog
```
