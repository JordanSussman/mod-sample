module github.com/JordanSussman/mod-sample

go 1.14

replace github.com/go-vela/types => github.com/JordanSussman/types v0.1.2-0.20200817233341-2853ce2956de

replace github.com/go-vela/compiler => github.com/JordanSussman/compiler v0.1.3-0.20200818133900-c4a5e4bebff2

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-vela/compiler v0.5.0
	github.com/go-vela/types v0.5.0
)
