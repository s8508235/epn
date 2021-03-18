# epn

A command line tool for encrypting phone number with sha256 and also check for file content format.


`${binary} help` for further information.
```
Usage:
  encrypted-phone-number [flags]
  encrypted-phone-number [command]

Available Commands:
  check       check will examine input file is a line-by-line sha256 file
  help        Help about any command

Flags:
  -h, --help                 help for encrypted-phone-number
  -i, --input_file string    Raw phone number csv file. (default "./input.csv")
      --log_level string      (default "info")
  -o, --output_file string   Encrypted phone number csv file. (default "./output.csv")

Use "encrypted-phone-number [command] --help" for more information about a command.
```
