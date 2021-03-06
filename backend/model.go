package main

type category struct {
	CategoryID      int    `json:"category_id"`
	Name            string `json:"name"`
	MetaTitle       string `json:"meta_title"`
	MetaKeywords    string `json:"meta_keywords"`
	MetaDescription string `json:"meta_description"`
	PageDescription string `json:"page_description"`
	PageTitle       string `json:"page_title"`
	CategoryName    string `json:"category_name"`
	ShortName       string `json:"short_name"`
	LongName        string `json:"long_name"`
	NumChildren     int    `json:"num_children"`
}
