# GO Exercises

GO Exercise is a library containing basic go programs covering all the basic concepts. This repository is a collection of exercise problems, and solutions.

### exercise_01.go

 * Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS ```“x”```, ```“y”```and ```“z”```

   * ```42```
     
   * ```“James Bond”```

   * ```true```

 * Now print the values stored in those variables using 

  * single print statement

  * multiple print statements. What are these values called?

### exercise_02.go

 * Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).

  * identifier ```“x” type int```

  * identifier ```“y” type string```

  * identifier ```“z” type bool```

 * in func main

  * print out the values for each identifier

  * The compiler assigned values to the variables. What are these values called?

### exercise_03.go

Using the code from the previous exercise,

 * At the package level scope, assign the following values to the three variables

 * for ```x``` assign ```42```

 * for ```y``` assign ```“James Bond”```

 * for ```z assign ```true..```

 * in ```func main```

  * use ```fmt.Sprintf``` to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER ```“s”```

  * print out the value stored by variable ```“s”```

### exercise_04.go

For this exercise,

 * Create your own type. Have the underlying type be an int.

 * create a VARIABLE of your new TYPE with the IDENTIFIER ```“x”```using the ```“VAR”``` keyword

 * in ```func main```

  * print out the value of the variable ```“x”```

  * print out the type of the variable ```“x”```

  * assign ```42``` to the VARIABLE ```“x”``` using the ```“=”``` OPERATOR

  * print out the value of the variable ```“x”```


### exercise_05.go

Building on the code from the previous example,

 * at the package level scope, using the ```“var”``` keyword, create a VARIABLE with the IDENTIFIER ```“y”```. The variable should be of the UNDERLYING TYPE of your custom TYPE ```“x”```

  * eg:

```golang
type hotdog int
var x hotdog
var y int
```
in func main,

  * this should already be done

   * print out the value of the variable ```“x”```

   * print out the type of the variable ```“x”```

   * assign your own VALUE to the VARIABLE ```“x”``` using the ```“=”```OPERATOR

   * print out the value of the variable ```“x”```

 * now do this

  * now use CONVERSION to convert the TYPE of the VALUE stored in ```“x”``` to the UNDERLYING TYPE

   * then use the ```“=”``` operator to ASSIGN that value to ```“y”```

   * print out the value stored in ```“y”```

   * print out the type of ```“y”```

### exercise_06.go

Write a program that prints a number in ```decimal, binary, and hex```

### exercise_07.go

Using the following operators, write expressions and assign their values to variables:

```golang
 * ==

 * <=

 * >=

 * !=

 * <

 * >
```

Now print each of the variables.

### exercise_08.go

Create TYPED and UNTYPED constants. Print the values of the constants.

### exercise_09.go

Write a program that

* assigns an int to a variable

* prints that int in ```decimal, binary, and hex```

* shifts the bits of that int over 1 position to the 
left, and assigns that to a variable

* prints that variable in ```decimal, binary, and hex```

### exercise_10.go

Create a variable of type string using a raw string literal . Print it.

### exercise_11.go
Using ```iota``` , create 4 constants for the NEXT 4 years. Print the constant values.

### exercise_12.go

Print every number from ```1 to 10,000```

### exercise_13.go

Print every rune code point of the uppercase alphabet three times. Your output should look like
this:

```golang
65

U+0041 'A'

U+0041 'A'

U+0041 'A'

66

U+0042 'B'

U+0042 'B'

