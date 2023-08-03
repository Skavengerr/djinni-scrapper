# Job Scraper Telegram Bot

This is a Telegram bot that scrapes job vacancies for Golang developers from the Djinni job board. The bot sends the latest job vacancies to users and filters out duplicate job postings.

## Requirements
- Djinni job board access (https://djinni.co/)

## Setup

1. Clone the repository:

```bash
git clone https://github.com/Skavengerr/job-scrapper.git
cd job-scraper
```


2. Install the necessary Go dependencies:

```bash
go mod download
go run cmd/main.go
```

## Usage

1. Start a chat with your bot on Telegram (https://t.me/job_go_scrapper_bot)
2. Send the `/start` command to the bot to initiate the job scraping process.

The bot will periodically scrape the Djinni job board for Golang developer vacancies and send you the latest jobs every 30 minutes. It will also filter out duplicate job postings that have already been sent to you.

## Data Storage

The bot stores job data in JSON files. Each user's data is stored in a separate JSON file named based on their Chat.ID. This allows the bot to remember the user's previous job postings and avoid duplicates.