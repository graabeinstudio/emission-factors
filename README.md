# Open source emission factors

This is a repository for calculating emission factors that can be use in sustainability reportorting etc.
Currently only contains some emission factors for Norway in 2022, but feel free to leave an issue or open a PR to add other emission factors.

## How it works

1. Install go (`brew install go` if on a mac)
2. Clone the repo
3. Enter the repo and type `go build .`
4. Then run the program with `./emission-factors <?emission-factor> --location <location> --year <year>`
