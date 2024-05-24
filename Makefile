build:
	@go build -o bin/mahasiswa-api

run: build
	@./bin/mahasiswa-api
