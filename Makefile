SHELL := /bin/bash

build: target/cli-1.0.jar
target/cli-1.0.jar:
	mvn package

run: target/cli-1.0.jar
	java -cp ./target/cli-1.0.jar io.pwmcintyre.CLI
