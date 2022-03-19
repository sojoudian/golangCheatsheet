How do you copy a slice, a map, and an interface?



You copy a slice by using the built-in copy() function:



a := []int{1, 2}

b := []int{3, 4}

check := a

copy(a, b)

fmt.Println(a, b, check)

// Output: [3 4] [3 4] [3 4]

Here, the check variable is used to hold a reference to the original slice description to make sure it is really copied.



In the next example, on the other hand, operation does not copy the slice contents, only the slice description:



a := []int{1, 2}

b := []int{3, 4}

check := a

a = b

fmt.Println(a, b, check)

// Output: [3 4] [3 4] [1 2]

You copy a map by traversing its keys. Yes, unfortunately, this is the simplest way to copy a map in Go:



a := map[string]bool{"A": true, "B": true}

b := make(map[string]bool)

for key, value := range a {

b[key] = value

}

Following example copies just the description of the map:



a := map[string]bool{"A": true, "B": true}

b := map[string]bool{"C": true, "D": true}

check := a

a = b

fmt.Println(a, b, check)

// Output: map[C:true D:true] map[C:true D:true] map[A:true B:true]

There’s no built-in way in Go to copy an interface. No, the reflect.DeepCopy() function does not exist