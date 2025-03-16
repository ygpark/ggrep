package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// ANSI 색상 코드
const redColor = "\033[31m"
const resetColor = "\033[0m"

func highlightMatch(text string, pattern *regexp.Regexp, useColor bool) string {
	if !useColor {
		return text
	}
	return pattern.ReplaceAllStringFunc(text, func(match string) string {
		return redColor + match + resetColor
	})
}

func grepFile(pattern *regexp.Regexp, filename string, showLineNumbers, invertMatch, countOnly, useColor bool) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "파일을 열 수 없습니다: %s\n", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	matchCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.MatchString(line)

		if invertMatch {
			matches = !matches
		}

		if matches {
			matchCount++
			if !countOnly {
				highlightedLine := highlightMatch(line, pattern, useColor)
				if showLineNumbers {
					fmt.Printf("%s:%d: %s\n", filename, lineNumber, highlightedLine)
				} else {
					fmt.Printf("%s: %s\n", filename, highlightedLine)
				}
			}
		}
		lineNumber++
	}

	if countOnly {
		fmt.Printf("%s: %d\n", filename, matchCount)
	}
}

func grepStdin(pattern *regexp.Regexp, invertMatch, countOnly, useColor bool) {
	scanner := bufio.NewScanner(os.Stdin)
	matchCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.MatchString(line)

		if invertMatch {
			matches = !matches
		}

		if matches {
			matchCount++
			if !countOnly {
				fmt.Println(highlightMatch(line, pattern, useColor))
			}
		}
	}

	if countOnly {
		fmt.Println(matchCount)
	}
}

func grepDirectory(pattern *regexp.Regexp, dir string, showLineNumbers, invertMatch, countOnly, useColor bool) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "파일 접근 오류: %s\n", err)
			return nil
		}
		if !info.IsDir() {
			grepFile(pattern, path, showLineNumbers, invertMatch, countOnly, useColor)
		}
		return nil
	})
}

func main() {
	ignoreCase := flag.Bool("i", false, "대소문자 구분 없이 검색")
	showLineNumbers := flag.Bool("n", false, "라인 번호 출력")
	invertMatch := flag.Bool("v", false, "매칭되지 않는 라인 출력")
	countOnly := flag.Bool("c", false, "일치 개수만 출력")
	recursive := flag.Bool("r", false, "디렉토리 재귀 검색")
	useColor := flag.Bool("color", false, "컬러 출력")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("사용법: ggrep [옵션] <패턴> [파일...]")
		os.Exit(1)
	}

	patternStr := flag.Arg(0)
	if *ignoreCase {
		patternStr = "(?i)" + patternStr
	}

	pattern, err := regexp.Compile(patternStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "잘못된 정규식: %s\n", err)
		os.Exit(1)
	}

	if flag.NArg() == 1 {
		grepStdin(pattern, *invertMatch, *countOnly, *useColor)
	} else {
		for _, path := range flag.Args()[1:] {
			fileInfo, err := os.Stat(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "파일 접근 오류: %s\n", path)
				continue
			}

			if fileInfo.IsDir() {
				if *recursive {
					grepDirectory(pattern, path, *showLineNumbers, *invertMatch, *countOnly, *useColor)
				} else {
					fmt.Fprintf(os.Stderr, "디렉토리입니다 (옵션 -r 필요): %s\n", path)
				}
			} else {
				grepFile(pattern, path, *showLineNumbers, *invertMatch, *countOnly, *useColor)
			}
		}
	}
}
