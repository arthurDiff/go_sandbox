new:
	mkdir $(name) && cd ./$(name) && go mod init $(name) && touch main.go; echo "package main" >> main.go
