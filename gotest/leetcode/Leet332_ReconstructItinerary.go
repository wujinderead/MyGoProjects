package main

import (
	"fmt"
	"sort"
)

// Given a list of airline tickets represented by pairs of departure and arrival 
// airports [from, to], reconstruct the itinerary in order. All of the tickets belong 
// to a man who departs from JFK. Thus, the itinerary must begin with JFK. 
// Note: If there are multiple valid itineraries, you should return the itinerary that
// has the smallest lexical order when read as a single string. For example, the itinerary 
// ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"]. 
// All airports are represented by three capital letters (IATA code). 
// You may assume all tickets form at least one valid itinerary. 
// Example 1: 
//   Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
//   Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]
// Example 2: 
//   Input: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
//   Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
//   Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"].
//     But it is larger in lexical order.

func findItinerary(tickets [][]string) []string {
    graph := make(map[string][]string, len(tickets)/2)
    tickettime := make(map[string]int)
    for i := range tickets {
    	graph[tickets[i][0]] = append(graph[tickets[i][0]], tickets[i][1])
    	tickettime[tickets[i][0]+tickets[i][1]] = tickettime[tickets[i][0]+tickets[i][1]]+1
    }
    for k := range graph {
    	sort.Sort(sort.StringSlice(graph[k]))
    }
    //fmt.Println(graph)
    //fmt.Println(tickettime)

    buf := make([]string, 0, len(tickets)+1)
    visit(len(tickets)+1, "JFK", graph, tickettime, &buf)
    return buf
}

func visit(n int, city string, graph map[string][]string, tickettime map[string]int, buf *[]string) bool {
	*buf = append(*buf, city)
	if len(*buf)==n {
		return true
	}
	for _, v := range graph[city] {
		key := city+v
		if t := tickettime[key]; t>0 {
			tickettime[key] = t-1
			if visit(n, v, graph, tickettime, buf) {
				return true
			}
			tickettime[key] = t
		}
	}
	*buf = (*buf)[:len(*buf)-1]
	return false
}

func main() {
	fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}))
	fmt.Println(findItinerary([][]string{{"EZE","AXA"},{"TIA","ANU"},{"ANU","JFK"},{"JFK","ANU"},
		{"ANU","EZE"},{"TIA","ANU"},{"AXA","TIA"},{"TIA","JFK"},{"ANU","TIA"},{"JFK","TIA"}}))
	fmt.Println(findItinerary([][]string{{"JFK","KUL"},{"JFK","NRT"},{"NRT","JFK"}}))
}