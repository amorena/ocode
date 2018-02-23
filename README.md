# Oracle Code Flow Demo

This is simply a set of "placeholder" functions that are called from a Java flow
func to demonstrate flow. Feel free to actually implement the funcs if you
want. ;)


## Pre-requisites

You'll need Fn installed locally.  Installation instructions 
[here](https://github.com/fnproject/fn).

## Setup

1. Clone this git respository and `cd` into the repo root folder.

2. Run `./start.sh` to start all required services (fnserver, fn ui, flow server,
flow ui).  If this is your first time starting these services it may take a
few minutes as all required Docker images are pulled.

3. Open the Flow UI in your browser: http://localhost:3002/

## Deploy and Configure the Application

Run the following:

1. `fn deploy --all --local`

2. `./configure.sh`

## Run the Function

1. `fn call ocode flow`
2. **QUICKLY** switch to Flow UI to show execution

## Shutdown

1. `./stop.sh` will stop all running service Docker containers.