build: main.go
	glide install
	go test -v `glide nv`
	go get -u github.com/golang/lint/golint
	golint `glide nv`
	go fmt `glide nv`
	go vet `glide nv`
	go build -i -o ${PWD}/bin/web main.go

run: build
	foreman start
