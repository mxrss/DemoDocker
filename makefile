all: installdeps build

format:
	go fmt *.go

build: format
	go build -o binary/redisping2

clean:
	rm -fr binary

installdeps:
	go get -u github.com/aws/aws-sdk-go/
	go get gopkg.in/redis.v3

run:
	chmod +x $(CURDIR)/binary/redisping2
	$(CURDIR)/binary/redisping2
