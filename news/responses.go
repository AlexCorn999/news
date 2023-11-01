package news

import "fmt"

type NewsResponse struct {
	TotalArticles int        `json:"totalArticles"`
	Articles      []articles `json:"articles"`
}

type articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	URL         string `json:"url"`
	Image       string `json:"image"`
	PublishedAt string `json:"publishedAt"`
}

func (n *NewsResponse) Info() string {
	var result string

	for _, n := range n.Articles {
		result += fmt.Sprintf("%s\n%s\n[IMAGE] %s\n[TIME] %s\n[URL] %s\n\n",
			n.Title, n.Description, n.Image, n.PublishedAt, n.URL)
	}

	return result
}
