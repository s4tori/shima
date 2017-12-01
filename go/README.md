<div align="center">
	<h1>Shima server</h1>
	<p>View the documentation and demo here:<br /> https://shima-grid.com</p>
</div>


<h2 align="center">Getting Started</h2>

### Linux

```sh
# Install Go in: <directory>/go/
# Set ENV variables (.bash_aliases)
# Example (directory = "~/Go")
export PATH=~/Go/go/bin:$PATH
export PATH=~/Go/wp/bin:$PATH
export GOPATH=~/Go/wp

# In the terminal:
./run.sh -i
./run.sh
```

### Windows

```sh
#   Install cmder
#   Install Go in         : <directory>/go/
#   Create WP directory in: <directory>/wp/

# Set ENV variables
# GOPATH: <directory>/wp/
# PATH  : <directory>/wp/bin

# With cmder:
sh run.sh -i
sh run.sh
```


<h2 align="center">Usage</h2>

```sh
# Install dependencies
# * "fresh" to rebuild the go app during development
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
