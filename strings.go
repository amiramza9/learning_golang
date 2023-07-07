package main
import ("fmt"
		"strings"
	)

func main(){
	pl := fmt.Println
	name := "amir hamza" //first time assignment with ":="
	pl(name)
	replc := strings.NewReplacer("i","ee")
	name =replc.Replace(name) // re-assignment with "="
	pl(name)
	pl("Length: ",len(name))
	pl("Contains mee", strings.Contains(name,"mee"))
	pl("index of 'z' is ",strings.Index(name,"z"))
	pl("Replace e with 3",strings.Replace(name,"e","3",1))
	pl("Replace e with 3",strings.Replace(name,"e","3",-1))

	myname:="Ameer\t   Hamza\t   GIKI"
	pl(myname)
	name=strings.TrimSpace(myname)
	pl(name)
	pl("Lower case: ",strings.ToLower(myname))
	pl("Upper case: ",strings.ToUpper(myname))
	arr := strings.Split("a b c d t x z o p"," ")
	pl("Split: ",arr)
	pl("Length of arr",len(arr))

	pl("Has the given prefix: ",strings.HasPrefix(myname,"Amee"))
	pl("Has the given Suffix: ",strings.HasSuffix(myname,"giki"))



}