package main

import (
	"fmt"
)

func print_feed(f *RSSFeed) {
	fmt.Println("Feed Title:", f.Channel.Title)
	fmt.Println("Feed Description:", f.Channel.Description)

	for i := range f.Channel.Item {
		fmt.Println(f.Channel.Item[i].Title)
		fmt.Println(f.Channel.Item[i].Description)
	}
}
