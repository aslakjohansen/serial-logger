TARGETS = \
	serial-logger \


all: ${TARGETS}

clean:
	rm ${TARGETS}


serial-logger: serial-logger.go
	go get github.com/tarm/serial
	go build serial-logger.go

