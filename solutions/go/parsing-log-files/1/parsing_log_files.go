package parsinglogfiles

import "regexp"
import "fmt"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(~|\*|=|-)*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var counter int
    re := regexp.MustCompile(`".*(?i)password.*"`)

	for _, line := range lines {
        if re.MatchString(line) {
            counter++
        }
    }
    
    return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+[^\s]+`)
    userRe := regexp.MustCompile(`User\s+`)
    for index, line := range lines {
        stringFound := re.FindString(line)
        if stringFound != "" {
            fmt.Println(line)
            fmt.Println("ha matchato")
            userName := userRe.ReplaceAllString(stringFound, "")
            fmt.Println("trovato username: " + userName)
            lines[index] = fmt.Sprintf("[USR] %s %s", userName, line)
            fmt.Println("nuova line: " + line)
        }
    }
    return lines
}
