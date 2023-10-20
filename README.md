# gorav1e
simple go api for rav1e rust av1 encoder/decoder

# Description
This started as a simple wrapper to expriment with av1 encoding/decoding from golang.
I decided to try and wrap the rust rav1e sdk.
However it seems that directly calling the rav1e C lib created from cargo-c seems to break the golang garbage collector.
I am not sure why but I feel that the work presented here https://words.filippo.io/rustgo/ provides some clue.
Given that I was in a bit of a rush I tried using C trampoline code to call the api which does appear to work.

# Limitations
Given the above it is going to take a little more time to correctly add trampoline code to cover the entire rust-c api

It does however show the general direction in order to get started.

# Building
First get the rav1e repo and follow the build instructions.
If like me you were working with a slightly older centos version then building nasm is a pain and if you contact me I will try to help you through that part.
Build the cargo-c api as documented and install the headers and shared lib in /usr/local/include and /usr/local/lib respectively

Build the trampline code first

cc -c -I /usr/local/include/rav1e trampoline.c

This will create the object file trampoline.o

Ensre you have LD_LIBRARY_PATH set (or build a static executable from go)

export LD_LIBRARY_PATH=/usr/local/lib

run the simple golang example

go run main.go

It doesn't do much as yet except call the rav1e api but it is a start...



