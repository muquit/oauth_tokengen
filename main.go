// oauth_tokengen is a command-line tool for obtaining oauth2 tokens
// using a local callback server.
// by Claude AI 3.5 Sonnet with instructions from muquit@muquit.com
package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"

    "golang.org/x/oauth2"
)

type Config struct {
    clientID     string
    clientSecret string
    authURL      string
    tokenURL     string
    redirectURL  string
    scopes       string
    port         int
}

func main() {
    cfg := parseFlags()

    // check environment variables if flags are not set
    if cfg.clientID == "" {
        cfg.clientID = os.Getenv("OAUTH_CLIENT_ID")
    }
    if cfg.clientSecret == "" {
        cfg.clientSecret = os.Getenv("OAUTH_CLIENT_SECRET")
    }

    // validate required configuration
    if cfg.clientID == "" || cfg.clientSecret == "" || cfg.authURL == "" || cfg.tokenURL == "" {
        fmt.Println("required configuration missing. either use flags or environment variables:")
        fmt.Println("  flags: --client-id, --client-secret, --auth-url, --token-url")
        fmt.Println("  environment: OAUTH_CLIENT_ID, OAUTH_CLIENT_SECRET")
        flag.Usage()
        os.Exit(1)
    }

    // create oauth2 config
    config := &oauth2.Config{
        ClientID:     cfg.clientID,
        ClientSecret: cfg.clientSecret,
        Endpoint: oauth2.Endpoint{
            AuthURL:  cfg.authURL,
            TokenURL: cfg.tokenURL,
        },
        RedirectURL: cfg.redirectURL,
        Scopes:      strings.Split(cfg.scopes, ","),
    }

    // generate random state
    state := "random-state" // todo: generate proper random state

    // channel to receive the token
    tokenChan := make(chan *oauth2.Token)

    // set up http server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        authURL := config.AuthCodeURL(state)
        http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
    })

    http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        if r.FormValue("state") != state {
            http.Error(w, "state mismatch", http.StatusBadRequest)
            return
        }

        code := r.FormValue("code")
        if code == "" {
            http.Error(w, "code not found", http.StatusBadRequest)
            return
        }

        token, err := config.Exchange(context.Background(), code)
        if err != nil {
            http.Error(w, "failed to exchange token", http.StatusInternalServerError)
            log.Printf("token exchange error: %v", err)
            return
        }

        fmt.Fprintf(w, "authentication successful! you can close this window.")
        tokenChan <- token
    })

    // start server
    serverAddr := fmt.Sprintf(":%d", cfg.port)
    go func() {
        if err := http.ListenAndServe(serverAddr, nil); err != nil {
            log.Fatalf("failed to start server: %v", err)
        }
    }()

    fmt.Printf("please visit http://localhost:%d to start the authentication flow\n", cfg.port)

    // wait for token
    token := <-tokenChan
    fmt.Printf("\naccess token: %s\n", token.AccessToken)
    fmt.Printf("refresh token: %s\n", token.RefreshToken)
    fmt.Printf("token type: %s\n", token.TokenType)
    fmt.Printf("expiry: %s\n", token.Expiry)
}

func parseFlags() *Config {
    cfg := &Config{}

    flag.StringVar(&cfg.clientID, "client-id", "", "oauth2 client id (or use OAUTH_CLIENT_ID env var)")
    flag.StringVar(&cfg.clientSecret, "client-secret", "", "oauth2 client secret (or use OAUTH_CLIENT_SECRET env var)")
    flag.StringVar(&cfg.authURL, "auth-url", "", "authorization url")
    flag.StringVar(&cfg.tokenURL, "token-url", "", "token url")
    flag.StringVar(&cfg.redirectURL, "redirect-url", "http://localhost:8080/callback", "redirect url")
    flag.StringVar(&cfg.scopes, "scopes", "", "comma-separated list of scopes")
    flag.IntVar(&cfg.port, "port", 8080, "local server port")

    flag.Parse()

    return cfg
}