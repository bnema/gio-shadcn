module github.com/bnema/gio-shadcn/components/titlebar

go 1.24.5

require (
	gioui.org v0.8.0
	github.com/bnema/gio-shadcn/components/button v0.0.0
	github.com/bnema/gio-shadcn/components/label v0.0.0
	github.com/bnema/gio-shadcn/theme v0.0.0
)

replace github.com/bnema/gio-shadcn/theme => ../../theme

replace github.com/bnema/gio-shadcn/utils => ../../utils

replace github.com/bnema/gio-shadcn/components/button => ../button

replace github.com/bnema/gio-shadcn/components/label => ../label

require (
	github.com/bnema/gio-shadcn/utils v0.0.0 // indirect
	github.com/go-text/typesetting v0.2.1 // indirect
	golang.org/x/exp v0.0.0-20250620022241-b7579e27df2b // indirect
	golang.org/x/exp/shiny v0.0.0-20240707233637-46b078467d37 // indirect
	golang.org/x/image v0.18.0 // indirect
	golang.org/x/text v0.16.0 // indirect
)
