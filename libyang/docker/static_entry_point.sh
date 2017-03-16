#!/bin/bash

cd /opt/yang/example
go build --ldflags '-extldflags "-static"'
