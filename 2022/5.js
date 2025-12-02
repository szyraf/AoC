const fs = require("fs")

fs.readFile("../../AoC-input/2022/5/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let lines = data.split('\n')
    let topLetters = ''

    let phase = 0
    let lettersLines = []
    let letters = []

    lines.forEach(line => {
        if (phase === 0) {
            if (line.replace(' ', '')[0] === '1') {
                let arrLength = parseInt(line.split(" ").join('').replace('\r', '')[line.split(" ").join('').replace('\r', '').length - 1])
                for (let i = 0; i < arrLength; i++) {
                    letters[i] = []
                }
                phase = 1
            }
            else {
                lettersLines.push(line)
            }
        }
        else if (phase === 1) {
            for (let i = 0; i < lettersLines.length; i++) {
                while (lettersLines[i].indexOf('[') != -1) {
                    let index = lettersLines[i].indexOf('[')
                    letters[index / 4].push(lettersLines[i][index + 1])
                    let newarr = lettersLines[i].split('')
                    newarr[index] = ']'
                    lettersLines[i] = newarr.join('')
                }
            }
            for (let i = 0; i < letters.length; i++) {
                letters[i] = letters[i].reverse()
            }
            phase = 2
        }
        else if (phase === 2) {
            let moveFromTo = line.split(' ')
            moveFromTo = [parseInt(moveFromTo[1]), parseInt(moveFromTo[3]), parseInt(moveFromTo[5].replace('\r', ''))]
            for (let i = 0; i < moveFromTo[0]; i++) {
                letters[moveFromTo[2] - 1].push(letters[moveFromTo[1] - 1].pop())
            }
        }
    })

    for (let i = 0; i < letters.length; i++) {
        topLetters += letters[i][letters[i].length - 1]
    }

    console.log(topLetters)
})

fs.readFile("../../AoC-input/2022/5/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let lines = data.split('\n')
    let topLetters = ''

    let phase = 0
    let lettersLines = []
    let letters = []

    lines.forEach(line => {
        if (phase === 0) {
            if (line.replace(' ', '')[0] === '1') {
                let arrLength = parseInt(line.split(" ").join('').replace('\r', '')[line.split(" ").join('').replace('\r', '').length - 1])
                for (let i = 0; i < arrLength; i++) {
                    letters[i] = []
                }
                phase = 1
            }
            else {
                lettersLines.push(line)
            }
        }
        else if (phase === 1) {
            for (let i = 0; i < lettersLines.length; i++) {
                while (lettersLines[i].indexOf('[') != -1) {
                    let index = lettersLines[i].indexOf('[')
                    letters[index / 4].push(lettersLines[i][index + 1])
                    let newarr = lettersLines[i].split('')
                    newarr[index] = ']'
                    lettersLines[i] = newarr.join('')
                }
            }
            for (let i = 0; i < letters.length; i++) {
                letters[i] = letters[i].reverse()
            }
            phase = 2
        }
        else if (phase === 2) {
            let moveFromTo = line.split(' ')
            moveFromTo = [parseInt(moveFromTo[1]), parseInt(moveFromTo[3]), parseInt(moveFromTo[5].replace('\r', ''))]
            let popArr = []
            for (let i = 0; i < moveFromTo[0]; i++) {
                popArr.push(letters[moveFromTo[1] - 1].pop())
            }
            for (let i = 0; i < moveFromTo[0]; i++) {
                letters[moveFromTo[2] - 1].push(popArr.pop())
            }
        }
    })

    for (let i = 0; i < letters.length; i++) {
        topLetters += letters[i][letters[i].length - 1]
    }

    console.log(topLetters)
})
