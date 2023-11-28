SHELL := /bin/bash

build: target/app-1.0.jar
target/app-1.0.jar:
	mvn package

run: target/app-1.0.jar
	java -cp ./target/app-1.0.jar com.mycompany.app.App