U+0042 'B'
```

through the rest of the alphabet characters

### exercise_14.go

Create a for loop using this syntax

```golang
* for condition { }
```
Have it print out the years you have been alive.

### exercise_15.go

Create a for loop using this syntax

```golang
* for { }
```

Have it print out the years you have been alive.

### exercise_16.go

Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

### exercise_17.go

Create a program that shows the ```“if statement”``` in action.

### exercise_18.go

Building on the previous hands-on exercise, create a program that uses ```“else if” and “else”```.

### exercise_19.go

Create a program that uses a switch statement with no switch expression specified.

### exercise_20.go

Create a program that uses a switch statement with the switch expression specified as a
variable of TYPE string with the IDENTIFIER “favSport”.

### exercise_21.go

Write down what these print:

```golang
* fmt.Println( true && true )
* fmt.Println( true && false )
* fmt.Println( true || true )
* fmt.Println( true || false )
* fmt.Println( !true )
```

### exercise_22.go

Using a COMPOSITE LITERAL:

 * create an ARRAY which holds 5 VALUES of TYPE int

 * assign VALUES to each index position.

 * Range over the array and print the values out.

 * Using format printing

  * print out the TYPE of the array

### exercise_23.go

Using a COMPOSITE LITERAL:

 * create a SLICE of TYPE int

 * assign 10 VALUES

 * Range over the slice and print the values out.

 * Using format printing

  * print out the TYPE of the slice

### exercise_24.go

Using the code from the previous example, use SLICING to create the following new slices which are then printed:

```golang
* [42 43 44 45 46]

* [47 48 49 50 51]

* [44 45 46 47 48]

* [43 44 45 46 47]
```

### exercise_25.go

Follow these steps:

 * start with this slice

```golang
 * x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
```

* append to that slice this value

 * 52

* print out the slice

* in ONE STATEMENT append to that slice these values

 * 53

 * 54

 * 55

* print out the slice

* append to the slice this slice

```golang
 * y := []int{56, 57, 58, 59, 60}
```

* print out the slice

### exercise_26.go

To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:

* start with this slice

```golang
 * x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
````

* use APPEND & SLICING to get these values here which you should ASSIGN to a variable ```“y”``` and then print:

```golang
 * [42, 43, 44, 48, 49, 50, 51]
```

### exercise_27.go

Create a slice to store the names of all of the states in the United States of America. What is the
length of your slice? What is the capacity? Print out all of the values, along with their index
position in the slice, without using the range clause. Here is a list of the states:

```golang
`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, ` Colorado`, ` Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`,`Idaho`, `Illinois`, ` Indiana`, ` Iowa`, `Kansas`,`Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`,`Minnesota`,`Mississippi`, `Missouri`, `Montana`, `Nebraska`,`Nevada`, ` New Hampshire`, ` New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, `Pennsylvania`, ` Rhode Island`, ` South Carolina`, `South Dakota`, ` Tennessee`, ` Texas`, `Utah`, ` Vermont`, ` Virginia`, ` Washington`, `West Virginia`, ` Wisconsin`, ` Wyoming`,
```

### exercise_28.go

Create a slice of a slice of string (```[][]string```). Store the following data in the multi-dimensional
slice:

```golang
"James", "Bond", "Shaken, not stirred"

"Miss", "Moneypenny", "Helloooooo, James."
```

Range over the records, then range over the data in each record.

### exercise_29.go

Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE ```[]string``` which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.

