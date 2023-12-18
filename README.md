# Meta

Meta is a simple python script to run shell scripts defined declaratively through a YAML configuration file.

## Installation

To install this script:

1. Clone the repository with `git clone https://github.com/juanbelieni/meta`.
1. `cd` into the cloned directory.
1. Run `./meta.py install`
1. Make sure that `.local/bin` is in your `$PATH`.

## Usage

First, you have to create a `meta.yaml` (or `meta.yml`) file inside the directory from where you will run the scripts. The syntax for this config file is very simple and an example can be found inside [`meta.yaml`](/meta.yaml).

After creating and specifying a Meta config file, just run `meta [script] [args]`.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the [MIT license](/LICENSE.txt).

