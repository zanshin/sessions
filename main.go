package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool.\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n\n")
		flag.PrintDefaults()
	}

	service := flag.String("s", "", "Name of service, including the tier")
	inst := flag.Int("i", 4, "Instance count, number of app nodes, defaults to 4")
	flag.Parse()

	// Use tabwriter to control columns in output
	writer := tabwriter.NewWriter(os.Stdout, 10, 5, 1, ' ', tabwriter.AlignRight)

	fmt.Fprintln(writer, "node\tblue\t   green")

	for i := 1; i <= *inst; i++ {
		host := fmt.Sprintf("ome-%s-app-0%d", *service, i)

		blueCmd, greenCmd := buildCurl(host, i)
		blue, err := runCurl(blueCmd)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		green, err := runCurl(greenCmd)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Fprintln(writer, fmt.Sprintf("%s\t%s\t%s", host, blue, green))

	}

	// Flush the write to get our output
	writer.Flush()
}

func buildCurl(host string, instance int) (string, string) {
	grepb := "| egrep 'id:.*:blue_cluster' | wc -l"
	grepg := "| egrep 'id:.*:green_cluster' | wc -l"
	mod := "/mod_cluster-manager/"

	fqdn := fmt.Sprintf("%s.prod.aws.ksu.edu", host)
	url := fmt.Sprintf("http://%s.%s", fqdn, mod)

	blue := fmt.Sprintf("curl -s %s %s", url, grepb)
	green := fmt.Sprintf("curl -s %s %s", url, grepg)

	return blue, green

}

func runCurl(cmd string) ([]byte, error) {
	result, err := exec.Command("bash", "-c", cmd).Output()

	// Strip newline from end of result
	result[len(result)-1] = 0x20

	return result, err
}
