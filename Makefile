run:
	go run .

install_deps:
	rm -rf go.mod go.sum
	go mod init github.com/golamrabbiazad/rest-jsonp-api-golang
	go mod tidy