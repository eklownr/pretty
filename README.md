# pretty
Pretty colors for the terminal.
```go
// replacement for fmt.Printf(string, multiple values). 
func Pl(s string, value ...interface{}) string {
	return Colorize(fmt.Sprintf(s+"\n", value...))
}

// colorize: text=Cyan, numbers=Green, special char and error=Red
func Colorize(text string) string {
  ...
}

// print string in one color
func PrintColor(text string, color string) {
	switch color {
	case "red":
		fmt.Printf("\033[31m%s\033[0m\n", text)
	case "green":
    ...
```
