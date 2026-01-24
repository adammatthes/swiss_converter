package requester

import (
	"fmt"
	"net/http"
	"regexp"
	"io"
)

func FindExchangeRates() (map[string]string, error) {
	urls := []string{
		"usd-cad",
	}

	rates := make(map[string]string)

	for _, url := range urls {
		fullURL := fmt.Sprintf("https://www.exchange-rates.org/converter/%s", url)
		response, err := http.Get(fullURL)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		pattern := regexp.MustCompile(`<span class="rate-to">(.*?)</span>`)
		match := pattern.FindStringSubmatch(string(body))[0]
		rates[url] = match
	}

	return rates, nil
}
