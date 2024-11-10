package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Role       string `json:"role"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
	Remote     bool   `json:"remote"`
	User       struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func getUsers(url string) ([]User, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var users []User
	if err = json.Unmarshal(data, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func createUser(url, apiKey string, data User) (User, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, nil
	}
	defer res.Body.Close()
	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil

}

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, nil
	}
	defer res.Body.Close()
	var user User
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, nil
	}
	defer res.Body.Close()
	var user User
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, nil
	}
	return user, nil
}

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("DELETE", fullURL, nil)
	req.Header.Set("X-API-Key", apiKey)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return err
	}
	return nil
}

func fetchTasks(baseURL, availability string) []Issue {
	amountOfIssues := map[string]int{
		"Low":    1,
		"Medium": 3,
		"High":   5,
	}

	fullURL := fmt.Sprintf("%s?sort=estimate&limit=%d", baseURL, amountOfIssues[availability])
	return getIssues(fullURL)
}

func fetchData(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Network error %v", err)
		return 0, err
	}
	statusCode := res.StatusCode

	if statusCode != 200 {
		fmt.Errorf("non-OK HTTP status: %s", statusCode)
		return statusCode, nil
	}

	return statusCode, nil
}
