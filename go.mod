module github.com/bnema/gio-shadcn

go 1.24.5

require (
	gioui.org v0.8.0
	github.com/bnema/gio-shadcn/components/button v0.0.0
	github.com/bnema/gio-shadcn/components/card v0.0.0-00010101000000-000000000000
	github.com/bnema/gio-shadcn/components/input v0.0.0-00010101000000-000000000000
	github.com/bnema/gio-shadcn/components/label v0.0.0
	github.com/bnema/gio-shadcn/components/titlebar v0.0.0-00010101000000-000000000000
	github.com/bnema/gio-shadcn/theme v0.0.0
	github.com/spf13/cobra v1.9.1
	golang.org/x/text v0.27.0
)

require (
	gioui.org/shader v1.0.8 // indirect
	github.com/bnema/gio-shadcn/utils v0.0.0 // indirect
	github.com/go-text/typesetting v0.2.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	golang.org/x/exp v0.0.0-20250620022241-b7579e27df2b // indirect
	golang.org/x/exp/shiny v0.0.0-20240707233637-46b078467d37 // indirect
	golang.org/x/image v0.18.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
)

replace github.com/bnema/gio-shadcn/components/button => ./components/button

replace github.com/bnema/gio-shadcn/components/card => ./components/card

replace github.com/bnema/gio-shadcn/components/input => ./components/input

replace github.com/bnema/gio-shadcn/components/label => ./components/label

replace github.com/bnema/gio-shadcn/components/titlebar => ./components/titlebar

replace github.com/bnema/gio-shadcn/theme => ./theme

replace github.com/bnema/gio-shadcn/utils => ./utils
