# Concurrent Web Scraper - Worker Pool Pattern Practice
This concurrent web scraper is a simple project designed to help me practice the worker pool pattern using Goroutines and channels in Go. The main purpose of this web scraper is to fetch the titles of multiple web pages concurrently.



# How it works
The web scraper consists of two main functions:

- fetchTitle(url string) (string, error): This function takes a URL as input, sends an HTTP GET request to fetch the web page, and extracts the title from the page's HTML content.

- fetchTitlesConcurrently(urls []string, numWorkers int) []string: This function takes a slice of URLs and an integer specifying the number of workers as input. It uses the worker pool pattern to concurrently fetch titles for the provided URLs. It controls the level of concurrency by limiting the number of Goroutines (workers) running at the same time.



# How to run
- Install Go if you haven't already.

- Clone or download this repository.

- Navigate to the directory containing the main.go file.

- Run the program with the following command:go run main.go



# Customizing the web scraper
You can customize the web scraper by modifying the list of URLs in the main function and adjusting the numWorkers value to control the level of concurrency. Remember that the numWorkers value acts as a rate limiter, limiting the number of concurrent tasks (Goroutines) that can run at the same time.
