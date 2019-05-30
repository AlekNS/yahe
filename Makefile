go_src				:= $(shell find . -type f -name '*.go' -not -path 'vendor/*')
lintfolders			:= cmd internal pkg
cmd 				?= yahe-auth
go_bin				:= ${GOPATH}/bin

include Makefile.inc
