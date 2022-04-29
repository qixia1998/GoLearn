package main

// 使用errgroup传播错误
import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"

	"golang.org/x/sync/errgroup"
)

const data = `line one
line two with more words
error: This is erroneous line`

func main() {
	log.Printf("Application %s starting.", "Error Detection")
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	// 对每一行发射一个goroutine
	g, _ := errgroup.WithContext(context.Background())
	for scanner.Scan() {
		row := scanner.Text()
		g.Go(func() error {
			return func(s string) error {
				if strings.Contains(s, "error:") {
					return fmt.Errorf(s)
				}
				return nil
			}(row)
		})
	}

	// 等待直到goroutine结束
	if err := g.Wait(); err != nil {
		fmt.Println("Error while waiting: " + err.Error())
	}

}
