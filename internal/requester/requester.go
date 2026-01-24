package requester

import (
	"fmt"
	"net/http"
	"regexp"
	"io"
	"strconv"
	"strings"
)

func FindExchangeRates() (map[string]float64, error) {
	urls := []string{
		"usd-cad",
	}

	rates := make(map[string]float64)

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
		match := pattern.FindStringSubmatch(string(body))[1]
		number := strings.Split(match, " ")[0]

		newRate, err := strconv.ParseFloat(number, 64)
		if err != nil {
			return nil, err
		}
		rates[url] = newRate
	}

	return rates, nil
}
