/*
 * @description       :
 * @version           : "1.0.0"
 * @creator           : Gordon Lim <honwei189@gmail.com>
 * @created           : 27/05/2020 13:54:45
 * @last modified     : 29/05/2020 22:05:54
 * @last modified by  : Gordon Lim <honwei189@gmail.com>
 */

package utilities

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	// . "github.com/logrusorgru/aurora"

	"github.com/gookit/color"
	"github.com/noaway/dateparse"
)

// Addslashes : Add slashes to the string
func Addslashes(str string) string {
	tmpRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

// BToMb : Convert bytes to Mb format
func BToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Clearscreen : Clear console screen
func Clearscreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// CmdExec : Execute command and return output
func CmdExec(args ...string) (string, error) {
	// Usage :
	// out, err := cmdExec("ls", "/home")
	baseCmd := args[0]
	cmdArgs := args[1:]

	// log.Println("Exec: ", args)

	cmd := exec.Command(baseCmd, cmdArgs...)
	out, err := cmd.Output()

	if err != nil {
		// return "", err

		var output string
		cmd = exec.Command(baseCmd, cmdArgs...)
		r, _ := cmd.StdoutPipe()

		// if err != nil {
		// 	return "", err
		// }

		// Use the same pipe for standard error
		cmd.Stderr = cmd.Stdout

		// Make a new channel which will be used to ensure we get all output
		done := make(chan struct{})

		// Create a scanner which scans r in a line-by-line fashion
		scanner := bufio.NewScanner(r)

		// Use the scanner to scan the output line by line and log it
		// It's running in a goroutine so that it doesn't block
		go func() {

			// Read line by line and process it
			for scanner.Scan() {
				output = output + "\n" + scanner.Text()
			}

			// We're all done, unblock the channel
			done <- struct{}{}

		}()

		// Start the command and check for errors
		err = cmd.Start()

		// Wait for all output to be processed
		<-done

		// Wait for the command to finish
		err = cmd.Wait()
		err = errors.New(output)
		output = ""
		return "", err
	}

	return string(out), nil
}

// func cmdRun(args ...string) (string, error) {
// 	baseCmd := args[0]
// 	cmdArgs := args[1:]

// 	out, err := exec.Command(baseCmd, cmdArgs...).Output()

// 	if err != nil {
// 		log.Fatal(err)
// 		return "", err
// 	}

// 	return string(out), nil
// }

// CmdRun : Execute command and return output as array
func CmdRun(args ...string) ([]string, error) {
	data := []string{}
	baseCmd := args[0]
	cmdArgs := args[1:]

	// The command you want to run along with the argument
	cmd := exec.Command(baseCmd, cmdArgs...)

	// log.Println("Exec: ", args)

	// Get a pipe to read from standard out
	r, err := cmd.StdoutPipe()

	if err != nil {
		return nil, err
	}

	// Use the same pipe for standard error
	cmd.Stderr = cmd.Stdout

	// Make a new channel which will be used to ensure we get all output
	done := make(chan struct{})

	// Create a scanner which scans r in a line-by-line fashion
	scanner := bufio.NewScanner(r)

	// Use the scanner to scan the output line by line and log it
	// It's running in a goroutine so that it doesn't block
	go func() {

		// Read line by line and process it
		for scanner.Scan() {
			line := scanner.Text()
			data = append(data, line)
		}

		// We're all done, unblock the channel
		done <- struct{}{}

	}()

	// Start the command and check for errors
	err = cmd.Start()

	// Wait for all output to be processed
	<-done

	// Wait for the command to finish
	err = cmd.Wait()
	return data, nil
}

// ConvertUTCDateTime : Convert date time to YYYY-mm-dd format
func ConvertUTCDateTime(datetime string) string {
	// loc, err := time.LoadLocation("+0800")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// time.Local = loc

	denverLoc, _ := time.LoadLocation("Asia/Kuching")
	t, err := dateparse.ParseIn(datetime, denverLoc)

	re := regexp.MustCompile(`(.?)[+-][0-9]{4}\b`)
	zone := string(re.Find([]byte(datetime)))
	datetime = strings.TrimSpace(strings.Replace(datetime, zone, "", -1))

	t, err = dateparse.ParseAny(datetime)
	if err != nil {
		panic(err.Error())
	}

	// const layout = "2020-12-31 00:00:00"
	// t, err = time.Parse(layout, fmt.Sprintf("%v", t))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// return fmt.Sprintf("%v", t)
	return fmt.Sprintf("%s", t.Format("2006-01-02 03:04:05 PM"))
}

// ConvertUTCDate : Convert date time to YYYY-mm-dd format
func ConvertUTCDate(datetime string) string {
	// loc, err := time.LoadLocation("+0800")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// time.Local = loc

	denverLoc, _ := time.LoadLocation("Asia/Kuching")
	t, err := dateparse.ParseIn(datetime, denverLoc)

	re := regexp.MustCompile(`(.?)[+-][0-9]{4}\b`)
	zone := string(re.Find([]byte(datetime)))
	datetime = strings.TrimSpace(strings.Replace(datetime, zone, "", -1))

	t, err = dateparse.ParseAny(datetime)
	if err != nil {
		panic(err.Error())
	}

	// const layout = "2020-12-31 00:00:00"
	// t, err = time.Parse(layout, fmt.Sprintf("%v", t))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// return fmt.Sprintf("%v", t)
	return fmt.Sprintf("%s", t.Format("2006-01-02"))
}

// ConvertUTCTime : Convert date time to HH:mm:ss AMPM format
func ConvertUTCTime(datetime string) string {
	denverLoc, _ := time.LoadLocation("Asia/Kuching")
	t, err := dateparse.ParseIn(datetime, denverLoc)

	re := regexp.MustCompile(`(.?)[+-][0-9]{4}\b`)
	zone := string(re.Find([]byte(datetime)))
	datetime = strings.TrimSpace(strings.Replace(datetime, zone, "", -1))

	t, err = dateparse.ParseAny(datetime)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", t.Format("03:04:05 PM"))
}

// Convert12Time : Convert date time to HH:mm:ss AMPM format
func Convert12Time(fulltime string) string {
	const (
		layoutISO = "15:04:05"
		layoutUS  = "January 2, 2006"
	)

	t, _ := time.Parse(layoutISO, fulltime)
	return fmt.Sprintf("%s", t.Format("03:04:05 PM"))
}

// Convert24Time : Convert date time to HH:mm:ss AMPM format
func Convert24Time(fulltime string) string {
	const (
		layoutISO = "15:04:05"
		layoutUS  = "January 2, 2006"
	)

	t, _ := time.Parse(layoutISO, fulltime)
	return fmt.Sprintf("%s", t.Format("15:04:05"))
}

// CreateFileWithLongLine : Create longtext file
// func CreateFileWithLongLine(fn string) (err error) {
// 	file, err := os.Create(fn)
// 	defer file.Close()

// 	if err != nil {
// 		return err
// 	}

// 	w := bufio.NewWriter(file)

// 	fs := 1024 * 1024 * 4 // 4MB

// 	// Create a 4MB long line consisting of the letter a.
// 	for i := 0; i < fs; i++ {
// 		w.WriteRune('a')
// 	}

// 	// Terminate the line with a break.
// 	w.WriteRune('\n')

// 	// Put in a second line, which doesn't have a linebreak.
// 	w.WriteString("Second line.")

// 	w.Flush()

// 	return
// }

//Deletefile : Delete file if exists
func Deletefile(f string) {

	if FileExists(f) {
		var err = os.Remove(f)
		if err != nil {
			color.Error.Println(err.Error())
		}
	}

}

// DirExists : Check is dir exists
func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// Escape : To handle special characters from string
func Escape(s string) string {
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "-"}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	s = re.ReplaceAllString(s, "")
	return s
}

