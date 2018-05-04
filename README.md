# pseudo [![CircleCI](https://circleci.com/gh/euforia/pseudo.svg?style=svg)](https://circleci.com/gh/euforia/pseudo)
Pseudo aims to be a 'pseudo' language.  Its primary purpose is simplify and
unitize development operations.

## Language

### Types
The following core types are supported:

- number
- float
- string
- list
- map

### Functions
There are a few builtin functions with the ability to register custom functions.
A complete set of functions can be found in the [functions.go](functions.go)
file.

### Syntax
A code block is any data wrapped in `${...}`.  This may be a variable, function
or any other supported operation

#### Escape sequence
A code block within `${...}` can be escaped by prepending a `$` to it like so
`$${...}`

### Development

#### Test

```shell
make test
```

#### Build

```shell
make pseudo
```

#### Run

```shell
./pseudo -h
```
