go_src				:= $(shell find . -type f -name '*.go' -not -path 'vendor/*')
lintfolders			:= cmd internal pkg
cmd 				?= yahe-explorer
go_bin				:= ${GOPATH}/bin

include Makefile.inc
