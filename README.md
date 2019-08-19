# Pupsniffer

See the new rescue pups added to Boulder Humane Society's website

## Usage

### Install

`go install github.com/asgaines/pupsniffer`

### CLI Usage

`pupsniffer -h`

### Make Archive

An archive is where the past scrapes are stored, used in a comparison against during new scrapes to show only the new pups.

`mkdir -p ~/.pupsniffer/kennel`

### Get Sniffed

`pupsniffer -kennel ~/.pupsniffer/kennel`

