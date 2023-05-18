package utils

import (
	goquery "github.com/google/go-querystring/query"
)

func NormalizeRequestContent(query interface{}, body interface{}) (string, error) {
	var ret string

	if query != nil {
		// attention: do not forget url tag after struct's fields
		q, err := goquery.Values(query)
		if err != nil {
			return "", err
		}
		ret += q.Encode()
	}

	if body != nil {
		// attention: do not forget url tag after struct's fields
		q, err := goquery.Values(body)
		if err != nil {
			return "", err
		}
		ret += q.Encode()
	}

	return ret, nil
}
