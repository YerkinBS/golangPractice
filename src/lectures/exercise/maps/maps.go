//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServer(serverStatus map[string]int) {
	fmt.Println("Number of servers:", len(serverStatus))
	
	onlineServCnt, offlineServCnt, maintServCnt, retiredServCnt := 0, 0, 0, 0
	for _, status := range serverStatus {
		switch status {
		case Online:
			onlineServCnt += 1
		case Offline:
			offlineServCnt += 1
		case Maintenance:
			maintServCnt += 1
		case Retired:
			retiredServCnt += 1
		}
	}

	fmt.Println("Number of online servers:", onlineServCnt)
	fmt.Println("Number of offline servers:", offlineServCnt)
	fmt.Println("Number of servers on maintenance:", maintServCnt)
	fmt.Println("Number of retired servers:", retiredServCnt)
	fmt.Println()
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatus := make(map[string]int)

	for i := 0; i < len(servers); i++ {
		serverStatus[servers[i]] = Online
	}

	displayServer(serverStatus)
	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline
	displayServer(serverStatus)

	for key, _ := range serverStatus {
		serverStatus[key] = Maintenance
	}

	displayServer(serverStatus)
}
