generate:
	@cd cmselect;goyacc -o parser.go -p main select.y;cd -

dep:
	@go install golang.org/x/tools/cmd/goyacc
