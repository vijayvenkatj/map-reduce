# MapReduce

A simple MapReduce implementation in Go.

## Overview

This project provides a basic MapReduce framework where a single master node orchestrates multiple worker nodes to process tasks in parallel. The communication between the master and workers is handled via Go's `net/rpc`.

## Project Structure

- `main.go`: The entry point of the application. Handles starting either a master or a worker node based on command-line flags.
- `internal/`: Contains the core MapReduce logic (master, worker, map and reduce functions).
  - `master.go`: Master node implementation handling task distribution.
  - `workers.go`: Worker node implementation.
  - `map.go` / `reduce.go`: Map and Reduce task processing logic.
  - `user/`: Contains user-defined `MapF` and `ReduceF` functions.

## Usage

You can run the application as either a master or a worker using the `-type` flag.

### Starting the Master

To start the master node:

```bash
go run main.go -type master -id 0 -nMap 2 -nReduce 1 -port 8080
```

### Starting a Worker

To start a worker node connecting to the master:

```bash
go run main.go -type worker -id 1 -master_addr localhost:8080 -nMap 2 -nReduce 1
```

### Command-line Flags

- `-type`: Node type (`master` or `worker`). Default: `master`
- `-id`: Node identifier. Default: `0`
- `-nMap`: Number of map tasks. Default: `1`
- `-nReduce`: Number of reduce tasks. Default: `1`
- `-port`: Port number for the master to listen on. Default: `8080`
- `-master_addr`: Address of the master node (for workers). Default: `localhost:8080`
