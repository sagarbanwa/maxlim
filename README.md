# Subdomain Enumeration Toolkit

This toolkit automates the process of subdomain enumeration using various tools and provides additional functionalities for organizing and analyzing the results.

## Tools Included:

- [subfinder](https://github.com/projectdiscovery/subfinder)
- [assetfinder](https://github.com/tomnomnom/assetfinder)
- [findomain](https://github.com/Edu4rdSHL/findomain)
- [amass](https://github.com/OWASP/Amass) (passive)
- Sorting and removing duplicate domains
- Live domain detection using [httpx](https://github.com/projectdiscovery/httpx) with technology detection (results saved in a separate file)
- Port scanning using [naabu](https://github.com/projectdiscovery/naabu)
- Screenshots of web pages using [gowitness](https://github.com/sensepost/gowitness)

Feel free to add more tools or customize the configuration as needed. (also you can contribute to the project)

## Usage:

Run the following command:

``` maxlim.exe -l urls.tx ```

``` ./maxlim -l urls.txt ```

``` maxlim.exe domain.com ```
