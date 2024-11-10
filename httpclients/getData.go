package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Issue struct {
	Title    string `json:"title"`
	Estimate int    `json:"estimate"`
}

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// ?
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var issues []Issue
	if err = json.Unmarshal(data, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

func marshalAll[T any](items []T) ([][]byte, error) {
  var result [][]byte
  for _, item : range items {
	data, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
  result = append(result, data)

  }
	return result, nil
}
