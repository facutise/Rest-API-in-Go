package main

import "os"

//"fmt"

func checkError(err error) {
	if err != nil {
		panic(err)

	}
}

func main() {

	//err:= os.Mkdir("subdir", 0755)
	//checkError(err)
	checkError(os.Mkdir("subdir1", 0755))

	//También se pueden remover directories
	//defer os.RemoveAll("subdir1")

	os.WriteFile("subdir1/file", []byte(""), 0755)

}
