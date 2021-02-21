# HUMID - HUMan ID

Easily generate memorable and human friendly IDs, as seen on Gyfcat and Netlify.

[![GitHub Actions status](https://github.com/kscarlett/humid/workflows/go.yml/badge.svg?branch=master)](https://github.com/kscarlett/humid/actions/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kscarlett/humid)](https://goreportcard.com/report/github.com/kscarlett/humid)
[![Documentation](https://godoc.org/github.com/kscarlett/humid?status.svg)](https://godoc.org/github.com/kscarlett/humid)

Support for custom wordlists:

- Animals (default)

## Usage

First, get the package with `go get github.com/kscarlett/humid`.

After that, import it as usual and call either `humid.Generate()` for the default or preferably, use `humid.GenerateWithOptions`, feeding in a `humid.Options` struct.

You can check [the docs](https://godoc.org/github.com/kscarlett/humid) for more information.

## Adding word lists

To add a new wordlist, you first create a new file in `wordlist/` using the name of the list.

Then add a new variable, follow the example in the existing wordlists. When you add your words in here, make sure they are all lowercased. If there are multiple words in entries, leave in the spaces, as these will be replaced by the appropriate separator. If there are dashes or other separators than spaces, replace these with spaces so the final result is consistent.

To ensure the quality and compatibility of the wordlists, add a new file for the tests, using an existing one as example.

Make sure to also add the wordlist to the top of this readme file.

After these steps, the wordlist should be available. If you want to contribute the list back to this project, please take a look at [our contribution guidelines](.github/CONTRIBUTING.md) and open a pull request.

## Versioning

At the moment, this package is in development. It should be fine to use, but I can't make any guarantees about the compatibility of the API between versions.

Starting at version 1.0, there will be no breaking changes within each major release.

For more information, see the [Semantic Versioning Specification](https://semver.org/).

## Dependencies

I wanted to make this a zero dependency package because that's quite nice and useful. Eventually though, it turned out using `math/rand` provided unsatisfactory randomness even on one system, and `crypto/rand` was more complex to use and quite a bit slower. In the end, `lukechampine.com/frand` was settled on for its speed, simplicity and randomness. The goal is to keep this as the only direct dependency unless there's a very good reason to add more.

## Acknowledgements

[human-id](https://github.com/RienNeVaPlus/human-id) - The JavaScript library that inspired this.

## License

This package was created by [Kellen Scarlett](https://github.com/kscarlett) and released under the [MIT License](LICENSE).
