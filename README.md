# Pupsniffer

See the new rescue pups added to Boulder Humane Society's website

## Usage

### Install

`go install github.com/asgaines/pupsniffer`

### CLI Usage

`pupsniffer -h`

### Make Kennel

A kennel is an archive of past scrapes, used in a comparison against during new scrapes to show only the new pups.

`mkdir -p ~/.pupsniffer/kennel`

### Get Sniffed

`pupsniffer -kennel ~/.pupsniffer/kennel`

