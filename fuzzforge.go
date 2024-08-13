package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "sync"
    "time"
)

// ANSI color codes
const (
    ColorReset  = "\033[0m"
    ColorRed    = "\033[31m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
    ColorCyan   = "\033[36m"
    ColorWhite  = "\033[37m"
)

// Colorize function to wrap text with a given ANSI color code
func Colorize(color, text string) string {
    return color + text + ColorReset
}

// PrintBanner prints the FuzzForge banner
func PrintBanner() {
    banner := `
 ###### #    # ###### ###### ######  ####  #####   ####  ###### 
 #      #    #     #      #  #      #    # #    # #    # #      
 #####  #    #    #      #   #####  #    # #    # #      #####  
 #      #    #   #      #    #      #    # #####  #  ### #      
 #      #    #  #      #     #      #    # #   #  #    # #      
 #       ####  ###### ###### #       ####  #    #  ####  ###### 
                                                                

                 Made By z3r0X0r                                              
    `
    fmt.Println(Colorize(ColorCyan, banner))
    fmt.Println(Colorize(ColorYellow, "FuzzForge: High-Speed Go-based Directory & Parameter Fuzzer\n"))
}

// Fuzz performs the fuzzing by sending HTTP requests with payloads
func Fuzz(target string, payloads []string, extensions []string, concurrency int, mode string, outputFile *os.File) {
    var wg sync.WaitGroup
    sem := make(chan struct{}, concurrency)

    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    validStatusCodes := map[int]bool{
        200: true,
        302: true,
        403: true,
        405: true,
        500: true,
    }

    for _, payload := range payloads {
        for _, ext := range extensions {
            wg.Add(1)
            sem <- struct{}{}

            go func(payload, ext string) {
                defer wg.Done()
                defer func() { <-sem }()

                var url string
                if mode == "dir" {
                    url = fmt.Sprintf("%s/%s%s", target, payload, ext)
                } else if mode == "param" {
                    url = fmt.Sprintf("%s?%s=value", target, payload)
                }

                resp, err := client.Get(url)
                if err != nil {
                    fmt.Println(Colorize(ColorRed, fmt.Sprintf("[!] Request Failed: %v", err)))
                    if outputFile != nil {
                        fmt.Fprintf(outputFile, "[!] Request Failed: %v\n", err)
                    }
                    return
                }
                defer resp.Body.Close()

                if validStatusCodes[resp.StatusCode] {
                    var color string
                    if resp.StatusCode == 200 {
                        color = ColorGreen // Change to bright green for 200 status
                    } else {
                        color = ColorYellow
                    }
                    fmt.Println(Colorize(color, fmt.Sprintf("[+] %d: Valid Path/Parameter Found: %s", resp.StatusCode, url)))
                    if outputFile != nil {
                        fmt.Fprintf(outputFile, "[+] %d: Valid Path/Parameter Found: %s\n", resp.StatusCode, url)
                    }
                }
            }(payload, ext)
        }
    }

    wg.Wait()
}

// LoadPayloads loads payloads from a file
func LoadPayloads(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var payloads []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        payloads = append(payloads, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return payloads, nil
}

func main() {
    PrintBanner()

    var target string
    var wordlist string
    var extensions string
    var concurrency int
    var mode string
    var outputFileName string

    // Command-line flags
    flag.StringVar(&target, "u", "", "Target URL")
    flag.StringVar(&wordlist, "w", "", "Wordlist file")
    flag.StringVar(&extensions, "x", ".php,.html,.js,.txt", "Comma-separated list of file extensions (e.g., .php,.html)")
    flag.IntVar(&concurrency, "c", 100, "Concurrency level")
    flag.StringVar(&mode, "m", "dir", "Fuzzing mode: 'dir' for directory fuzzing, 'param' for parameter fuzzing")
    flag.StringVar(&outputFileName, "o", "", "Output file to save results")
    flag.Parse()

    // Check if required flags are provided
    if target == "" || wordlist == "" {
        log.Fatalf(Colorize(ColorRed, "[!] Target URL (-u) and Wordlist (-w) are required fields."))
    }

    // Open the output file if specified
    var outputFile *os.File
    var err error
    if outputFileName != "" {
        outputFile, err = os.Create(outputFileName)
        if err != nil {
            log.Fatalf("Failed to create output file: %v", err)
        }
        defer outputFile.Close()
    }

    extList := strings.Split(extensions, ",")
    payloads, err := LoadPayloads(wordlist)
    if err != nil {
        log.Fatalf("Failed to load payloads: %v", err)
    }

    Fuzz(target, payloads, extList, concurrency, mode, outputFile)
}
