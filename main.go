package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/alexflint/go-arg"
	_ "github.com/joho/godotenv/autoload"
)

// we need the user, app password, and listen port to work
type AuthData struct {
	Identifier string `arg:"-i,--identifier,env:BSKY_ID" help:"bsky.social user id" placeholder:"BSKY_ID"`
	Password   string `arg:"-p,--password,env:BSKY_PW" help:"bsky.social app password" placeholder:"BSKY_PW"`
	ListenPort int    `arg:"-l,--listen-port,env:PORT" help:"port to listen on" placeholder:"PORT" default:"3000"`
}

// HTTP API endpoints we use
const (
	bskyAuthURL       = "https://bsky.social/xrpc/com.atproto.server.createSession"
	bskyGetPostThread = "https://bsky.social/xrpc/app.bsky.feed.getPostThread"
)

// build a valid AT Protocol URI/URL
func buildATURL(userDid, postId string) string {
	return fmt.Sprintf("at://%s/app.bsky.feed.post/%s", userDid, postId)
}

// given user and app password, authenticate to the bsky api and retrive a jwt
func bskyAuth(authId, password string) (string, error) {

	requestData := map[string]string{
		"identifier": authId,
		"password":   password,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(bskyAuthURL, "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth request failed with status code: %d", resp.StatusCode)
	}

	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result["accessJwt"], nil
}

// given a did and postId, along with proper app credentials, return the post thread from bsky
func getPostThread(did, postId, authId, password string) (interface{}, error) {
	atURL := buildATURL(did, postId)
	getPostThreadURL := fmt.Sprintf("%s?uri=%s", bskyGetPostThread, url.QueryEscape(atURL))
	token, err := bskyAuth(authId, password)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", getPostThreadURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Println(did, postId, resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get post thread request failed with status code: %d", resp.StatusCode)
	}

	var result interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// pull in values from cli args or env file and start the server
func main() {

	var authData AuthData
	arg.MustParse(&authData)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api.hrbrmstr.dev"))
	})

	http.HandleFunc("/bsky/", func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) != 4 {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		did, postId := pathParts[2], pathParts[3]
		resp, err := getPostThread(did, postId, authData.Identifier, authData.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Success",
			"data":    resp,
		})
	})

	server := &http.Server{Addr: fmt.Sprintf(":%d", authData.ListenPort)}

	go func() {
		<-interrupt
		log.Println("Shutting down the server...")

		if err := server.Shutdown(context.TODO()); err != nil {
			log.Fatal(err)
		}

	}()

	log.Println("Server started on port", authData.ListenPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
