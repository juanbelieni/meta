# Meta

Meta is a simple go script to run shell commands defined declaratively through a configuration file.

## Installation

To install this script:

1. Clone the repository with `git clone https://github.com/juanbelieni/meta`.
1. `cd` into the cloned directory.
1. Run `go run meta.go install`
1. Make sure that `.local/bin` is in your `$PATH`.

## Usage

First, you have to create a `Metafile` (or `meta.yml`) file inside the directory from where you will run the scripts. The syntax for this config file is very simple and an example can be found inside [`Metafile`](/Metafile).

After creating and specifying a Meta config file, just run `meta [script] [args]`.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the [MIT license](/LICENSE.txt).

