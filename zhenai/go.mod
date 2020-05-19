module zhenai

go 1.14

require (
	engine v0.0.0
	fetcher v0.0.0
	github.com/Go-zh/tools v0.0.0-20190318102628-d7df357d3e99 // indirect
	github.com/chromedp/cdproto v0.0.0-20200424080200-0de008e41fa0 // indirect
	github.com/chromedp/chromedp v0.5.3
	github.com/gobwas/ws v1.0.3 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/tools/gopls v0.4.1 // indirect
	model v0.0.0
	parser v0.0.0
)

replace (
	engine v0.0.0 => ../engine
	fetcher v0.0.0 => ../fetcher
	model v0.0.0 => ../model
	parser v0.0.0 => ./parser
)
