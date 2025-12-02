const fs = require("fs")

fs.readFile("../../AoC-input/2022/12/input.txt", (err, input) => {
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
})

fs.readFile("../../AoC-input/2022/12/input.txt", (err, input) => {
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
})
