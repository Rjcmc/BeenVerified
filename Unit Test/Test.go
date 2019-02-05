package main


func main() {
	
    CheckErr(nil)
}



func CheckErr(err error) int{
        if err != nil {
            //panic(err)
	    return 1
        }
	return 0
    }
