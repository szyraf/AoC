const fs = require("fs")

fs.readFile("../../AoC-input/2022/10/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let number = 1
    let cycles = 0
    let when = 20
    let sumOfSignals = 0

    lines.forEach(line => {
        let lineArr = line.replace('\r', '').split(' ').map((a, index) => index === 1 ? parseInt(a) : a)
        
        if (lineArr.length === 1) {
            cycles++
            if (cycles === when) {
                sumOfSignals += number * when
                if (when <= 220) when += 40
            }
        }
        else {
            cycles ++
            if (cycles === when) {
                sumOfSignals += number * when
                if (when <= 220) when += 40
            }
            cycles ++
            if (cycles === when) {
                sumOfSignals += number * when
                if (when <= 220) when += 40
            }
            number += lineArr[1]
        }
    })

    console.log(sumOfSignals)
})

fs.readFile("../../AoC-input/2022/10/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let number = 1

    let crt = ''

    lines.forEach(line => {
        let lineArr = line.replace('\r', '').split(' ').map((a, index) => index === 1 ? parseInt(a) : a)

        if (Math.abs(crt.length % 40 - number) <= 1) {
            crt += '#'
        }
        else {
            crt += '.'
        }
        if (lineArr.length === 2) {
            if (Math.abs(crt.length % 40 - number) <= 1) {
                crt += '#'
            }
            else {
                crt += '.'
            }
            number += lineArr[1]
        }
    })

    for (let i = 0; i < 6; i ++) {
        console.log(crt.substring(i * 40, (i + 1) * 40))
    }
})
