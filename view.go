package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title       string
	Body        string
	NavLinks    []NavLink
	CurrentPage string
}

type NavLink struct {
	Name   string
	URL    string
	Active bool
}

func getNavLinks(currentPage string) []NavLink {
	return []NavLink{
		{Name: "Home", URL: "/", Active: currentPage == "Home"},
		{Name: "About", URL: "/about", Active: currentPage == "About"},
		{Name: "Contact", URL: "/contact", Active: currentPage == "Contact"},
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	currentPage := "Home"
	p := &Page{
		Title:       currentPage,
		Body:        "Hello, World!",
		NavLinks:    getNavLinks(currentPage),
		CurrentPage: currentPage,
	}
	t, _ := template.ParseFiles("view.tmpl")
	t.Execute(w, p)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	currentPage := "About"
	p := &Page{
		Title:       currentPage,
		Body:        "About Us",
		NavLinks:    getNavLinks(currentPage),
		CurrentPage: currentPage,
	}
	t, _ := template.ParseFiles("view.tmpl")
	t.Execute(w, p)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	currentPage := "Contact"
	p := &Page{
		Title:       currentPage,
		Body:        "Contact Us",
		NavLinks:    getNavLinks(currentPage),
		CurrentPage: currentPage,
	}
	t, _ := template.ParseFiles("view.tmpl")
	t.Execute(w, p)
}
