SRC=$(wildcard *.go)

all:
	@make install
	@go build -o bio $(SRC)

install:
	@cd binfo && go install

test:
	@make all
	@./bio --reverse test_files/reverse/reverse_complement_data.txt

clean:
	@rm -fr ../../../pkg/darwin_386/bioinfo/test04
	@rm -f ./bio
