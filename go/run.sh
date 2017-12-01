#!/bin/bash
# 547021
#
# Go source-code:
# * Local packages inside ./src
# * Local binary inside ./bin
# * Global dependencies inside the global workspace ($GOPATH)


# ------------------- [Variables] -------------------
# Set "GOPATH" to help Go to find our local package
OLD_GOPATH=$GOPATH
export GOPATH=$PWD
export GOBIN=$PWD/bin
GO_MAIN=src/shima.go
GO_DEPS=(
	"github.com/pilu/fresh"
)

MSG_HEADER="\e[1;39m"
MSG_DIM="\e[1;2m"
MSG_ERROR="\e[1;31m"
MSG_OK="\e[0;32m"
MSG_RESET="\e[0;0m"


# ------------------- [Functions] -------------------

function log {
	echo -e "${MSG_HEADER} $1 ${MSG_DIM}${2:-$GO_MAIN}${MSG_RESET}..."
}

function dev {
	log "Run (DEV mode with fresh)"
	export __GO_DEV__=1
	fresh
}

function build {
	log "Build"
	go install $GO_MAIN
}

function init {
	# Unset "GOBIN" to install global dependencies inside $GOPATH/bin
	# Undo "GOPATH" to install global dependencies inside the original global worskpace
	# It prevents dependencies to pollute the local workspace (src, pkc, bin)
	unset GOBIN
	export GOPATH=$OLD_GOPATH

	for i in "${GO_DEPS[@]}"
	do
		log "Install dependency:" $i
		go get $i
	done
}

function doc {
	log "Doc"
	log "See" "http://localhost:6060/pkg/"
	godoc -http=:6060 -goroot="."
	#godoc -http=:6060
}

function test {
	log "Test" "config ga svg util"
	go test config ga svg util
}

function format {
	log "Format"
	gofmt -l .
}

function clean {
	log "Clean"
	go clean -x .
	go clean -i .
}

function env {
	go env
}

# Exit (SUCCESS)
function quit {
	echo -e ""
	echo -e "${MSG_OK} Done !"
	echo -e "${MSG_RESET}"
	exit 0
}

# Exit (FAILURE)
function die {
	echo -e ""
	echo -e "${MSG_ERROR} Script aborted !"
	echo -e "${MSG_RESET}"
	exit 1
}


# -------------------- [Options] --------------------

# Read command line options
while getopts "rbidtfceh" Option; do
	case $Option in
		r) dev;;
		b) build;;
		i) init;;
		d) doc;;
		t) test;;
		f) format;;
		c) clean;;
		e) env;;
		h)
			echo "  `basename $0` - Manage Go sources"
			echo "  "
			echo "  Usage: ./`basename $0` options (-rbidtfceh)"
			echo "  "
			echo "  Arguments"
			echo "     -r (default)   Run the project in DEV/WATCH mode"
			echo "     -i             Initialize the project (install dependencies)"
			echo "     -b             Build and install sources"
			echo "     -d             Create the documentation"
			echo "     -t             Launch tests"
			echo "     -f             Format sources (output only)"
			echo "     -c             Clean object files"
			echo "     -e             List Go configuration"
			echo "  "
			exit 0
			;;
		*) die;;
	esac
done


# --------------------- [Main] ----------------------

# No command line options => launch "dev" function by default
if [ "$#" -eq "0" ]; then
	dev
fi

quit