```golang
`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`

`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`

`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
```

### exercise_30.go

Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

### exercise_31.go

Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop

### exercise_32.go

Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:

 * first name
 * last name
 * favorite ice cream flavors

Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

### exercise_33.go

Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

### exercise_34.go

Create a new type: vehicle.

 * The underlying type is a struct.

 * The fields:

  * doors

  * color

 * Create two new types: truck & sedan.

  * The underlying type of each of these new types is a struct.

  * Embed the “vehicle” type in both truck & sedan.

  * Give truck the field “fourWheel” which will be set to bool.

  * Give sedan the field “luxury” which will be set to bool. solution

 * Using the vehicle, truck, and sedan structs:

  * using a composite literal, create a value of type truck and assign values to the
fields;

  * using a composite literal, create a value of type sedan and assign values to the
fields.

 * Print out each of these values.

 * Print out a single field from each of these values.

### exercise_34.go

Create and use an anonymous struct.

### exercise_37.go

 * create a func with the identifier foo that returns an int

 * create a func with the identifier bar that returns an int and a string

 * call both funcs

 * print out their results

### exercise_38.go

create a func with the identifier foo that

 * takes in a variadic parameter of type int

  * pass in a value of type ```[]int``` into your func (unfurl the ```[]int```)

  * returns the sum of all values of type int passed in

 * create a func with the identifier bar that

  * takes in a parameter of type ```[]int```

  * returns the sum of all values of type int passed in

### exercise_39.go

Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

### exercise_40.go

 * Create a user defined struct with

  * the identifier “person”

  * the fields:

   * first

   * last

   * age

 * attach a method to type person with

  * the identifier “speak”

  * the method should have the person say their name and age

 * create a value of type person

 * call the method from the value of type person

### exercise_41.go

 * create a type SQUARE

 * create a type CIRCLE

 * attach a method to each that calculates AREA and returns it


  * circle ```area= πr2```

  * square ```area = L * W```

 * create a type SHAPE that defines an interface as anything that has the AREA method

 * create a func INFO which takes type shape and then prints the area

 * create a value of type square

 * create a value of type circle

 * use func info to print the area of square

 * use func info to print the area of circle

### exercise_42.go

Build and use an anonymous func

### exercise_43.go

Assign a func to a variable, then call that func

### exercise_44.go

 * Create a func which returns a func

 * assign the returned func to a variable

 * call the returned func

### exercise_45.go

“callback” is when we pass a func into a func as an argument. For this exercise,

 * pass a func into a func as an argument

### exercise_46.go

Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable

### exercise_47.go

Create a value and assign it to a variable.

 * Print the address of that value.

### exercise_48.go

 * create a person struct

 * create a func called “changeMe” with a ```*person``` as a parameter

 * change the value stored at the ```*person``` address

  * important

 * to dereference a struct, use ```(*value).field```

  * ```p1.first```

  * ```(*p1).first```

 * “As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
x.f is shorthand for (*x).f.”

  * [https://golang.org/ref/spec#Selectors](https://golang.org/ref/spec#Selectors)

 * create a value of type person

  * print out the value

 * in ```func main```

  * call “changeMe”

 * in ```func main```

  * print out the value

### exercise_51.go

Starting [with this code](https://play.golang.org/p/_fQUGm9Utvl) , marshal the ```[]user``` to JSON. There is a little bit of a curve ball in this
hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.

### exercise_52.go

Starting [with this code](https://play.golang.org/p/b_UuCcZag9) , unmarshal the JSON into a Go data structure. 

Hint: [https://mholt.github.io/json-to-go/](https://play.golang.org/p/b_UuCcZag9)

### exercise_53.go

Starting [with this code](https://play.golang.org/p/BVRZTdlUZ_) , encode to JSON the ```[]user``` sending the results to ```Stdout```.

Hint: you will need to use ```json.NewEncoder(os.Stdout).encode(v interface{})```

### exercise_54.go

Starting [with this code](https://play.golang.org/p/H_q75mpmHW) , sort the ```[]int``` and ```[]string``` for each person.

### exercise_55.go

Starting [with this code](https://play.golang.org/p/BVRZTdlUZ_), sort the ```[]user``` by

 * age

 * last

Also sort each ```[]string``` “Sayings” for each user

 * print everything in a way that is pleasant

### exercise_56.go

 * in addition to the main goroutine, launch two additional goroutines

  * each additional goroutine should print something out

  * use waitgroups to make sure each goroutine finishes before your program exists

### exercise_57.go

This exercise will reinforce our understanding of method sets:

 * create a type person struct

 * attach a method speak to type person using a pointer receiver

  * *person

 * create a type human interface

  * to implicitly implement the interface, a human must have the speak method

 * create func “saySomething”

  * have it take in a human as a parameter

  * have it call the speak method

 * show the following in your code

  * you CAN pass a value of type *person into saySomething

  * you CANNOT pass a value of type person into saySomething

 * here is a hint if you need some help

### exercise_58.go

Using goroutines, create an incrementer program

 * have a variable to hold the incrementer value

 * launch a bunch of goroutines

  * each goroutine should

   * read the incrementer value

   * store it in a new variable

   * yield the processor with ```runtime.Gosched()```

   * increment the new variable

   * write the value in the new variable back to the incrementer variable

 * use waitgroups to wait for all of your goroutines to finish

 * the above will create a race condition.

 * Prove that it is a race condition by using the ```-race``` flag

 * if you need help, here is a hint: [https://play.golang.org/p/FYGoflKQej](https://play.golang.org/p/FYGoflKQej)

### exercise_59.go

Fix the race condition you created in the previous exercise by using a mutex

 * it makes sense to remove 

 ```golang
 runtime.Gosched()
