const fs = require("fs")

fs.readFile("../../AoC-input/2022/4/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let lines = data.split('\n')
    let sum = 0

    lines.forEach(line => {
        let numbers = [line.split(',')[0].split('-')[0], line.split(',')[0].split('-')[1], line.split(',')[1].split('-')[0], line.split(',')[1].split('-')[1]].map(num => parseInt(num))
        if (numbers[0] <= numbers[2] && numbers[1] >= numbers[3] || numbers[0] >= numbers[2] && numbers[1] <= numbers[3]) {
            sum++
        }
    })

    console.log(sum);
})

fs.readFile("../../AoC-input/2022/4/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let lines = data.split('\n')
    let sum = 0

    lines.forEach(line => {
        let numbers = [line.split(',')[0].split('-')[0], line.split(',')[0].split('-')[1], line.split(',')[1].split('-')[0], line.split(',')[1].split('-')[1]].map(num => parseInt(num))
        if (numbers[0] <= numbers[2] && numbers[1] >= numbers[2] || numbers[0] <= numbers[3] && numbers[1] >= numbers[3] || (numbers[0] <= numbers[2] && numbers[1] >= numbers[3] || numbers[0] >= numbers[2] && numbers[1] <= numbers[3])) {
            sum++
        }
    })

    console.log(sum);
})
