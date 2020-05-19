module douban

go 1.14
require (
    engine v0.0.0
    fetcher v0.0.0
    doubanparser v0.0.0
    model v0.0.0
    scheduler v0.0.0
)
replace (
    engine v0.0.0 => ../engine
    fetcher v0.0.0 => ../fetcher
    doubanparser v0.0.0 => ./doubanparser
    model v0.0.0 => ../model
    scheduler v0.0.0 => ../scheduler
)