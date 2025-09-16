# WebScraper

````markdown
````
#  Crawler

A simple concurrent web crawler written in Go.  
It fetches pages, extracts data (headings, first paragraph, links, images), and exports results to CSV.

---

##  Requirements
- [Go 1.21+](https://go.dev/dl/)
- Internet connection 😅

---

##  Usage

Build the binary:

```bash
go build -o crawler
````

Run the crawler:

```bash
./crawler "https://example.com" <maxConcurrency> <maxPages>
```

Example:

```bash
./crawler "https://wikipedia.org" 3 10
```

This crawls up to **10 pages**, using **3 concurrent workers**, starting from `wikipedia.org`.
The results will be saved to `report.csv`.

---

##  Features

* Concurrency control (limit active requests).
* Configurable max pages to crawl.
* Normalizes URLs to avoid duplicates.
* Extracts:

  * `<h1>` heading
  * First `<p>` paragraph
  * Outgoing links
  * Images
* Exports results to **CSV report**.

---

##  Output

The crawler generates a `report.csv` file with these columns:

* `page_url`
* `h1`
* `first_paragraph`
* `outgoing_link_urls`
* `image_urls`

---

##  Notes

* Be kind to servers: don’t set concurrency too high.
* Use `Ctrl+C` to stop crawling at any time.

```

