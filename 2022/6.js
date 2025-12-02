const fs = require("fs")

fs.readFile("../../AoC-input/2022/6/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let index = 3
    lines.forEach(line => {
        let allDifferentCharacters = false
        while (allDifferentCharacters === false) {
            let isSomeDifferentCharacters = false
            let letters = []
            for (let i = 3; i > 0; i--) {
                letters.push(line[index - i])
                for (let j = 0; j < letters.length; j++) {
                    if (line[index - i + 1] === line[index - 3 + j]) {
                        isSomeDifferentCharacters = true
                        i = 0
                        break
                    }
                }
            }
            if (isSomeDifferentCharacters) {
                index++
            }
            else {
                allDifferentCharacters = true
            }
        }
    })

    console.log(index + 1)
})

fs.readFile("../../AoC-input/2022/6/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let index = 13
    lines.forEach(line => {
        let allDifferentCharacters = false
        while (allDifferentCharacters === false) {
            let isSomeDifferentCharacters = false
            let letters = []
            for (let i = 13; i > 0; i--) {
                letters.push(line[index - i])
                for (let j = 0; j < letters.length; j++) {
                    if (line[index - i + 1] === line[index - 13 + j]) {
                        isSomeDifferentCharacters = true
                        i = 0
                        break
                    }
                }
            }
            if (isSomeDifferentCharacters) {
                index++
            }
            else {
                allDifferentCharacters = true
            }
        }
    })

    console.log(index + 1)
})
