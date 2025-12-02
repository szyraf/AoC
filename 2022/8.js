const fs = require("fs")

fs.readFile("../../AoC-input/2022/8/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let forest = []
    let visibleTrees = []

    lines.forEach(line => {
        forest.push(line.replace('\r', '').split('').map(Number))
        visibleTrees.push(line.replace('\r', '').split('').map(x => false))
    })

    for (let direction = 0; direction < 4; direction++) {
        for (let i = 0; i < forest.length; i++) {
            maxValue = -1
            for (let j = 0; j < forest.length; j++) {
                let x, y
                if (direction == 0) {
                    x = i
                    y = j
                }
                else if (direction == 1) {
                    x = forest.length - 1 - j
                    y = i
                }
                else if (direction == 2) {
                    x = i
                    y = forest.length - 1 - j
                }
                else {
                    x = j
                    y = i
                }
                area = forest[y][x]
                if (area > maxValue) {
                    maxValue = area
                    visibleTrees[y][x] = true
                }
            }
        }
    }

    sum = 0
    visibleTrees.flat().forEach(tree => {
        if (tree) sum++
    })

    console.log(sum);
})

fs.readFile("../../AoC-input/2022/8/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let forest = [] // y, x

    lines.forEach(line => {
        forest.push(line.replace('\r', '').split('').map(Number))
    })

    let bestValue = 0
    for (let startX = 1; startX < forest.length - 1; startX++) {
        for (let startY = 1; startY < forest.length - 1; startY++) {
            let values = [0, 0, 0, 0]
            let yourValue = forest[startY][startX]
            for (let direction = 0; direction < 4; direction++) {
                let x = 0
                let y = 0
                while (startX + x >= 0 && startX + x < forest.length && startY + y >= 0 && startY + y < forest.length) {
                    if (x !== 0 || y !== 0) {
                        values[direction] = values[direction] + 1
                        if (forest[startY + y][startX + x] >= yourValue) {
                            break
                        }
                    }

                    if (direction == 0) {
                        y++
                    }
                    else if (direction == 1) {
                        x++
                    }
                    else if (direction == 2) {
                        y--
                    }
                    else {
                        x--
                    }
                }
            }
            let score = values[0] * values[1] * values[2] * values[3]
            if (score > bestValue) {
                bestValue = score
            }
        }
    }

    console.log(bestValue);
})
