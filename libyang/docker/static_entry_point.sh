#!/bin/bash

cd /opt/yang/example
go build -a --ldflags '-extldflags " -ldl -static"'
cd /opt/yang/xpath
go build -a --ldflags '-extldflags " -ldl -static"'
