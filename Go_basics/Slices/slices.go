package main
import ("fmt")


func main(){
	
    // 1) creating a slice using []datatype{values} method

    sl1 := []int{}

    fmt.Println("Size of sl1 = ", len(sl1))


    // 2) creating slice from an array

    arr:= [...]int{24, 53, 1, 49, 91, 77}

    sl2:= arr[2:5]      // arr[2] to arr[5-1]

    fmt.Println(sl2)


    // 3) creating slice using make() function 

    sl3:= make([]int, 10, 20)
    sl4:= make([]int, 10)

    fmt.Println(sl3, "Length = ", len(sl3), "Capacity = ", cap(sl3))
    fmt.Println("Capacity of sl4 = ", cap(sl4))


    // ---------------------------------------------

    // Modifying slices


    // ❌ we cannot directly insert items to the slice like this; it will give runtime error saying index is out of range

    // for i:=0; i<10; i++{
    //     sl1[i]=i
    // }

    // fmt.Println(sl1)


    // ✅ We use append() to add an element to the end of the slice

    sl5:=[]int{3, 4}

    fmt.Println("Slice before appending: ", sl5)

    sl5 = append(sl5, 45, 12, 0, 1)

    fmt.Println("Slice after appending some elements: ", sl5)


    // appending one slice to another
    sl6:= append(sl5, sl2...)

    fmt.Println("Sl6 = ", sl6)


    sl7:=sl6[:4]
    fmt.Println(sl7)



    // Implementing copy() function

    arr2:=[]int{2, 52, 67, 13, 83, 14, 67, 95}

    // creating a smaller array from elements of arr2, which is needed to be used somewhere

    neededNums:=arr2[:5]    // although this new slice contains the desired elements, it is still dependent on the original array

    copyNums:=make([]int, len(neededNums))
    copy(copyNums, neededNums)      // the copyNums is now independent of the original array/slice and the original array can be garbage collected if not needed

    fmt.Print(copyNums)

    






}