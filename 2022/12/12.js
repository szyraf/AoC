/*
Description:

--- Day 12: Hill Climbing Algorithm ---
You try contacting the Elves using your handheld device, but the river you're following must be too low to get a decent signal.

You ask the device for a heightmap of the surrounding area (your puzzle input). The heightmap shows the local area from above broken into a grid; the elevation of each square of the grid is given by a single lowercase letter, where a is the lowest elevation, b is the next-lowest, and so on up to the highest elevation, z.

Also included on the heightmap are marks for your current position (S) and the location that should get the best signal (E). Your current position (S) has elevation a, and the location that should get the best signal (E) has elevation z.

You'd like to reach E, but to save energy, you should do it in as few steps as possible. During each step, you can move exactly one square up, down, left, or right. To avoid needing to get out your climbing gear, the elevation of the destination square can be at most one higher than the elevation of your current square; that is, if your current elevation is m, you could step to elevation n, but not to elevation o. (This also means that the elevation of the destination square can be much lower than the elevation of your current square.)

For example:

Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
Here, you start in the top-left corner; your goal is near the middle. You could start by moving down or right, but eventually you'll need to head toward the e at the bottom. From there, you can spiral around to the goal:

v..v<<<<
>v.vv<<^
.>vv>E^^
..v>>>^^
..>>>>>^
In the above diagram, the symbols indicate whether the path exits each square moving up (^), down (v), left (<), or right (>). The location that should get the best signal is still E, and . marks unvisited squares.

This path reaches the goal in 31 steps, the fewest possible.

What is the fewest steps required to move from your current position to the location that should get the best signal?
*/

const fs = require("fs")

// it takes about 40s
fs.readFile("input.txt", (err, input) => {
    if (err) throw err
    let data = input.toString()
    let lines = data.replaceAll('\r', '').split('\n')

    let grid = []
    
    lines.forEach(line => {
        grid.push(line.split('').map(char => char.charCodeAt(0) - 97))
    })

    let minSteps = grid.flat().length
    let minStepsGrid = Array(grid.length).fill().map(() => Array(grid[0].length).fill(minSteps))

    function check(occupiedCoordinates, direction) {
        let thisX = occupiedCoordinates[occupiedCoordinates.length - 2]
        let thisY = occupiedCoordinates[occupiedCoordinates.length - 1]
 
        let x = 0, y = 0
        if      (direction === 0) y--
        else if (direction === 1) x++
        else if (direction === 2) y++
        else if (direction === 3) x--

        let isFinish = false

        if (!(thisX + x < 0 || thisX + x >= grid[0].length || thisY + y < 0 || thisY + y >= grid.length)) {
            if (grid[thisY + y][thisX + x] === -28) {
                if (grid[thisY][thisX] >= 24) {
                    isFinish = true
                    if (occupiedCoordinates.length / 2 < minSteps) {
                        minSteps = occupiedCoordinates.length / 2
                    }
                }
            }
            else if (grid[thisY + y][thisX + x] <= (grid[thisY][thisX] === -14 ? 1 : (grid[thisY][thisX] + 1))) {
                let isOccupied = false
                for (let i = 0; i < occupiedCoordinates.length; i += 2) {
                    if (occupiedCoordinates[i] === thisX + x && occupiedCoordinates[i + 1] === thisY + y) {
                        isOccupied = true
                    }
                }
                if (!isOccupied) {
                    if (occupiedCoordinates.length / 2 < minStepsGrid[thisY + y][thisX + x]) {
                        minStepsGrid[thisY + y][thisX + x] = occupiedCoordinates.length / 2

                        let copyOccupiedCoordinates = [...occupiedCoordinates]
                        copyOccupiedCoordinates.push(thisX + x)
                        copyOccupiedCoordinates.push(thisY + y)
                        check(copyOccupiedCoordinates, 0)
                    }
                }
            }
        }

        if (direction !== 3 && !isFinish) {
            check(occupiedCoordinates, direction + 1)
        }
    }

    let startIndex = grid.flat().indexOf(-14)
    minStepsGrid[startIndex / grid[0].length, startIndex % grid[0].length] = 0
    check([startIndex % grid[0].length, startIndex / grid[0].length], 0)

    console.log(minSteps);
    // 481
})

