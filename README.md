# Crawler

A simple concurrent web crawler written in Go.  
It fetches pages, extracts data (headings, first paragraph, links, images), and exports results to CSV.

---

## Requirements
- Go 1.21 or higher
- Internet connection

---

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/Witnot/crawler.git
   cd crawler
   ```

2. Build the binary:
   ```bash
   go build -o crawler
   ```

---

## Usage

Run the crawler with the following syntax:

```bash
./crawler <url> <maxConcurrency> <maxPages>
```

**Example:**
```bash
./crawler "https://example.com" 3 10
```

**Parameters:**
- `<url>`: Starting page to crawl
- `<maxConcurrency>`: Maximum number of concurrent requests (e.g., 3)
- `<maxPages>`: Maximum number of pages to crawl (e.g., 10)

The crawler will output results to `report.csv` in the current directory.

---

## Features

- Concurrency control to limit active requests
- Configurable maximum pages to crawl
- Normalizes URLs to avoid duplicates
- Extracts:
  - `<h1>` heading
  - First `<p>` paragraph
  - Outgoing links
  - Images
- Exports results to a CSV file

---

## Output

The `report.csv` file contains the following columns:
- `page_url`
- `h1`
- `first_paragraph`
- `outgoing_link_urls`
- `image_urls`

**Sample CSV output:**
```
page_url,h1,first_paragraph,outgoing_link_urls,image_urls
https://example.com,Welcome to Example,"This is the first paragraph of content.","https://example.com/about,https://example.com/contact","https://example.com/logo.png,https://example.com/banner.jpg"
```

---

## Notes

- Use a reasonable concurrency value to avoid overloading servers
- Stop the crawler at any time with `Ctrl+C`
- The crawler respects the same domain and follows links within the starting domain
