# Open source emission factors

This is a repository for calculating emission factors that can be use in sustainability reportorting etc.
Currently only contains some emission factors for Norway in 2022, but feel free to leave an issue or open a PR to add other emission factors.

## How it works

1. Install go (`brew install go` if on a mac)
1. Clone the repo

### API

1. Run `docker compose up` to start API

### CLI

1. Type `CGO_ENABLED=0 go build -o emission-factors-cli cli/main.go` to build the CLI.
1. Then run the CLI with `./emission-factors-cli <?emission-factor> --location <location> --year <year>`
