# URL Submission Tool

![Go](https://img.shields.io/badge/Go-1.18+-blue.svg)
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-green.svg)
[![License](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

A tool that submits URLs to search engines (Baidu, Bing, Google) via command line or HTTP API.

## Features

- Submit URLs to multiple search engines simultaneously
- Simple command line interface
- RESTful HTTP API
- Support for Baidu, Bing and Google
- Configurable submission parameters
- Lightweight and fast

## Installation

### Prerequisites

- Go 1.23 or higher

### From source

```bash
git clone git@github.com:kyle-cao/submitter_search_engine.git
cd submitter_search_engine
go mod tidy
go build -o submitter
```

## Usage

### Command Line

Submit a single URL:

```bash
./submitter cmd --urls=https://example.com,https://example.com --engines=baidu,bing,google
```

### HTTP API

Start the API server:

```bash
submitter http --host=0.0.0.0 --port=8080
```

## Configuration

Create a config file `./config/config.json`:

```json
{
  "baidu": {
    "site": "https://example.com",
    "token": "123123"
  },
  "bing": {
    "siteUrl": "https://example.com",
    "apiKey": "123123"
  },
  "google": {
    "type": "service_account",
    "project_id": "xx",
    "private_key_id": "xxx",
    "private_key": "xxx",
    "client_email": "xxx",
    "client_id": "123",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_x509_cert_url": "xxx",
    "universe_domain": "googleapis.com"
  }
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)