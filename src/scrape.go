package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

func makeCall() {

	fmt.Println("Enter a link to which game you want to parse:")
	var userLink string
	fmt.Scan(&userLink)

	collector := colly.NewCollector()
	trophy := TrophyModel{}
	re := regexp.MustCompile(`<[^>]*>`)
	reImg := regexp.MustCompile(`<img.*?src="([^"]+)".*?>`)

	var trophies []TrophyModel

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("Got a response from: ", response.Request.URL)
	})
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("Blimey, an error occurred: ", err)
	})

	collector.OnHTML("div.entry-content p", func(main *colly.HTMLElement) {
		tdHTML, _ := main.DOM.Html()
		result := re.ReplaceAllString(tdHTML, "")
		resultImg := reImg.FindStringSubmatch(tdHTML)

		if main.DOM.Find("img").Length() > 0 {
			parts := strings.Split(result, "\n")
			imgUrl := strings.TrimLeft(resultImg[1], "/")
			trophy.title = parts[0]
			trophy.description = parts[1]
			trophy.trophyImage = imgUrl

			main.ForEach("img", func(i int, element *colly.HTMLElement) {
				for _, attr := range element.DOM.Nodes[0].Attr {
					if attr.Key == "alt" {
						switch element.Attr("title") {
						case "Platinum Trophy":
							trophy.trophyGrade = PLATINUM
						case "Gold Trophy":
							trophy.trophyGrade = GOLD
						case "Silver Trophy":
							trophy.trophyGrade = SILVER
						case "Bronze Trophy":
							trophy.trophyGrade = BRONZE
						}
					}
				}
			})
			trophies = append(trophies, trophy)
		}

	})
	collector.Visit(userLink)
	fmt.Println(trophies)
}

type TrophyModel struct {
	title, description, trophyImage string
	trophyGrade                     TrophyGrade
}

type TrophyGrade int64

const (
	BRONZE   TrophyGrade = 0
	SILVER   TrophyGrade = 1
	GOLD     TrophyGrade = 2
	PLATINUM TrophyGrade = 3
)

func (tg TrophyGrade) String() string {
	switch tg {
	case PLATINUM:
		return "Platinum"
	case GOLD:
		return "Gold"
	case SILVER:
		return "Silver"
	case BRONZE:
		return "Bronze"
	}
	return "unknown"
}
