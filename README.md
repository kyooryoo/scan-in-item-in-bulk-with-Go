# Scan in Item in Bulk with Go
Implement Alma Scan in API with GoLang

## This is a command line tool with following args:
- library, "l", "MAIN", "Specify library code. Default is MAIN."
- circ_desk, "p", "DEFAULT_CIRC_DESK", "Specify circulation desk code. Default is DEFAULT_CIRC_DESK."
- in_house_use, "i", "false", "Specify if register in house use. Default is false."
- barcode, "b", "NOTEXIST", "Specify the barcode to scan. Default is NOTEXIST."
- file, "f", "./barcodes.txt", "Specify the file path for barcode list. Default is ./barcodes.txt"

## Examples in Dev mode
* Use barcodes.txt with library name specified
`% APIKEY="<YOURKEY>" go run main.go -l LYNN`
* Scan in one specified barcode
`% APIKEY="<YOURKEY>" go run main.go -l LYNN -b SUST0000174622`

## Examples in Prod mode
* Use barcodes.txt with library name specified
`% APIKEY="<YOURKEY>" go run main.go -l LYNN`
* Scan in one specified barcode
`% APIKEY="<YOURKEY>" ./main -l LYNN -b SUST0000174622`

## Performance in scan 10 items
* Dev Mode: about 22 seconds
* Prod Mode: about 7 seconds

Current executable *main* is built on MacOS 10.15.7.
Run `go build main.go` locally for your own executable. 