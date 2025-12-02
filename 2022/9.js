const fs = require("fs")

fs.readFile("../../AoC-input/2022/9/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let minimap = [
        [0, 0, 0],
        [0, 3, 0],
        [0, 0, 0]
    ]

    let xT = 0, yT = 0
    const visitedCoordinates = new Set()
    visitedCoordinates.add(JSON.stringify({x: xT, y: yT}))

    lines.forEach(line => {
        let lineArr = line.replace('\r', '').split(' ').map((a, index) => index === 1 ? parseInt(a) : a)
        for (let i = 0; i < lineArr[1]; i++) {
            let xH = 0, yH = 0

            if      (lineArr[0] === 'U') yH++
            else if (lineArr[0] === 'R') xH++
            else if (lineArr[0] === 'D') yH--
            else if (lineArr[0] === 'L') xH--

            if (minimap.flat().includes(3)) {
                minimap[1 - yH][1 - xH] = 1
                minimap[1][1] = 2
            }
            else {
                let isMoving = false
                for (let j = 0; j < 3; j++) {
                    if (minimap[yH === 0 ? j : 1 - yH][xH === 0 ? j : 1 - xH] === 1) {
                        xT += xH + (xH === 0 ? 1 : 1 - xH) - (xH === 0 ? j : 1 - xH)
                        yT += yH + (yH === 0 ? 1 : 1 - yH) - (yH === 0 ? j : 1 - yH)
                        visitedCoordinates.add(JSON.stringify({x: xT, y: yT}))
                        minimap[yH === 0 ? j : 1 - yH][xH === 0 ? j : 1 - xH] = 0
                        minimap[yH === 0 ? 1 : 1 - yH][xH === 0 ? 1 : 1 - xH] = 1
                        isMoving = true
                        break
                    }
                }
                if (!isMoving) {
                    if (minimap[yH === 0 ? 1 : 1 + yH][xH === 0 ? 1 : 1 + xH] === 1) {
                        minimap[yH === 0 ? 1 : 1 + yH][xH === 0 ? 1 : 1 + xH] = 0
                        minimap[1][1] = 3
                    }
                    else {
                        let index = minimap.flat().indexOf(1)
                        minimap[Math.floor(index / 3)][index % 3] = 0
                        minimap[Math.floor(index / 3) - yH][index % 3 - xH] = 1
                    }
                }
            }
        }
    })

    console.log(visitedCoordinates.size)
})

fs.readFile("../../AoC-input/2022/9/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let knotsCoordinates = []
    for (let i = 0; i < 10; i++) {
        knotsCoordinates.push([0, 0])
    }

    const visitedCoordinates = new Set()
    visitedCoordinates.add(JSON.stringify({x: 0, y: 0}))

    lines.forEach(line => {
        let lineArr = line.replace('\r', '').split(' ').map((a, index) => index === 1 ? parseInt(a) : a)
        for (let i = 0; i < lineArr[1]; i++) {
            let prevX = 0, prevY = 0

            if      (lineArr[0] === 'U') prevY++
            else if (lineArr[0] === 'R') prevX++
            else if (lineArr[0] === 'D') prevY--
            else if (lineArr[0] === 'L') prevX--

            for (let knots = 0; knots < 10; knots++) {
                if (knots === 0) {
                    knotsCoordinates[0][0] += prevX
                    knotsCoordinates[0][1] += prevY
                }
                else {
                    if (Math.abs(knotsCoordinates[knots][0] - knotsCoordinates[knots - 1][0]) > 1 || Math.abs(knotsCoordinates[knots][1] - knotsCoordinates[knots - 1][1]) > 1) {
                        let isNextTo = false
                        for (let i = 0; i < 4; i++) {
                            let x = 0, y = 0
                            if      (i === 0) y++
                            else if (i === 1) x++
                            else if (i === 2) y--
                            else if (i === 3) x--

                            if (!(Math.abs(knotsCoordinates[knots][0] - (knotsCoordinates[knots - 1][0] + x)) > 1 || Math.abs(knotsCoordinates[knots][1] - (knotsCoordinates[knots - 1][1] + y)) > 1)) {
                                prevX = knotsCoordinates[knots - 1][0] + x - knotsCoordinates[knots][0]
                                prevY = knotsCoordinates[knots - 1][1] + y - knotsCoordinates[knots][1]
                                knotsCoordinates[knots][0] = knotsCoordinates[knots - 1][0] + x
                                knotsCoordinates[knots][1] = knotsCoordinates[knots - 1][1] + y

                                if (knots === 9) {
                                    visitedCoordinates.add(JSON.stringify({x: knotsCoordinates[knots][0], y: knotsCoordinates[knots][1]}))
                                }
                                isNextTo = true
                                break
                            }
                        }
                        if (!isNextTo) {
                            for (let i = 0; i < 4; i++) {
                                let x = 0, y = 0
                                if      (i === 0) {x++; y++}
                                else if (i === 1) {x++; y--}
                                else if (i === 2) {x--; y--}
                                else if (i === 3) {x--; y++}
    
                                if (!(Math.abs(knotsCoordinates[knots][0] - (knotsCoordinates[knots - 1][0] + x)) > 1 || Math.abs(knotsCoordinates[knots][1] - (knotsCoordinates[knots - 1][1] + y)) > 1)) {
                                    prevX = knotsCoordinates[knots - 1][0] + x - knotsCoordinates[knots][0]
                                    prevY = knotsCoordinates[knots - 1][1] + y - knotsCoordinates[knots][1]
                                    knotsCoordinates[knots][0] = knotsCoordinates[knots - 1][0] + x
                                    knotsCoordinates[knots][1] = knotsCoordinates[knots - 1][1] + y
    
                                    if (knots === 9) {
                                        visitedCoordinates.add(JSON.stringify({x: knotsCoordinates[knots][0], y: knotsCoordinates[knots][1]}))
                                    }
                                    break
                                }
                            }
                        }
                    }
                }
            } 
        }
    })

    console.log(visitedCoordinates.size)
})
