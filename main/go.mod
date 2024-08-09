module jsonparser/main

go 1.22.5

replace jsonparser/parser => ../parser

replace jsonparser/tokenizer => ../tokenizer

require (
	jsonparser/parser v0.0.0-00010101000000-000000000000
	jsonparser/tokenizer v0.0.0-00010101000000-000000000000
)

require github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3 // indirect
