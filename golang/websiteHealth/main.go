package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func checker(ctx context.Context, ch <-chan string, client *http.Client, logger *zap.Logger) {
	for {
		select {
		case <-ctx.Done():
			return

		case url := <-ch:
			res, err := client.Head(url)

			if err != nil {
				logger.Warn("Website Down: ", zap.String("url", url), zap.Error(err))
				continue
			}
			res.Body.Close()
			if res.StatusCode == http.StatusOK {
				//logger.Info("Website Up: ", zap.String("url", url))
			} else {
				logger.Info("Website down: ", zap.String("url", url), zap.Int("statusCode", res.StatusCode))
			}
		}
	}
}

func main() {
	file, err := os.Open("sites.txt")
	client := &http.Client{Timeout: 2 * time.Second}
	ch := make(chan string)
	defer close(ch)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	noW := flag.Int("worker", 3, "No of Worker")
	noT := flag.Int("delay", 10, "delay time")
	flag.Parse()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	defer file.Close()
	defer stop()
	var sites []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		sites = append(sites, line)
	}

	if err := scanner.Err(); err != nil {
		logger.Panic("error while reading file")
	}

	for i := 0; i < *noW; i++ {
		go checker(ctx, ch, client, logger)
	}

	for {
		select {
		case <-ctx.Done():
			logger.Info("Shutting Down")
			return

		default:
			for _, value := range sites {
				ch <- value
			}
			time.Sleep(time.Duration(*noT) * time.Second)

		}
	}

}
