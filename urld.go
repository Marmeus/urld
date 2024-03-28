package main

import (
    "bufio"
    "flag"
    "fmt"
    "net/url"
    "os"
)

func main() {
    inputFlag := flag.Bool("i", false, "Read data passed through a Linux pipe")
    encodeFlag := flag.Bool("e", false, "URL encode data (default: Decode)")
    fileFlag := flag.String("f", "", "Read data stored in a file")

    flag.Parse()

    if flag.NFlag() == 0 || flag.Arg(0) == "-h" {
        flag.Usage()
        return
    }

    var data []string

    // Read data from pipe if -i flag is provided
    if *inputFlag {
        data = readPipe()
    }

    // Read data from file if -f flag is provided
    if *fileFlag != "" {
        fileData, err := readFile(*fileFlag)
        if err != nil {
            fmt.Println("Error reading file:", err)
            return
        }
        data = append(data, fileData...)
    }

    // Encode or decode data based on -e flag
    var processedData []string
    if *encodeFlag {
        for _, d := range data {
            processedData = append(processedData, url.QueryEscape(d))
        }
    } else {
        for _, d := range data {
            decoded, err := url.QueryUnescape(d)
            if err != nil {
                fmt.Println("Error decoding data:", err)
                return
            }
            processedData = append(processedData, decoded)
        }
    }

    // Print the processed data
    for _, d := range processedData {
        fmt.Println(d)
    }
}

func readPipe() []string {
    var data []string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        data = append(data, scanner.Text())
    }
    return data
}

func readFile(filename string) ([]string, error) {
    var data []string
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        data = append(data, scanner.Text())
    }
    return data, scanner.Err()
}

