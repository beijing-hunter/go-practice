package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

//此方式有点：根据逻辑cpu个数开启固定个数的goroutine,减少goroutine的创建数量及堆分配
func freqNumCPU(topic string, docs []string) int {
	var found int32

	g := runtime.NumCPU() //逻辑处理器个数
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for i := 0; i < g; i++ { //根据逻辑cpu个数开启多少个goroutine

		go func() {
			var lFound int32
			defer func() {
				atomic.AddInt32(&found, lFound)
				wg.Done()
			}()

			for doc := range ch { //goroutine从channel中取文件名称
				file := fmt.Sprintf("%s.xml", doc[:8])
				f, err := os.OpenFile(file, os.O_RDONLY, 0)
				if err != nil {
					log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
					return
				}

				data, err := ioutil.ReadAll(f)
				if err != nil {
					f.Close()
					log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
					return
				}
				f.Close()

				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
					log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
					return
				}

				for _, item := range d.Channel.Items {
					if strings.Contains(item.Title, topic) {
						lFound++
						continue
					}

					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
		}()
	}

	for _, doc := range docs { //将文件名称写入channel
		ch <- doc
	}
	close(ch)

	wg.Wait()
	return int(found)
}
