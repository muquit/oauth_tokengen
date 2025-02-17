# oauth_tokengen

`oauth_tokengen` is a command line tool to obtain OAuth2 tokens for different providers using a local callback server. It is written in [golang](https://golang.org/) and uses [oauth2](https://pkg.go.dev/golang.org/x/oauth2) package for OAuth2 implementation.

All the heavy lifting is done by the package [oauth2](https://pkg.go.dev/golang.org/x/oauth2).

If you have any question, request or suggestion, please enter it in the [Issues](https://github.com/muquit/oauth_tokengen/issues) with appropriate label.

## Quick Start

```bash
go install github.com/muquit/oauth_tokengen@latest
```

Or build from source:
```bash
git clone https://github.com/muquit/oauth_tokengen.git
cd oauth_tokengen
go build
```

## Usage Examples

### Yahoo Mail
```bash
# Using flags
./oauth_tokengen \
  --client-id="your-client-id" \
  --client-secret="your-client-secret" \
  --auth-url="https://api.login.yahoo.com/oauth2/request_auth" \
  --token-url="https://api.login.yahoo.com/oauth2/get_token" \
  --scopes="mail-w"

# Using environment variables
export OAUTH_CLIENT_ID="your-client-id"
export OAUTH_CLIENT_SECRET="your-client-secret"
./oauth_tokengen \
  --auth-url="https://api.login.yahoo.com/oauth2/request_auth" \
  --token-url="https://api.login.yahoo.com/oauth2/get_token" \
  --scopes="mail-w"
```

### Google Mail
```bash
# Using flags
./oauth_tokengen \
  --client-id="your-client-id" \
  --client-secret="your-client-secret" \
  --auth-url="https://accounts.google.com/o/oauth2/v2/auth" \
  --token-url="https://oauth2.googleapis.com/token" \
  --scopes="https://mail.google.com/"

# Using environment variables
export OAUTH_CLIENT_ID="your-client-id"
export OAUTH_CLIENT_SECRET="your-client-secret"
./oauth_tokengen \
  --auth-url="https://accounts.google.com/o/oauth2/v2/auth" \
  --token-url="https://oauth2.googleapis.com/token" \
  --scopes="https://mail.google.com/"
```

## Command line flags and Environment Variables

Required (can be provided via flags or environment variables):
- `--client-id` - OAuth2 client ID (or `OAUTH_CLIENT_ID` env var)
- `--client-secret` - OAuth2 client secret (or `OAUTH_CLIENT_SECRET` env var)
- `--auth-url` - Authorization URL
- `--token-url` - Token URL
- `--scopes` - Comma-separated list of scopes

Optional:
- `--redirect-url` - Redirect URL (default: "http://localhost:8080/callback")
- `--port` - Local server port (default: 8080)

## Common Provider Endpoints

### Google
- Auth URL: `https://accounts.google.com/o/oauth2/v2/auth`
- Token URL: `https://oauth2.googleapis.com/token`

### Yahoo
- Auth URL: `https://api.login.yahoo.com/oauth2/request_auth`
- Token URL: `https://api.login.yahoo.com/oauth2/get_token`

### Microsoft
- Auth URL: `https://login.microsoftonline.com/common/oauth2/v2.0/authorize`
- Token URL: `https://login.microsoftonline.com/common/oauth2/v2.0/token`

## License

[MIT License](LICENSE)

## Author

Claude AI 3.5 Sonnet with instructions from muquit@muquit.com