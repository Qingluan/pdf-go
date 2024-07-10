package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Qingluan/pdf-go"
)

func main() {
	r, err := pdf.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	pages := r.NumPage()
	var buf bytes.Buffer

	for i := 1; i <= pages; i++ {
		p := r.Page(i)
		// fmt.Println("1")
		text, err := p.GetText()
		if err != nil {
			log.Fatal(err)
		}
		imgs, err := p.GetImgs()
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(text)
		for j, img := range imgs {
			buf, err := io.ReadAll(img)
			if err != nil {
				log.Fatal(err)
			}
			os.WriteFile(fmt.Sprintf("page_%d.png", i*10+j), buf, 0644)
		}
		// fmt.Println("text:", text)
		if err != nil {
			return
		}
		buf.WriteString(text)
	}

}
