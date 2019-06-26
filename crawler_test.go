package metacriticCrawler

import "testing"

func TestClient_GetSwitchScores(t *testing.T) {
	c := NewClient(nil)
	games, err := c.GetSwitchScores()
	if err != nil {
		t.Error(err)
	}

	t.Log(games)
}
