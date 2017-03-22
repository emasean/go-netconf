#!/bin/bash

cd /opt/yang/example
go build --ldflags '-extldflags "-static"'
cd /opt/yang/check_example/
go build --ldflags '-extldflags "-static"'
