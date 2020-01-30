package main

import "io/ioutil"

// Page is a wiki document
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(getPageLocation(p.Title), p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	body, err := ioutil.ReadFile(getPageLocation(title))
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getPageLocation(title string) string {
	return "data/" + title + ".txt"
}
