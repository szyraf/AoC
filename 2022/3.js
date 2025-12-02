const fs = require("fs")

fs.readFile("../../AoC-input/2022/3/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let lines = data.split('\n')
    let sum = 0

    lines.forEach(line => {
        let firstCompartment = line.slice(0, line.length / 2)
        let secondCompartment = line.slice(line.length / 2)

        let letter = '';
        for (let i = 0; i < firstCompartment.length; i++) {
            for (let j = 0; j < secondCompartment.length; j++) {
                if (firstCompartment[i] === secondCompartment[j]) {
                    letter = firstCompartment[i]
                    i = firstCompartment.length
                    break
                }
            }
        }

        let letterCharCode = letter.charCodeAt(0)

        sum += letterCharCode - (letterCharCode < 91 ? 38 : 96)
    })

    console.log(sum);
})

fs.readFile("../../AoC-input/2022/3/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let lines = data.split('\n')
    let sum = 0

    let prevLine1 = ""
    let prevLine2 = ""
    lines.forEach((line, lineIndex) => {
        if (lineIndex % 3 === 0) {
            prevLine1 = line
        }
        else if (lineIndex % 3 === 1) {
            prevLine2 = line
        }
        else {
            let letter = '';
            for (let i = 0; i < line.length; i++) {
                isInPrevLine1 = false
                for (let j = 0; j < prevLine1.length; j++) {
                    if (line[i] === prevLine1[j]) {
                        isInPrevLine1 = true
                        break
                    }
                }
                if (isInPrevLine1) {
                    for (let j = 0; j < prevLine2.length; j++) {
                        if (line[i] === prevLine2[j]) {
                            letter = line[i]
                            i = line.length
                            break
                        }
                    }
                }
            }

            let letterCharCode = letter.charCodeAt(0)
            sum += letterCharCode - (letterCharCode < 91 ? 38 : 96)
        }
    })

    console.log(sum);
})
