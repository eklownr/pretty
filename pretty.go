package pretty

import (
	"fmt"
	"regexp"
	"strings"
)

// colorized text. Replace fmt.Printf(string, multiple values)
func Pf(s string, value ...interface{}) string {
	return Colorize(fmt.Sprintf(s+"\n", value...))
}

// Println with color. Use like fmt.Println
func Pl(args ...interface{}) {
	var arg string = fmt.Sprint(args...)
	println(Colorize(arg))
}

// colorize: text=Cyan, numbers=Green, specialchar=Yellow and error=Red
func Colorize(s string) string {
	w := strings.Split(s, " ") // split s into words
	for i := range w {
		// numbers w[i] = "\033[32m" + w[i] + "\033[0m"
		if regexp.MustCompile(`\d`).MatchString(w[i]) || regexp.MustCompile(`\d`).MatchString(w[i]) {
			w[i] = "\033[32m" + w[i] + "\033[0m" // Green
		}
		// if 'Error!' or 'Warning!' red
		if w[i] == "Error!" || w[i] == "Warning!" {
			w[i] = "\033[31m" + w[i] + "\033[0m"
		}
		// Special char Yellow
		if regexp.MustCompile(`[^a-zA-Z:,.!åäöÅÄÖ\n]`).MatchString(w[i]) {
			w[i] = "\033[33m" + w[i] + "\033[0m"
		}
		// Cyan text
		w[i] = "\033[36m" + w[i] + "\033[0m"
	}
	return strings.Join(w, " ")
}
// print one color for one string
func PrintColor(text string, color string) {
	switch color {
	case "red":
		fmt.Printf("\033[31m%s\033[0m\n", text)
	case "green":
		fmt.Printf("\033[32m%s\033[0m\n", text)
	case "yellow":
		fmt.Printf("\033[33m%s\033[0m\n", text)
	case "blue":
		fmt.Printf("\033[34m%s\033[0m\n", text)
	case "magenta":
		fmt.Printf("\033[35m%s\033[0m\n", text)
	case "cyan":
		fmt.Printf("\033[36m%s\033[0m\n", text)
	case "white":
		fmt.Printf("\033[37m%s\033[0m\n", text)
	case "orange":
		fmt.Printf("\033[38:5:166m%s\033[0m\n", text)
	default: // white
		fmt.Printf("\033[37m%s\033[0m\n", text)
	}
}

// print red text
func Pr(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[31m"+format+"\033[0m", a...)
}

// print green text
func Pg(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[32m"+format+"\033[0m", a...)
}

// print yellow text
func Py(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[33m"+format+"\033[0m", a...)
}

// print blue text
func Pb(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[34m"+format+"\033[0m", a...)
}

// print magenta text
func Pm(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[35m"+format+"\033[0m", a...)
}

// print cyan text
func Pc(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[36m"+format+"\033[0m", a...)
}

// print purple text
func Pp(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[38m"+format+"\033[0m", a...)
}

// print orange text
func Po(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[39m"+format+"\033[0m", a...)
}

// print black text
func Pk(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[40m"+format+"\033[0m", a...)
}

// print white text
func Pw(format string, a ...interface{}) string {
	return fmt.Sprintf("\033[37m"+format+"\033[0m", a...)
}
