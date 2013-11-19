SRC=$(wildcard *.go)

all:
	@make install
	@go run $(SRC) test/frequent_words_data1.txt

install:
	@cd kmer && go install

clean:
	@rm -fr ../../../pkg/darwin_386/string_test/test03
