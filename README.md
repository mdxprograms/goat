# GOAT (go + react + binary style builds)

This is a working example of using `Go` + `React` style app as a binary build using `go.rice`

Once it is packaged you can run the binary for whichever OS the build is packaged for
using `./goat` for this example or whatever the app name is given in the makefile

### Run local dev

`make dev`

### Run build

`make build`

### Run package (build react app + build go app + rice to binary all_the_things)

`make package`

### Serve binary

`./goat`