```

### exercise_60.go

Fix the race condition you created in ```exercise_56.go``` by using 

```golang
package atomic
```
### exercise_61.go

Create a program that prints out your OS and ARCH. Use the following commands to run it

```golang
 * go run

 * go build

 * go install
```

### exercise_62.go

Get [this code](https://play.golang.org/p/j-EA6003P0) working:

 * using func literal, aka, anonymous self-executing func

 * a buffered channel

### exercise_63.go

Get this code running:

 * [https://play.golang.org/p/oB-p3KMiH6](https://play.golang.org/p/oB-p3KMiH6)

 * [https://play.golang.org/p/_DBRueImEq](https://play.golang.org/p/oB-p3KMiH6)

### exercise_64.go

Starting [with this code](https://play.golang.org/p/sfyu4Is3FG), pull the values off the channel using a for range loop

### exercise_65.go

Starting [with this code](https://play.golang.org/p/sfyu4Is3FG) , pull the values off the channel using a select statement

### exercise_66.go

Show the comma ok idiom starting with [this code](https://play.golang.org/p/YHOMV9NYKK)

### exercise_67.go

Write a program that

 * puts 100 numbers to a channel

 * pull the numbers off the channel and print them

### exercise_68.go

Write a program that

 * launches 10 goroutines

  * each goroutine adds 10 numbers to a channel

 * pull the numbers off the channel and print them

### exercise_69.go

Start [with this code](https://play.golang.org/p/3W69TH4nON). Instead of using the blank identifier, make sure the code is checking and
handling the error.

### exercise_70.go

Start [with this code](https://play.golang.org/p/9a1IAWy5E6) . Create a custom error message using “fmt.Errorf”.

### exercise_71.go

Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. If you need a hint, [here is one](https://play.golang.org/p/L5QWV8-p11).

### exercise_72.go

Starting [with this code](https://play.golang.org/p/wlEM1tgfQD) , use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your

```golang
 * lat "50.2289 N"

 * long "99.4656 W"
```

### exerciseseventythree

Create a dog package. The dog package should have an exported func “Years” which takes human years and turns them into dog years (1 human year ```=``` 7 dog years). Document your code with comments. Use this code in ```func main```.

 * run your program and make sure it works

 * run a local server with godoc and look at your documentation

### exercise_74.go

Use godoc at the command line to look at the documentation for:


 * fmt

 * fmt Print

 * strings

 * strconv

### exerciseseventyfive

Start [with this code](https://play.golang.org/p/S87WsSd4noR). Get the code ready to BET on the code (add benchmarks, examples, tests). Run the following in this order:

 * tests

 * benchmarks

 * coverage

 * coverage shown in web browser

 * examples shown in documentation in a web browser

### exerciseseventysix

Start [with this code](https://play.golang.org/p/lnGw07GfpxJ) . Get the code ready to BET on (add benchmarks, examples, tests) however
do not write an example for the func that returns a map; and only write a test for it as an extra challenge. Add documentation to the code. Run the following in this order:

 * tests

 * benchmarks

 * coverage

 * coverage shown in web browser

 * examples shown in documentation in a web browser

### exerciseseventyseven

Start [with this code](https://play.golang.org/p/MwZ1xAD5Z5s). Get the code ready to BET on (add benchmarks, examples, tests). Write a table test. Add documentation to the code. Run the following in this order:

* tests

 * benchmarks

 * coverage

 * coverage shown in web browser

 * examples shown in documentation in a web browser

### exercise\_callback\_35.go
### exercise\_factorial\_36.go
### exercise\_json\_marshal\_49.go
### exercise\_json\_unmarsha\_50.go

## godoc offline

If `$ godoc` does not work in command line or git bash run: 

|`$ go get -v  golang.org/x/tools/cmd/godoc`

To run local server to acces `godoc` offline run:

|`$ godoc -http=:[port number]`

To read `godoc` in browser:

|`localhost:[port number]`

## Testing

To test go file run:

`$ go test`

## Benchmarking

To benchmark go code run:

`go test -bench .`

## Coverage

To test coverage:

`go test -cover`

To print coverage to a file:

`go test -coverprofile cover_file_name.out`

To view coverage in browser:

`go tool cover -html=cover_file_name.out`
