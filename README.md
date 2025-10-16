<p align="center">
    <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" align="center" width="30%">
</p>
<p align="center"><h1 align="center">GO-WEATHER-MCP</h1></p>
<p align="center">
	<em><code>❯ REPLACE-ME</code></em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/alvarolspeixoto/go-weather-mcp?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/alvarolspeixoto/go-weather-mcp?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/alvarolspeixoto/go-weather-mcp?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/alvarolspeixoto/go-weather-mcp?style=default&color=0080ff" alt="repo-language-count">
</p>
<p align="center"><!-- default option, no dependency badges. -->
</p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>
<br>

##  Table of Contents

- [ Overview](#-overview)
- [ Features](#-features)
- [ Project Structure](#-project-structure)
  - [ Project Index](#-project-index)
- [ Getting Started](#-getting-started)
  - [ Prerequisites](#-prerequisites)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
---

##  Overview

This project implements an MCP Server that provides weather information based on an input city, besides implementing a API REST with the same functionality. This is made using Go and [Open Weather API](https://openweathermap.org/api) as information source.

An MCP (Model Context Protocol) server is a backend application that manages and provides the context required for an AI model to operate and respond to a request.

---

##  Features

MCP tools:
- get_weather_by_city (Retrieves the current weather for a specified city.)

---

##  Project Structure

```sh
└── go-weather-mcp/
    ├── cmd
    │   └── server
    ├── go.mod
    ├── go.sum
    └── internal
        ├── api
        ├── app
        ├── config
        ├── domain
        ├── infra
        └── mcp
```


###  Project Index
<details open>
	<summary><b><code>GO-WEATHER-MCP/</code></b></summary>
	<details> <!-- __root__ Submodule -->
		<summary><b>__root__</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/go.mod'>go.mod</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/go.sum'>go.sum</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- cmd Submodule -->
		<summary><b>cmd</b></summary>
		<blockquote>
			<details>
				<summary><b>server</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/cmd/server/main.go'>main.go</a></b></td>
						<td><code>❯ Entrypoint to the server</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- internal Submodule -->
		<summary><b>internal</b></summary>
		<blockquote>
			<details>
				<summary><b>mcp</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/mcp/server.go'>server.go</a></b></td>
						<td><code>❯ MCP Server configuration</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>config</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/config/config.go'>config.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>infra</b></summary>
				<blockquote>
					<details>
						<summary><b>openweathermap</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/infra/openweathermap/geocode_mapper.go'>geocode_mapper.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/infra/openweathermap/dto.go'>dto.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/infra/openweathermap/weather_mapper.go'>weather_mapper.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/infra/openweathermap/client.go'>client.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/infra/openweathermap/geocode_repository.go'>geocode_repository.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/infra/openweathermap/weather_repository.go'>weather_repository.go</a></b></td>
							</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
			<details>
				<summary><b>api</b></summary>
				<blockquote>
					<details>
						<summary><b>http</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/api/http/router.go'>router.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/api/http/handler.go'>handler.go</a></b></td>
							</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
			<details>
				<summary><b>domain</b></summary>
				<blockquote>
					<details>
						<summary><b>weather</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/domain/weather/repository.go'>repository.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/domain/weather/usecase.go'>usecase.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/domain/weather/entity.go'>entity.go</a></b></td>
							</tr>
							</table>
						</blockquote>
					</details>
					<details>
						<summary><b>geocode</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/domain/geocode/repository.go'>repository.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/domain/geocode/usecase.go'>usecase.go</a></b></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/domain/geocode/entity.go'>entity.go</a></b></td>
							</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
			<details>
				<summary><b>app</b></summary>
				<blockquote>
					<details>
						<summary><b>weather</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/app/weather/service.go'>service.go</a></b></td>
							</tr>
							</table>
						</blockquote>
					</details>
					<details>
						<summary><b>geocode</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/alvarolspeixoto/go-weather-mcp/blob/master/internal/app/geocode/service.go'>service.go</a></b></td>
							</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
		</blockquote>
	</details>
</details>

---
##  Getting Started

###  Prerequisites

Before getting started with go-weather-mcp, ensure your runtime environment meets the following requirements:

- **Go**: 1.21 or newer
- [OpenWeather API Key](https://openweathermap.org/api)

###  Installation

Install go-weather-mcp using one of the following methods:

**Build from source:**

1. Clone the go-weather-mcp repository:
```sh
❯ git clone https://github.com/alvarolspeixoto/go-weather-mcp
```

2. Navigate to the project directory:
```sh
❯ cd go-weather-mcp
```

3. Install the project dependencies:

Copy and paste .env.example, it to .env and put your api key in `OPENWEATHER_API_KEY` variable.

**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build
```




###  Usage

Run go-weather-mcp using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run ./cmd/server --mode={mode} --port={port}
```
- Modes: `mcp` (for running only MCP server), `http` (only HTTP server) and `dual` (runs both servers)

For testing the MCP Server in Claude Desktop, you should configure it in your `claude_desktop_config.json`. Example for Windows:

```js
{
  "mcpServers": {
    "weather": {
      "command": "C:\\Users\\alvar\\OneDrive\\Documentos\\dev\\go-mcp\\go-weather-mcp.exe",
      "args": [],
      "env": {
        "MODE": "mcp",
        "OPENWEATHER_API_KEY": "your-open-weather-api-key"
      }
    }
  }
}

```
---