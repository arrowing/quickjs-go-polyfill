package polyfill

import (
	"github.com/buke/quickjs-go"
	"github.com/buke/quickjs-go-polyfill/pkg/base64"
	"github.com/buke/quickjs-go-polyfill/pkg/console"
	"github.com/buke/quickjs-go-polyfill/pkg/fetch"
	"github.com/buke/quickjs-go-polyfill/pkg/timer"
	"github.com/buke/quickjs-go-polyfill/pkg/window"
)

func InjectAll(ctx *quickjs.Context) {
	window.InjectTo(ctx)
	fetch.InjectTo(ctx)
	console.InjectTo(ctx)
	base64.InjectTo(ctx)
	timer.InjectTo(ctx)
}
