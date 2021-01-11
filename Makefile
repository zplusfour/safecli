build-win:
	go build -o bin/safe.exe
build-lin:
	go build -o bin/safe.elf
build-macos:
	go build -o bin/safe.app
clean:
	rm -f *.exe