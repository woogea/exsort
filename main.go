package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	rank       int
	column     int
	reg        string
	asc        bool
	onlyColumn bool
	delimiter  string
	rootCmd    = &cobra.Command{
		Use:   "exsort",
		Short: "sort from stdin by selected columns for version",
		Long:  ` sort from stdin by selected columns for version. sort v1.2.3 as 1002003 when rank is 1000`,
		Run: func(cmd *cobra.Command, args []string) {
			lines := []string{}
			stdin := bufio.NewScanner(os.Stdin)
			for stdin.Scan() {
				lines = append(lines, stdin.Text())
			}
			result := exsort(lines, reg, column, rank, asc)
			fmt.Println(result)
			for _, l := range result {
				if onlyColumn {
					fmt.Println(strings.Split(l, delimiter))
				} else {
					fmt.Println(l)
				}
			}
		},
	}
)

func exsort(lines []string, regex string, column int, rank int, asc bool) []string {
	//remove extra lines
	proclines := []string{}
	for i := range lines {
		if len(strings.Split(lines[i], delimiter)) > column && len(lines[i]) > 0 {
			proclines = append(proclines, lines[i])
		}
	}
	lines = proclines
	fmt.Println(lines, regex, column, rank, asc)
	less := func(i, j int) bool {
		if len(lines[i]) < column || len(lines[j]) < column {
			panic(lines)
		}
		tmpa := strings.Split(lines[i], delimiter)[column]
		tmpb := strings.Split(lines[j], delimiter)[column]
		re := regexp.MustCompile(regex)
		va := 0
		vb := 0
		astr := re.Split(tmpa, 255)
		for _, x := range astr {
			va *= rank
			t, _ := strconv.Atoi(x)
			va += t
		}
		bstr := re.Split(tmpb, 255)
		for _, x := range bstr {
			vb *= rank
			t, _ := strconv.Atoi(x)
			vb += t
		}
		return va > vb
	}
	fmt.Println(lines)
	sort.SliceStable(lines, less)
	if asc {
		rev := []string{}
		for i := range lines {
			rev = append(rev, lines[len(lines)-1-i])
		}
		lines = rev
	}
	return lines
}

func main() {
	rootCmd.PersistentFlags().IntVar(&rank, "rank", 100, "set 10 for 1.2.3 means 1*10^2+2*10^1+3*10^0")
	rootCmd.PersistentFlags().IntVar(&column, "column", 0, "which is used on sort")
	rootCmd.PersistentFlags().BoolVar(&asc, "asc", false, "sort order.true is asc, false is desc")
	rootCmd.PersistentFlags().StringVar(&reg, "reg", "[.]", "regexp for split. when 1.2.3 to \"[.]\"")
	rootCmd.PersistentFlags().BoolVar(&onlyColumn, "only-column", false, "output only specified column")
	rootCmd.PersistentFlags().StringVar(&delimiter, "delimiter", " ", "column separator")
	er := rootCmd.Execute()
	if er != nil {
		println(er)
	}
}
