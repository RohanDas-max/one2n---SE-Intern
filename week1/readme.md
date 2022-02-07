# Grep
   ### commands/methods
   *  search in STDIN values ` go run . ["value"]` then enter values (can enter multiple values) then hit `ctrl+c`. The terminal will print the matched values.
      * It is defined in the `searchFromStdin.go`. The method needs arg parameter to run. then method will enter a for loop where the bufio.newScanner will be invoked which is reading from the STDIN it will be keep looking for values until it is stopped manually (`ctrl+c`) then the values retrived from STDIN it will be stored in a slice then another method Checking will be invoked where the method is checking wheater the []string contains the argument or not if it is present it will return the matched values else the methods will return error

   * search in a file/directory `go run . ["value"] [filename/directory name]`
      * The method search in search.go is invoked when you enter the above command it will read the file which you entered or read the files which are available in the directory you entered. Then the `filepath.walk()` will be invoked it will check error if existed then it will return error if not then os.readfile will invoked which is taking the `path` parameter which is provided by `filepath.walk()` then a `[]byte`  recieved and stored one by one inside a `map`.Then checking if the argument is present in the map by iterating the `map`. then returning the available strings and errors.
   * storing the outpus in a file `go run . ["value"] [filename/directory name] -o [filename]`.
      * This is defined in `oflag()` method it will this method is reusing the search() if the search method return a value it will store it in the filename provided in the command.
   * All of these methods are being called in the `core()` where the method is checking the length of the argument and according the commands length the methods are running.


# Josephus Simulation
   ### commands/method
   *  `go run . -n [number] -k [number]
      * Here the -n represents the numbers of players participated
      * -k represents the numbers of skips
   * In the `controller.go` the function controller holds the logic to find out the survivor/winner of the game it is a recursive which is first checking if the number of player is more then 1 if not then the 1 is the only player so he's the winner else the recursive function will be called and eventually will find out the winner.
   * In the `flags.go` the function holds the the falgs operations where the -n & -k flags are invoked and the the values of n and k are returned
   * In the `main.go` there are two method one is `exec()` which is executing the `flags()` and the values of n and k are used to execute the `controller.go` which is returning the position of the winner.