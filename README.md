# edist

**E**dit **Dist**ance

`edist` calculate edit distance between two string

# Install

## Using Go
```shell
go install github.com/xztaityozx/edist@latest
```

## Download PreCompiled Binary

Download from pre-compiled binary from [GitHub Releases](https://github.com/xztaityozx/edist/releases)


# Usage

## Calculate `Levenshtein Distance`

```shell
$ edist levenshtein "abc" "abd"
1

$ edist levenshtein "ca" "abc"
3

# `L` subcommand is alias for levenshtein
$ edist L "キュアミラクル" "キュアマジカル"
3
```


### Options

```shell
  -h, --help               help for levenshtein
  -c, --replace-cost int   cost for replace (default 1)
```

## Calculate `Dameral-Levenshtein Distance`

```shell
$ edist damerau-levenshtein "abc" "abd"
1

# ca -> a -> ab -> abc
$ edist damerau-levenshtein "ca" "abc"
3

# unlimited damerau-levenshtein(`--limited=false`)
# ca -> ac -> abc
$ edist DL --limited=false "ca" "abc"
2

# `DL` subcommand is alias for damerau-levenshtein
$ edist DL "キュアミラクル" "キュアミラクル"
3
```

### Options
```shell
  -h, --help               help for damerau-levenshtein
  -l, --limited            apply limited damerau-levenshtein distance algorithm (default true)
  -c, --replace-cost int   cost for replace (default 1)
  -s, --swap-cost int      cost for swap (default 1)
```

## Calculate `Hamming Distance`

```shell
$ edist hamming 1011 1010
1

# `H` subcommand is alias for hamming
$ edist H 1011 1010
1
```


## Global Options

```shell
  -n, --normalize   apply normalization
```

# LICENSE

[MIT](./LICENSE)
