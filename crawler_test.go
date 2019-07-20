package metacriticCrawler

import "testing"

func TestClient_GetSwitchScores(t *testing.T) {
	games, err := GetSwitchScores()
	if err != nil {
		t.Error(err)
	}

	t.Log(games)
}
