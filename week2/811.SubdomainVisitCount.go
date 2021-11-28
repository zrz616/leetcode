import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	subdomainVisitedCounts := make(map[string]int)
	for _, cpdomain := range cpdomains {
		splitSubdomain(cpdomain, subdomainVisitedCounts)
	}
	var ans []string
	for domain, visited := range subdomainVisitedCounts {
		ans = append(ans, fmt.Sprintf("%d %s", visited, domain))
	}
	return ans
}

func splitSubdomain(cpdomain string, subdomainVisitedcounts map[string]int) {
	cpdomainStringsSplit := strings.Split(cpdomain, " ")
	visited, domain := cpdomainStringsSplit[0], cpdomainStringsSplit[1]
	subDomains := strings.Split(domain, ".")
	for i, _ := range subDomains {
		visited_count, err := strconv.Atoi(visited)
		if err != nil {
			log.Fatal(err)
		}
		subdomainVisitedcounts[strings.Join(subDomains[i:], ".")] += visited_count
	}
}
