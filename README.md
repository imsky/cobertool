# cobertool

`cobertool` is a simple utility for processing [Cobertura](https://cobertura.github.io/cobertura/) code coverage reports. It can compute code coverage from multiple reports and generate HTML coverage reports.

## Why?

The Cobertura format is one of the most common code coverage report formats across different programming language ecosystems. However, the Cobertura tools are very Java-centric and infrequently updated. It's somewhat cumbersome to merge coverage reports or generate HTML reports from Cobertura reports exported by coverage tools from different language ecosystems. `cobertool` was built to address these issues and make dealing with Cobertura reports as painless as possible.

## Compatibility

Several "standard" coverage libraries for various languages have Cobertura output built in:

* JavaScript supports Cobertura with [istanbul](https://istanbul.js.org/)
* Python supports Cobertura with [nosetest](http://nose.readthedocs.io/en/latest/) and [pytest](https://docs.pytest.org/en/latest/)
* Scala supports Cobertura with [scoverage](https://github.com/scoverage/scalac-scoverage-plugin)

There are also several coverage format conversion tools:

* Go has [gocov-xml](https://github.com/AlekSi/gocov-xml)
* Erlang has [covertool](https://github.com/idubrov/covertool)
* .NET has [OpenCoverToCoberturaConverter](https://github.com/danielpalme/OpenCoverToCoberturaConverter)
* JaCoCo has [cover2cover](https://github.com/rix0rrr/cover2cover)
* Finally, LCOV reports can be converted using [lcov-to-cobertura-xml](https://github.com/eriwen/lcov-to-cobertura-xml)

## License

MIT

## Credits

Made by [Ivan Malopinsky](http://imsky.co).
