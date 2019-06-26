package metacriticCrawler

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"strconv"
	"strings"
)

const urlPattern = "https://www.metacritic.com/browse/games/release-date/available/switch/metascore?view=condensed&page=%d"

func (c *Client) GetSwitchScores() ([]Game, error) {
	pageNum := 0
	htmlByte, err := c.sendNewRequest(fmt.Sprintf(urlPattern, pageNum))
	if err != nil {
		return nil, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(string(htmlByte)))
	if err != nil {
		return nil, err
	}

	var games []Game

	gameEles := htmlquery.Find(doc, "//li[contains(@class,'product game_product')]")
	for _, gameEle := range gameEles {
		titleEle := htmlquery.FindOne(gameEle, "//div/div[1]/a")
		titleStr := strings.TrimSpace(htmlquery.InnerText(titleEle))
		scoreEle := htmlquery.FindOne(gameEle, "//div/div[2]/div")
		score, err := strconv.Atoi(strings.TrimSpace(htmlquery.InnerText(scoreEle)))
		if err != nil {
			return nil, err
		}
		fmt.Printf("%s: %d\n", titleStr, score)
		game := Game{
			Name:  titleStr,
			Score: score,
		}
		games = append(games, game)
	}

	finalPageEle := htmlquery.FindOne(doc, "//li[contains(@class,'page last_page')]/a")
	finalPage, err := strconv.Atoi(htmlquery.InnerText(finalPageEle))
	if err != nil {
		return nil, err
	}

	for i := pageNum + 1; i < finalPage; i++ {
		htmlByte, err := c.sendNewRequest(fmt.Sprintf(urlPattern, i))
		if err != nil {
			return nil, err
		}

		doc, err := htmlquery.Parse(strings.NewReader(string(htmlByte)))
		if err != nil {
			return nil, err
		}

		gameEles := htmlquery.Find(doc, "//li[contains(@class,'product game_product')]")
		for _, gameEle := range gameEles {
			titleEle := htmlquery.FindOne(gameEle, "//div/div[1]/a")
			titleStr := strings.TrimSpace(htmlquery.InnerText(titleEle))
			scoreEle := htmlquery.FindOne(gameEle, "//div/div[2]/div")
			score, err := strconv.Atoi(strings.TrimSpace(htmlquery.InnerText(scoreEle)))
			if err != nil {
				return nil, err
			}
			fmt.Printf("%s: %d\n", titleStr, score)
			game := Game{
				Name:  titleStr,
				Score: score,
			}
			games = append(games, game)
		}
	}

	return games, nil
}
