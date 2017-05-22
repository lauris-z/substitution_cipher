package main

import (
  "fmt"
  "strings"
  "os"
  "strconv"
)

var COLONNES=17

func main() {


 TAB:=[]string{"H","E","R",">","!P","!L","^","V","P","!K","|","OB","L","T","G","OH","!D",
	       "N","!P","+","B","O|","sqrF","O","sqrD","D","W","Y",".","<","sqrG","K","!F","O-",
	       "B","!Y","I.","!C","M","+","U","Z","G","W","O|","O-","L","sqrF","O+","H","J",
	       "S","!P","!P","tri.","^","!L","triF","sqrG","V","OD","!P","O","+","+","R","K","OH",
	       "sqr","tri","M","+","O+","!T","!J","!D","|","OO","F","P","+","PP","OG","!K","/",
               "!P","triF","R","^","F","!L","O","-","sqrG","!D","C","!K","F",">","OH","D","O|",
	       "sqrF","OO","+","K","Q","sqrD",".I","OH","U","!C","X","G","V",".","O+","L","|",
	       "O|","G","OH","J","!F","!J","sqrF","O","+","sqr","N","Y","O+","+","sqr.","L","tri",
	       "!D","<","M","+","!B","+","Z","R","OH","F","B","!C","!Y","A","O.","OG","K",
	       "-","O+","!L","U","V","+","^","J","+","O","!P","tri.","<","F","B","!Y","-",
	       "U","+","R","/","OO","!T","E","|","D","Y","B","!P","!B","T","M","K","O",
	       "OH","<","!C","!L","R","J","|","sqrG","OO","T","OG","M",".","+","PP","B","F",
	       "O+","O.","tri","S","!Y","sqrF","+","N","|","OO","F","B","!C","O|",".I","triF","R",
	       "!L","G","F","N","^","!F","OO","OH","OG","!B",".","!C","V","OG","!T","+","+",
	       "!Y","B","X","OB","sqrG","I.","OG","tri","C","E",">","V","U","Z","OO","-","+",
	       "|","!C",".","OD","O+","B","K","O|","O","!P","^",".","!F","M","Q","G","OH",
	       "R","!C","T","+","L","OB","O.","C","<","+","F","!L","W","B","|","O-","L",
	       "+","+","O-","W","C","O+","W","!C","P","O","S","H","T","/","O|","O-","!P",
	       "|","F","!K","!D","W","<","tri.","!T","B","sqr","Y","O","B","sqrG","-","C","!C",
	       ">","M","D","H","N","!P","!K","S","O+","Z","O","triF","A","|","K",".I","+"}

	TAB=tabDigit(TAB)
	printTabString(TAB,COLONNES)
	TAB=replaceTabString(TAB,os.Args[1])
	printTabString(TAB,COLONNES)
}

func printTabString(tab []string, colonnes int) {
	for i:=0;i<len(tab);i++ {
                fmt.Print(tab[i])
		fmt.Print(" ")
		if((i+1)%colonnes==0) {
			fmt.Print("\n")
		}
        }
}

func containsCar(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func replaceTabString(tab []string, modifs string) []string {
        newTab:=tab
        dico:=strings.Split(modifs,",")
        for i:=0;i<len(dico);i++ {
                tabModif:=strings.Split(dico[i],"=")
                for k:=0;k<len(tab);k++ {
                        value:=tab[k]
                        if value == tabModif[0] {
                                newTab[k]=tabModif[1]
                        }
                }
        }
        tabAlphabet:=[]string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
        for i:=0;i<len(newTab);i++ {
                if !containsCar(tabAlphabet,newTab[i]) {
                        newTab[i]="?"
                }
        }
        return newTab
}

func tabDigit(tab []string) []string {
	var SYMBOLS map[string]string
	SYMBOLS=make(map[string]string)
	last:=1
	for i:=0;i<len(tab);i++ {
		if(SYMBOLS[tab[i]]=="") {
			if(last<10){
				SYMBOLS[tab[i]]="0"+strconv.Itoa(last)
			} else {
				SYMBOLS[tab[i]]=strconv.Itoa(last)
			}
			last+=1
		}
	}
	for i:=0;i<len(tab);i++ {
		tab[i]=SYMBOLS[tab[i]]
	}
	return tab
}
