PKG_NAME=homelists
DUSER=aceberg

mod:
	cd src && \
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/HomeLists && \
	go mod tidy

run:
	cd src && \
	go run .

go-build:
	cd src && \
	go build .

docker-build:
	docker build -t $(DUSER)/$(PKG_NAME) .

docker-run:
	docker rm $(PKG_NAME) || true
	docker run --name $(PKG_NAME) \
	-e "TZ=Asia/Novosibirsk" \
	-v ~/.dockerdata/$(PKG_NAME):/data/$(PKG_NAME) \
	-p 8842:8842 \
	$(DUSER)/$(PKG_NAME)