module fronted

go 1.14

require (
    controller v0.0.0
    engine v0.0.0
    fetcher v0.0.0
    model v0.0.0
    view v0.0.0
)
replace (
    controller v0.0.0 => ./controller
    engine v0.0.0 => ../engine
    fetcher v0.0.0 => ../fetcher
    model v0.0.0 => ../model
    view v0.0.0 => ./view
)