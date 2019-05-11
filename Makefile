go_src				:= $(shell find . -type f -name '*.go' -not -path 'vendor/*')
lintfolders			:= cmd internal pkg
cmd 				?= yaheapi
skip_dep			?= true
go_bin				:= ${GOPATH}/bin

include Makefile.inc
