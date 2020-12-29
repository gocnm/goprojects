package main

import "fmt"

func main() {
	input := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(input));
}

func longestCommonPrefix(strs []string) string {    
	prefix := ""
	
	if(len(strs) != 0) {
			prefix = strs[0]
	} 

	for i:=1; i<len(strs); i++ {                          
			prefix = getCommonPrefix(prefix, strs[i])			
			if(prefix == "") {
				break 				
			}
	}
	return prefix
}

func getCommonPrefix(prev string, current string) string {
	counter := 0

	for i:=0; (i<len(prev) && i<len(current)); i++ {
			if(prev[i] == current[i]) {
					counter++
			} else {
				return prev[0:counter]
			}
	}
	return prev[0:counter]
}