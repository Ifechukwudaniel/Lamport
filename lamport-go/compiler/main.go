package main

import (
	"fmt"
	//"os"
	"os/user"
)

func main() {
	const LAMPORT_COMPILER = `
      ___       ___           ___           ___           ___           ___           ___     
     /\__\     /\  \         /\__\         /\  \         /\  \         /\  \         /\  \    
    /:/  /    /::\  \       /::|  |       /::\  \       /::\  \       /::\  \        \:\  \   
   /:/  /    /:/\:\  \     /:|:|  |      /:/\:\  \     /:/\:\  \     /:/\:\  \        \:\  \  
  /:/  /    /::\~\:\  \   /:/|:|__|__   /::\~\:\  \   /:/  \:\  \   /::\~\:\  \       /::\  \ 
 /:/__/    /:/\:\ \:\__\ /:/ |::::\__\ /:/\:\ \:\__\ /:/__/ \:\__\ /:/\:\ \:\__\     /:/\:\__\
 \:\  \    \/__\:\/:/  / \/__/~~/:/  / \/__\:\/:/  / \:\  \ /:/  / \/_|::\/:/  /    /:/  \/__/
  \:\  \        \::/  /        /:/  /       \::/  /   \:\  /:/  /     |:|::/  /    /:/  /     
   \:\  \       /:/  /        /:/  /         \/__/     \:\/:/  /      |:|\/__/     \/__/      
    \:\__\     /:/  /        /:/  /                     \::/  /       |:|  |                  
     \/__/     \/__/         \/__/                       \/__/         \|__|


_________  ________      _____   __________ .___ .____     _____________________ 
\_   ___ \ \_____  \    /     \  \______   \|   ||    |    \_   _____/\______   \
/    \  \/  /   |   \  /  \ /  \  |     ___/|   ||    |     |    __)_  |       _/
\     \____/    |    \/    Y    \ |    |    |   ||    |___  |        \ |    |   \
 \______  /\_______  /\____|__  / |____|    |___||_______ \/_______  / |____|_  /
        \/         \/         \/                         \/        \/         \/ 

`
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("%s \n", LAMPORT_COMPILER)
	fmt.Printf("Hello %s! This is the Lamport compiler writing in go!\n", user.Username)
	fmt.Printf("Fell free to type in commands \n")
	//fmt.Printf(os.Stdin, os.Stdout)
}
