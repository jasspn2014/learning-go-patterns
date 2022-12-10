package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry() {
	// Logic
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Separation on concern
// Which can be used Separetly by multiple other structs
// eg SaveToWeb(), SaveToFile() which takes a type and store

// Not a Good way
func (j *Journal) SavetoFile(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0664)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// Package Based
var LineSeperator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeperator)), 0664)
}

// More Better Way

type Persistance struct {
	lineSeperator string
}

func (p *Persistance) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeperator)), 0664)
}

func main() {
	j := Journal{}
	j.AddEntry("I drank coffee")
	j.AddEntry("I went to the gym")

	fmt.Println(j.String())

	/*
		All Does the Same Job but here Responsibilty of Journal Package
		Should be of Data Persistance as well. It should ideally
		be handled by some other package, various ways are shown in Way 2 & 3.
	*/

	// Way 1 -
	j.SavetoFile("sample1.txt")

	// Way 2 -
	SaveToFile(&j, "sample2.txt")

	// Way 3 -
	p := Persistance{"\n"}
	p.SaveToFile(&j, "sample3.txt")
}
