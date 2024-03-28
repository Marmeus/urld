# urld
URL decoder and encoder written in go

## Usage
```bash
Usage of ./urld:
  -e    URL encode data (default: Decode)
  -f string
        Read data stored in a file
  -i    Read data passed through a Linux pipe
```

## Installation

The best way to install the program is with go:

```bash
go install github.com/Marmeus/urld
```

As an alternative, you can build it manually by executing the following commands:

```bash
git clone https://github.com/Marmeus/urld.git
cd urld
go build urld.go
```
