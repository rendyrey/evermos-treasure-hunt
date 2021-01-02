# evermos-treasure-hunt

- `#` represents an obstacle. 
- `.` represents a clear path. 
- `X` represents the player’s starting position. 
  
A treasure is hidden within one of the clear path points, and the user must find it.

From the starting position, the user must navigate in a specific order: 

- Up/North A step(s)
- then Right/East B step(s)
- then Down/South C step(s). 

- So the first move have to be up (can move up more than 1 time)
- the second move have to be right (can move right more than 1 time)
- and the third move have to down (can move down more than 1 time)

The program must output a list of probable coordinate points where the treasure might be located.
display the grid with the probable treasure locations marked with a `$` symbol.


# how to run
`go run main.go`