/*
--- Part Two ---
As you walk up the hill, you suspect that the Elves will want to turn this into a hiking trail. The beginning isn't very scenic, though; perhaps you can find a better starting point.

To maximize exercise while hiking, the trail should start as low as possible: elevation a. The goal is still the square marked E. However, the trail should still be direct, taking the fewest steps to reach its goal. So, you'll need to find the shortest path from any square at elevation a to the square marked E.

Again consider the example from above:

Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
Now, there are six choices for starting position (five marked a, plus the square marked S that counts as being at elevation a). If you start at the bottom-left square, you can reach the goal most quickly:

...v<<<<
...vv<<^
...v>E^^
.>v>>>^^
>^>>>>>^
This path reaches the goal in only 29 steps, the fewest possible.

What is the fewest steps required to move starting from any square with elevation a to the location that should get the best signal?
*/

// it takes about 40s
fs.readFile("input.txt", (err, input) => {
    if (err) throw err
    let data = input.toString()
    let lines = data.replaceAll('\r', '').split('\n')

    let grid = []
    
    let minOccupiedCoordinates = []

    lines.forEach(line => {
        grid.push(line.split('').map(char => char.charCodeAt(0) - 97))
    })

    let minSteps = grid.flat().length
    let minStepsGrid = Array(grid.length).fill().map(() => Array(grid[0].length).fill(minSteps))

    function check(occupiedCoordinates, direction) {
        let thisX = occupiedCoordinates[occupiedCoordinates.length - 2]
        let thisY = occupiedCoordinates[occupiedCoordinates.length - 1]
 
        let x = 0, y = 0
        if      (direction === 0) y--
        else if (direction === 1) x++
        else if (direction === 2) y++
        else if (direction === 3) x--

        let isFinish = false

        if (!(thisX + x < 0 || thisX + x >= grid[0].length || thisY + y < 0 || thisY + y >= grid.length)) {
            if (grid[thisY + y][thisX + x] === -28) {
                if (grid[thisY][thisX] >= 24) {
                    isFinish = true
                    if (occupiedCoordinates.length / 2 < minSteps) {
                        minSteps = occupiedCoordinates.length / 2
                        minOccupiedCoordinates = [...occupiedCoordinates]
                    }
                }
            }
            else if (grid[thisY + y][thisX + x] <= (grid[thisY][thisX] === -14 ? 1 : (grid[thisY][thisX] + 1))) {
                let isOccupied = false
                for (let i = 0; i < occupiedCoordinates.length; i += 2) {
                    if (occupiedCoordinates[i] === thisX + x && occupiedCoordinates[i + 1] === thisY + y) {
                        isOccupied = true
                    }
                }
                if (!isOccupied) {
                    if (occupiedCoordinates.length / 2 < minStepsGrid[thisY + y][thisX + x]) {
                        minStepsGrid[thisY + y][thisX + x] = occupiedCoordinates.length / 2

                        let copyOccupiedCoordinates = [...occupiedCoordinates]
                        copyOccupiedCoordinates.push(thisX + x)
                        copyOccupiedCoordinates.push(thisY + y)
                        check(copyOccupiedCoordinates, 0)
                    }
                }
            }
        }

        if (direction !== 3 && !isFinish) {
            check(occupiedCoordinates, direction + 1)
        }
    }

    let startIndex = grid.flat().indexOf(-14)
    minStepsGrid[startIndex / grid[0].length, startIndex % grid[0].length] = 0
    check([startIndex % grid[0].length, startIndex / grid[0].length], 0)

    console.log(minOccupiedCoordinates);
    for (let i = 2; i < minOccupiedCoordinates.length; i += 2) {
        if (minOccupiedCoordinates[i] === 0) {
            minSteps--
        }
        else {
            break
        }
    }

    console.log(minSteps);
    // 480
})