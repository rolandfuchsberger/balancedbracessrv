package web

// use v2 - otherwise it won't work when files are included in the binary (see: https://github.com/gobuffalo/packr/issues/106)
import "github.com/gobuffalo/packr/v2"

// set up a new box by giving it a (relative) path to a folder on disk:
var box = packr.New("my templates", "./templates")
