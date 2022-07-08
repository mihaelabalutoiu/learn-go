package main

import (
	"fmt"
)

func main() {
	var publisher string
	var writer string
	var artist string
	var title string
	var year int
	var pageNumber int
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."

	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "publisher by", publisher, "in", year, "with", pageNumber, "page numbers", "of grade", grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc.2"

	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "publisher by", publisher, "in", year, "with", pageNumber, "page numbers", "of grade", grade)
}
