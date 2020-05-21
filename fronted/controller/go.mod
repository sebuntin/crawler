module controller

go 1.14

require (
	github.com/olivere/elastic/v7 v7.0.15
	view v0.0.0
)

replace (
	engine v0.0.0 => ../../engine
	fetcher v0.0.0 => ../../fetcher
	model v0.0.0 => ../../model
	view v0.0.0 => ../view
)
