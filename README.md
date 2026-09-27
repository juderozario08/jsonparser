# Simple JSON Parser in Go

## Overview

This is a basic JSON parser written in Go for personal use. It demonstrates how to parse JSON data into Go data structures and convert Go data structures to JSON. This project serves as a learning tool and example of JSON parsing in Go.

## Things you are expected to learn from this project
- **Hashmaps**/**Maps**
- **Stack**
- **Recursion**
- **Pointers in Golang**
- **Structs**
- **Interfaces**
- **Types**
- **Closures**
- **UserInput and Syntax validation**

## Features

- **Basic JSON Parsing**: Convert JSON strings to Go data structures.
- **JSON Encoding**: Convert Go data structures to JSON strings.
- **Learning Example**: Ideal for understanding JSON parsing in Go.

## Tokenizer Folder
- Responsible for reading through the entire JSON and creating Tokens based on the type of data being read

## Parser Folder
- Responsible for Parsing and Validating the data being sent through the JSON and turned into data structures used in Golang

## Encoder
- Responsible for taking the Golang data structures and converting them into JSON format

## Installation

Clone the repository and build the project:

```sh
git clone https://github.com/juderozario08/jsonparser.git
cd jsonparser/main
go build
go run .
```

Feel free to change the main file to any inputs you like
The project includes test cases for the parser, tokenizer, and encoder that can be run using:

```sh
# Run all tests across the entire project
go test -v ./...

# Or run specific test suites
go test -v ./parser
go test -v ./tokenizer
go test -v ./encoder
```
Happy Coding :)
