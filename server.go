package main

import (
	"bytes"
	"flag"
	"image"
	"image/color"
	"image/gif"

	"github.com/gin-gonic/gin"

	"1px/log"
)

var (
	logFile    = flag.String("log", "./1px.log", "log file path")
	listenHost = flag.String("host", "0.0.0.0", "listen host")
	listenPort = flag.String("port", "9982", "listen port")
	help = flag.Bool("h", false, "print help info")
)

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	setupLogger()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	img, err := gen1pxGif()
	if err != nil {
		panic("fail to gen 1px gif " + err.Error())
	}

	r.Any("/1px/gif", func(ctx *gin.Context) {
		ctx.Data(200, "image/gif", img)
	})

	err = r.Run(*listenHost + ":" + *listenPort)
	if err != nil {
		panic("fail to run 1px service " + err.Error())
	}
}

func gen1pxGif() ([]byte, error) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{R: 255, A: 255})

	var buf bytes.Buffer
	err := gif.Encode(&buf, img, nil)
	if err != nil {
		return nil ,err
	}

	return buf.Bytes(), nil
}

func setupLogger() {
	opt := log.Options{
		Stdout:      false,
		ConsoleMode: false,
		Filename:    *logFile,
		MaxSize:     500, // 500MB
		MaxAge:      30,  // 30 days
		MaxBackups:  20,  // 20 logs
		Level:       log.InfoLevel,
	}
	log.SetOptions(opt)
}
