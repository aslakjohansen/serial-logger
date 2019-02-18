TARGETS = \
	serial-logger \
	serial-logger-win64.exe \
	serial-logger-win32.exe \


all: ${TARGETS}

clean:
	touch ${TARGETS}
	rm    ${TARGETS}

mrproper: clean
	touch dummy~
	rm        *~


serial-logger: serial-logger.go
	go get github.com/tarm/serial
	go build serial-logger.go

serial-logger-win64.exe: serial-logger.go
	go get github.com/tarm/serial
	GOOS=windows GOARCH=amd64 go build -o serial-logger-win64.exe serial-logger.go
	
serial-logger-win32.exe: serial-logger.go
	go get github.com/tarm/serial
	GOOS=windows GOARCH=386 go build -o serial-logger-win32.exe serial-logger.go
