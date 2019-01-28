# catxslx
> Prints an excel worksheet to the console.

It takes an excel (xslx) file as input, and prints all rows and columns
of a given worksheet to the STDOUT. This is very handy when you want
to automate stuff.

![alt text](https://github.com/derhackler/catexcel/demo.gif "Overview")

## Installation

Download from the releases folder and put into your path.

## Usage example

### Example 1: 
Print the first sheet of an Excel file to STDOUT:

```sh
./catexcel example.xlsx
```

### Example 2:
Print a sheet with a given name:

```sh
./catexcel -sheet "2nd sheet" example.xlsx
```

### Example 3: 
Given you have an Excel file with the columns 'firstname, lastname, haircolor, birthday'.
Search for everyone who is called Smith and has brown hair
Print firstname and birthday.

```sh
catexcel example.xslx | awk '/Smith\tbrown/' | cut -f 1,4
```

## Development setup
1. Make sure you have golang installed.
2. Checkout the repository.
3. Install the required dependencies:

```ps
go get github.com/360EntSecGroup-Skylar/excelize
```
xlsx
4. Building on Windows (for both Windows and Unix) from Powershell:

```ps
$Env:GOOS = "windows"; $Env:GOARCH = "amd64"; go install
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go install
```

## Release History

* 1.0.0
    * First release

## Meta

Benedikt Eckhard – [@derhackler](https://twitter.com/derhackler) – public@derhackler.at

Distributed under the MIT license. See ``LICENSE`` for more information.

[https://github.com/derhackler/catexcel](https://github.com/derhackler/catexcel)

## Contributing

1. Fork it (<https://github.com/derhackler/catexcel/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request