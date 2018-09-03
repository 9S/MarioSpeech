all : windows macos linux

windows :
	env GOOS=windows GOARCH=amd64 go build -o bin/MarioSpeech_Win64.exe

macos :
	env GOOS=darwin GOARCH=amd64 go build -o bin/MarioSpeech_macOS

linux :
	env GOOS=linux GOARCH=amd64 go build -o bin/MarioSpeech_linux