// FileExists : Check file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FilePutContents : Create / Overwrite file and write contents
func FilePutContents(fn string, content string) (err error) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)
	w.WriteString(content)
	w.Flush()

	return

	// file, err := os.Create(fn)
	// defer file.Close()

	// if err != nil {
	// 	return err
	// }

	// w := bufio.NewWriter(file)

	// w.WriteString(content)
	// // Terminate the line with a break.
	// // w.WriteRune('\n')

	// w.Flush()

	// return
}

// Isset : Check is array key exists
func Isset(arr []string, index int) bool {
	return (len(arr) > index)
}

// LimitLength : Return portion of string specified from 0 and length.
func LimitLength(s string, length int) string {
	if len(s) < length {
		return s
	}

	return s[:length]
}

// PrintMemUsage : Show current memory used
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", BToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", BToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", BToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

//ReadFileWithReadString : Read file with using bufio.NewReader method.
func ReadFileWithReadString(fn string) (err error) {
	// fmt.Println("readFileWithReadString")

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')

		// fmt.Printf(" > Read %d characters\n", len(line))

		// Process the line here.
		// fmt.Println(" > > " + LimitLength(line, 50))

		fmt.Println(line)

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return
}

//Readfile : Read file by using bufio.NewScanner.  This method is fast
func Readfile(fn string) (err error) {
	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return err
	}

	// Start reading from the file using a scanner.
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// fmt.Printf(" > Read %d characters\n", len(line))

		// Process the line here.
		// fmt.Println(" > > " + LimitLength(line, 50))
		fmt.Println(line)
	}

	if scanner.Err() != nil {
		fmt.Printf(" > Failed!: %v\n", scanner.Err())
	}

	return
}

