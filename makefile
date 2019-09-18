BUILD=go build
BINARY_NAME=mydictionary-demo

windows_386=GOOS=windows GOARCH=386 $(BUILD) -o $(BINARY_NAME)_windows_386.exe
windows_amd64=GOOS=windows GOARCH=amd64 $(BUILD) -o $(BINARY_NAME)_windows_amd64.exe
darwin_386=GOOS=darwin GOARCH=386 $(BUILD) -o $(BINARY_NAME)_darwin_386
darwin_amd64=GOOS=darwin GOARCH=amd64 $(BUILD) -o $(BINARY_NAME)_darwin_amd64
linux_386=GOOS=linux GOARCH=386 $(BUILD) -o $(BINARY_NAME)_linux_386
linux_amd64=GOOS=linux GOARCH=amd64 $(BUILD) -o $(BINARY_NAME)_linux_amd64

default:
	$(BUILD)

all:
	$(windows_386)
	$(windows_amd64)
	$(darwin_386)
	$(darwin_amd64)
	$(linux_386)
	$(linux_amd64)

windows_386:
	$(windows_386)

windows_amd64:
	$(windows_amd64)

darwin_386:
	$(darwin_386)

darwin_amd64:
	$(darwin_amd64)

linux_386:
	$(linux_386)

linux_amd64:
	$(linux_amd64)

clean:
	rm -fr $(BINARY_NAME)
	rm -fr $(BINARY_NAME)_*
