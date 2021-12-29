package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const colorReset = "\033[0m"

const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const colorWhite = "\033[37m"

const bgYellow = "\033[43m"

const banner = `
	$$\   $$\  $$$$$$\  $$\ $$$$$$$$\ 
	$$ |  $$ |$$  __$$\ \__|$$  _____|
	\$$\ $$  |$$ /  \__|$$\ $$ |      
	 \$$$$  / $$ |$$$$\ $$ |$$$$$\    
	 $$  $$<  $$ |\_$$ |$$ |$$  __|   
	$$  /\$$\ $$ |  $$ |$$ |$$ |      
	$$ /  $$ |\$$$$$$  |$$ |$$ |      
	\__|  \__| \______/ \__|\__|      

	XGiF (eXposed Git Finder)
	https://github.com/prasant-paudel/xgif
`

const help = `
Usage: xgif [OPTIONS]

OPTIONS:
 -t		Target URL
 -T		List of targets [File]
 -v		Verbose mode
 -o		Output file

Example: xgif -t https://github.com -vv
Example: xgif -T target_urls.txt -v
Example: xgif -T target_urls.txt -o output.txt
`

func printLegends() {
	fmt.Println(colorCyan + "----: Legends :----" + colorReset)
	fmt.Println(colorRed + bgYellow + "Potentially Exploitable" + colorReset)
	fmt.Println(colorRed + "Connection Error" + colorReset)
	fmt.Println(colorYellow + "Other HTTP Errors" + colorReset)
	fmt.Println(colorCyan + "-------------------" + colorReset)
	fmt.Println()
}

var target string
var targets_path string
var verbose bool
var veryVerbose bool
var outputFile string

func argParse() {
	flag.StringVar(&target, "t", "", "Target URL")
	flag.StringVar(&targets_path, "T", "", "List of targets [File]")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.BoolVar(&veryVerbose, "vv", false, "Very verbose mode")
	flag.StringVar(&outputFile, "o", "", "Output file")

	flag.Parse()

	fmt.Println(target, targets_path, verbose)
}

func main() {
	argParse()
	fmt.Println(colorRed + banner + colorReset)
	start := time.Now()
	ch := make(chan string)

	var targets []string
	var _targets []string

	if targets_path != "" {
		_targets = readLines(targets_path)
	}
	if target != "" {
		_targets = append(_targets, strings.TrimSpace(target))
	}

	if target == "" && targets_path == "" {
		fmt.Printf(colorCyan + help + colorReset)
	} else {
		printLegends()
		for _, ln := range _targets {
			ln = strings.TrimSpace(ln)
			if len(ln) != 0 && contains(targets, ln) == false {
				targets = append(targets, ln)
			}
		}

		for _, ln := range targets {
			go checkGitConfig(getBaseUrl(ln), ch)
		}

		for range targets {
			fmt.Print(<-ch)

		}
		fmt.Println(colorCyan+"Targets tested: ", len(targets), colorReset)
		fmt.Printf(colorCyan+"Time elapsed  : %.2fs\n"+colorReset, time.Since(start).Seconds())
	}

}

func checkGitConfig(baseUrl string, ch chan<- string) {
	url := baseUrl + "/.git/config"
	resp := getReq(url)
	if resp == "Connection Error" && veryVerbose {
		ch <- fmt.Sprintln(colorRed + "[Connection Error] " + url + colorReset)
	} else if strings.Contains(resp, "Status:") && verbose {
		ch <- fmt.Sprintln(colorYellow + resp + " " + url + colorReset)
	} else if strings.Contains(resp, "[core]") {
		ch <- fmt.Sprintln(bgYellow+colorRed+url, " *** Potentially Exploitable *** "+colorReset)
	} else {
		ch <- ""
	}
}

func getReq(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Connection Error"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		return string(body)
	}
	return fmt.Sprint("Status:", resp.StatusCode)
}

func readLines(filename string) []string {
	var lines []string

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()
	return lines
}

func getBaseUrl(str string) string {
	if !strings.Contains(str, "//") {
		return "http://" + str
	}

	schema := strings.Split(str, "://")[0]
	domain := strings.Split(strings.Split(str, "://")[1], "/")[0]
	return fmt.Sprint(schema + "://" + domain)
}

func contains(arr []string, e string) bool {
	for _, i := range arr {
		if i == e {
			return true
		}
	}
	return false
}
