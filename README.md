# CNAB to OCI

The intent of CNAB to OCI is to provide a translation tool in between the two formats 

## Getting Started

To get and build the project:
```sh
$ go get github.com/docker/cnab-to-oci
$ cd $GOPATH/src/github.com/docker/cnab-to-oci
$ make
```
By now you should have a binary of the project in the `bin` folder. To run it, execute:
```sh
$ bin/cnab-to-oci --help
```

### Prerequisites

Make
Golang 1.9
Git

### Installing

For installing, make sure your `$GOPATH/bin` makes part of the `$PATH`

```sh
$ make install
```

This will build and install `cnab-to-oci` into `$GOPATH/bin` 

## Development

### Running the tests

```sh
$ make test
```

### Running the e2e tests

```sh
$ make e2e
```

### Using it as a dependency on your project

Please refer to [godoc.org]()


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Maintainers

See also the list of [maintainers](MAINTAINERS) who participated in this project.

## Contributors

See also the list of [contributors](https://github.com/docker/cnab-to-oci/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
