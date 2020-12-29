package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(input));
	fmt.Println(alternateSolution(input));
}

//*************************************** Solution 1 ***************************************//
func longestCommonPrefix(strs []string) string {    
	if(len(strs) == 0) {
			return ""
	}    
	prefix := strs[0]        
	
	for i:=1; i<len(strs); i++ {
			for(strings.Index(strs[i], prefix) != 0) {
					prefix = prefix[0:len(prefix)-1]  
					if(prefix == "") {
							return ""
					}
			}               
	}
	return prefix
}

//*************************************** Solution 2 ***************************************//

func alternateSolution(strs []string) string {    
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

//********************************************************************************************//