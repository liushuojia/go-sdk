.PHONY:all mod git help
all: mod

mod:
	go mod tidy
	go mod download
	go mod vendor
	go mod verifymak

git:
	git add .
	git commit -m "change"
	git push

help:
	@echo "make mod - 运行 Go mod"
	@echo "make help - 查看帮助"
