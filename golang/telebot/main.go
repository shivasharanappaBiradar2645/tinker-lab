package main

import (
	"os"
"io"
	"bytes"
	"time"
	"fmt"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	godotenv.Load()
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		var (
			user = c.Sender()
			text = c.Text()
		)

		return c.Send(fmt.Sprintf("hello %s, %s",user.FirstName,text))
	})

b.Handle(tele.OnPhoto, func(c tele.Context) error {
    photo := c.Message().Photo 

    buf := &bytes.Buffer{}

    
    file := photo.MediaFile()
    rc, err := b.File(file)
    if err != nil {
        return err
    }
    defer rc.Close()

    
    _, err = io.Copy(buf, rc)
    if err != nil {
        return err
    }
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImageFromBytes(buf.Bytes())
	text, err := client.Text()
	if err != nil {
		return c.Send("failed")
	}
	c.Send(text)
    
    

    
    return b.Delete(c.Message())
})
b.Start()
}
