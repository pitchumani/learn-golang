package newspaper

type Newspaper struct {
	headlines chan string
}

// TopHeadlines returns a read-only channel of strings
func (n Newspaper) TopHeadlines() <-chan string {
	return n.headlines
}

// ReportStory returns a write-only channel of strings
func (n Newspaper) ReportStory() chan<- string {
	return n.headlines
}

func main()
