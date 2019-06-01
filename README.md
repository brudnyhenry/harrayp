[![Build Status](https://travis-ci.org/brudnyhenry/harrayp.svg?branch=master)](https://travis-ci.org/brudnyhenry/harrayp)

<img src="./img/logo.png" width="200"/>

# HarrayP

HarrayP is a cli tool for managing HP P2000 arrays

## Installation

Download latest binary from [Releases](https://github.com/brudnyhenry/harrayp/releases)

## Usage
```bash
Usage:
  harrayp [command]

Available Commands:
  get         A brief description of your command
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.harrayp.yaml)
  -h, --help            help for harrayp

Use "harrayp [command] --help" for more information about a command.
```

## Configuration
Example config file with array login details:
```yaml
url: "http://XXX"
login: "user"
password: "password"
```

Configuration file default location is $HOME/.harrayp.yaml

It can be overwritten with `--config` flag


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
