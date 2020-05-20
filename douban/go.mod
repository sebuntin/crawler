module douban

go 1.14

require (
	doubanparser v0.0.0
	engine v0.0.0
	fetcher v0.0.0
	github.com/parnurzeal/gorequest v0.2.16 // indirect
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7 // indirect
	model v0.0.0
	moul.io/http2curl v1.0.0 // indirect
	persist v0.0.0
	scheduler v0.0.0
)

replace (
	doubanparser v0.0.0 => ./doubanparser
	engine v0.0.0 => ../engine
	fetcher v0.0.0 => ../fetcher
	model v0.0.0 => ../model
	persist v0.0.0 => ../persist
	scheduler v0.0.0 => ../scheduler
)
