<div align="center">
	<h1>Shima server</h1>
	<p>The documentation with practical examples is available online:<br /> https://shima-grid.com</p>
</div>


<h2 align="center">🚀 Getting started</h2>

### Prerequisites

Install:
* [Go](https://go.dev/dl/) (1.18.3 or higher) 

### Mac

```sh
# With homebrew
brew install go
export PATH=~/go/bin:$PATH
go env

./run.sh -i
./run.sh
```

### Linux

```sh
# Install Go in: <directory>/go/
# Set ENV variables (.bash_aliases or ~/.zshrc)
# Example (directory = "~/Go")
export PATH=~/Go/go/bin:$PATH
export PATH=~/Go/wp/bin:$PATH
export GOPATH=~/Go/wp
go env

./run.sh -i
./run.sh
```

### Windows

```sh
# Install cmder or run a WLS2 distribution in Windows terminal
# Install Go in         : <directory>/go/
# Create WP directory in: <directory>/wp/

# Set ENV variables
# GOPATH: <directory>/wp/
# PATH  : <directory>/wp/bin
go env

# With cmder:
sh run.sh -i
sh run.sh
```


<h2 align="center">👨‍💻 Usage</h2>

```sh
# Install dependencies
# * "refresh" to rebuild the go app during development
./run.sh -i

# Run in DEV/WATCH mode
./run.sh

# Test the project
./run.sh -t

# Build and install the project
./run.sh -b

# See all commands available
./run.sh -h
```
