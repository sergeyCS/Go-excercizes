Go syntax cheatpaper.

Function make produces the same slice as allocating an array and slicing it:
make([]int, 50, 100)
new([100]int)[0:50]