//ReadFileWithReadLine : Read file by using bufio.NewReader method
func ReadFileWithReadLine(fn string) (err error) {
	fmt.Println("readFileWithReadLine")

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	for {
		var buffer bytes.Buffer

		var l []byte
		var isPrefix bool
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}

			// If we're just at the EOF, break
			if err != nil {
				break
			}
		}

		if err == io.EOF {
			break
		}

		line := buffer.String()

		// fmt.Printf(" > Read %d characters\n", len(line))

		// Process the line here.
		// fmt.Println(" > > " + LimitLength(line, 50))
		fmt.Println(line, 50)
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return
}

//RegSplit : regex split string.  e.g: utilities.RegSplit(line, `\S+`) //split by whitespace
func RegSplit(text string, pattern string) []string {
	re := regexp.MustCompile(pattern)

	// fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	// fmt.Printf("String contains any match: %v\n", re.MatchString(line)) // True

	submatchall := re.FindAllString(text, -1)
	// for _, element := range submatchall {
	// 	fmt.Println(element)
	// }

	return submatchall
}

// RemoveArrayIndex : Remove array element by key
func RemoveArrayIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
	// newArr := make([]string, (len(s) - 1))
	// k := 0
	// for i := 0; i < (len(s) - 1); {
	// 	if i != index {
	// 		newArr[i] = s[k]
	// 		k++
	// 	} else {
	// 		k++
	// 	}
	// 	i++
	// }

	// return newArr
}

//SplitLines : Split string with new lines
func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

// Stripslashes : Remove slashes from string
func Stripslashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

// Substr : Display string from end with specific length
func Substr(s string, optional ...int) string {
	str := []rune(s)
	trimmed := ""
	start := 0
	end := 0

	// fmt.Println(string(str))

	if len(optional) == 2 {
		start = optional[0]
		end = optional[1]

		if start < 0 {
			trimmed = string(str[len(str)-3:])
			trimmed = Substr(trimmed, 0, end)
		} else {
			if end < 0 {
				trimmed = string(str[start : len(str)+end])
			} else {
				trimmed = string(str[start:end])
			}
		}
	} else {
		start = optional[0]

		if start < 0 {
			trimmed = string(str[len(str)+start:])
		}
	}

	return strings.TrimSpace(trimmed)
}

// SubstrLast : Remove characters from end with specific length
func SubstrLast(s string, length int) string {
	// r, size := utf8.DecodeLastRuneInString(s)
	// if r == utf8.RuneError && (size == 0 || size == 1) {
	// 	size = 0
	// }

	return strings.TrimSpace(s[:len(s)-length])
}

// os.IsExist() would be blind to EMPTY FILE. Please always consider IsNotExist() instead.
// func isExist(err error) bool {
// 	err = os.underlyingError(err)
// 	return err == syscall.EEXIST || err == syscall.ENOTEMPTY || err == ErrExist
// }

// func isNotExist(err error) bool {
// 	err = underlyingError(err)
// 	return err == syscall.ENOENT || err == ErrNotExist
// }
