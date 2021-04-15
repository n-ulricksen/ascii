## ascii

![demo](media/ascii.gif)

### Built With

- [Go](https://golang.org/)

## Installation

### Build Using Git and Go

Clone the source code

```sh
git clone https://github.com/n-ulricksen/ascii.git
cd ascii
```

Build the binary

```sh
go build .

# you can optionally specify the output path with '-o'
go build -o /usr/local/bin/ascii .
```

## Usage

ascii will print an ascii table to the terminal.

### Examples

You may specify the number of columns in the table with the -c flag.

```sh
# print 4 columns wide
ascii -c 4

# 2 columns wide
ascii -c 2
```

## Contributing

Pull requests welcome and appreciated!

## License

[MIT](https://choosealicense.com/licenses/mit/)
