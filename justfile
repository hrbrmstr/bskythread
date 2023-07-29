# This is a justfile https://github.com/casey/just

# use .env here
set dotenv-load := true

# default recipe to display help information
default:
  @just --list

# build the binary
build:
  go build

# run the server with .env values
run:
  ./bskythread

# test retrieving a live post and `/`
test:
  curl -s "http://localhost:4141/bsky/did:plc:3t37x6vfigdzzp2gjcfnzlz4/3k3ewlgjhpy2m" | jq
  curl -s "http://localhost:4141/"
