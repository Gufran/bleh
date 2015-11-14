.PHONY : build
build:
	@go-bindata -o assets.go assets/
	@go build ./
