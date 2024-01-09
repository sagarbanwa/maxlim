package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}


func main() {
	var domain string

	if len(os.Args) > 1 {
		if os.Args[1] == "-l" && len(os.Args) > 2 {
			// Read the domain from the file specified by -l option
			filePath := os.Args[2]
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error opening file:", err)
				os.Exit(1)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			if scanner.Scan() {
				domain = scanner.Text()
			} else {
				fmt.Println("Error reading domain from file.")
				os.Exit(1)
			}
		} else {
			// Use the domain from command-line argument
			domain = os.Args[1]
		}
	} else {
		// Prompt the user to enter a domain interactively
		fmt.Print("Enter domain name: ")
		fmt.Scanln(&domain)
	}
	fmt.Println("[*] Running subdomain enumeration...")
	runCommand("subfinder", "-d", domain, "-silent", "-all", "-t", "100", "-o", fmt.Sprintf("%s/subfinder-01.txt", domain))

	fmt.Println("[*] Running subdomain enumeration for found domain")
	runCommand("subfinder", "-dL", fmt.Sprintf("%s/subfinder-01.txt", domain), "-silent", "-all", "-t", "70", "-o", fmt.Sprintf("%s/subfinder-02.txt", domain))

	fmt.Println("[*] Running subdomain enumeration for found domain")
	runCommand("subfinder", "-dL", fmt.Sprintf("%s/subfinder-02.txt", domain), "-silent", "-all", "-t", "70", "-o", fmt.Sprintf("%s/subfinder-03.txt", domain))

	fmt.Println("[*] Running assetfinder enumeration...")
	runCommand("assetfinder", "--subs-only", domain, ">", fmt.Sprintf("%s/assetfinder.txt", domain))

	runCommand("subfinder", "-dL", fmt.Sprintf("%s/assetfinder.txt", domain), "-silent", "-all", "-t", "70", "-o", fmt.Sprintf("%s/subfinder-05-assetfinder.txt", domain))

	fmt.Println("[*] Running findomain enumeration...")
	runCommand("findomain", "--quiet", "-t", domain, "-u", fmt.Sprintf("%s/findomain.txt", domain))

	fmt.Println("[*] Running findomain enumeration...")
	runCommand("findomain", "--quiet", "-t", domain, "-u", fmt.Sprintf("%s/findomainx.txt", domain))
	
	runCommand("amass", "enum", "-passive", "-d", domain, fmt.Sprintf("%s/findomainx.txt", domain))

	fmt.Println("[*] Sorting and removing duplicates from all subdomain files...")
	runCommand("sh", "-c", fmt.Sprintf("grep -vE '^_' %s/*.txt | sed 's/[^a-zA-Z0-9.-]/ /g' | sort | uniq > %s/finaloutput.txt", domain, domain))

	fmt.Println("[*] Running httpx to find live domains status !!")
	runCommand("sh", "-c", fmt.Sprintf("cat %s/finaloutput.txt | httpx -silent -threads 50 > %s/httpx-live.txt", domain, domain))

	fmt.Println("[*] Running httpx to get status code, title, and technologies...")
	runCommand("sh", "-c", fmt.Sprintf("cat %s/httpx-live.txt | httpx -status-code -title -tech-detect -follow-redirects -threads 60 > %s/httpx-status.txt", domain, domain))

	fmt.Println("[*] Running naabu for ports scan...")
	runCommand("naabu", "-l", fmt.Sprintf("%s/finaloutput.txt", domain), "-p", "21,22,80,81,280,300,443,583,591,593,832,981,1010,1099,1311,2082,2087,2095,2096,2480,3000,3128,3333,4243,4444,4445,4567,4711,4712,4993,5000,5104,5108,5280,5281,5601,5800,6543,7000,7001,7002,7396,7474,8000,8001,8008,8009,8014,8042,8060,8069,8080,8081,8083,8088,8090,8091,8095,8118,8123,8172,8181,8222,8243,8280,8281,8333,8337,8443,8500,8530,8531,8834,8880,8887,8888,8983,9000,9001,9043,9060,9080,9090,9091,9092,9200,9443,9502,9800,9981,10000,10250,10443,11371,12043,12046,12443,15672,16080,17778,18091,18092,20720,28017,32000,55440,55672", "-silent", "-rate", "60", fmt.Sprintf("%s/naabu-ports-scan.txt", domain))

	fmt.Println("[*] Running gowitness to capture screenshots...")
	runCommand("gowitness", "file", "-f", fmt.Sprintf("%s/httpx-live.txt", domain), "--fullpage", "--threads", "2", "-P", fmt.Sprintf("%s/screenshots", domain))
	

	fmt.Println("[*] END of the Scan")
}
