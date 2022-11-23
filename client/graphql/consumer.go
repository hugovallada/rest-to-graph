package graphql

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func ExecuteGraphQLCall(url, headers, queryData string) ([]byte, error) {
	jsonData, err := json.Marshal(mountGraphQLQuery(queryData))
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	if headers != "" {
		properties := splitHeaders(headers)
		for k, v := range properties {
			req.Header.Set(k, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	return data, err
}

func mountGraphQLQuery(queryBody string) map[string]string {
	return map[string]string{"query": queryBody}
}

func splitHeaders(headers string) map[string]string {
	properties := make(map[string]string)
	props := strings.Split(headers, ";")
	for _, prop := range props {
		keyValue := strings.Split(prop, "=")
		properties[keyValue[0]] = keyValue[1]
	}
	return properties
}
