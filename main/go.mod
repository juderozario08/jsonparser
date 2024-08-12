module jsonparser/main

go 1.22.5

replace jsonparser/parser => ../parser

replace jsonparser/tokenizer => ../tokenizer

require jsonparser/encoder v0.0.0-00010101000000-000000000000

replace jsonparser/encoder => ../encoder
