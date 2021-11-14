package main

import (
	"log"

	"github.com/mxschmitt/playwright-go"
)

func main() {
	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		SlowMo:   playwright.Float(800.00),
		Headless: playwright.Bool(false),
	})

	if err != nil {
		log.Fatalf("could not launch chromium: %v", err)
	}

	page, err := browser.NewPage()

	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	_, err = page.Goto("https://mangaplus.shueisha.co.jp/updates")

	if err != nil {
		log.Fatalf("could not goto: %v", err)

	}

	err = page.Fill(`[placeholder="ค้นหาตามชื่อเรื่องหรือผู้แต่ง"]`, "one piece")

	if err != nil {
		log.Fatalf("could not fill: %v", err)
	}

	err = page.Keyboard().Press("Enter")

	if err != nil {
		log.Fatalf("could not press: %v", err)
	}

	err = page.Click("text=วันพีซ")

	if err != nil {
		log.Fatalf("could not click: %v", err)
	}

	list, err := page.QuerySelectorAll("text=ตอนที่")

	if err != nil {
		log.Fatalf("could not query selector all: %v", err)
	}

	latestChapter, err := list[len(list)-1].TextContent()

	if err != nil {
		log.Fatalf("could not get text content: %v", err)
	}

	log.Println("-----------------------------")
	log.Printf("ตอนล่าสุด -> %v", latestChapter)
	log.Println("-----------------------------")
}
