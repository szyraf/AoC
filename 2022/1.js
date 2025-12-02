const fs = require("fs")

fs.readFile("../../AoC-input/2022/1/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let elves = data.split('\n\n')
    let maxSum = 0
    elves.forEach(element => {
        let localSum = 0
        element.split('\n').forEach(number => {
            localSum += parseInt(number)
        })
        if (localSum > maxSum) maxSum = localSum
    })

    console.log(maxSum);
});

fs.readFile("../../AoC-input/2022/1/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let elves = data.split('\n\n')
    let topThreeSums = [0, 0, 0]
    elves.forEach(element => {
        let localSum = 0
        element.split('\n').forEach(number => {
            localSum += parseInt(number)
        })

        for (let i = 0; i < 3; i++) {
            if (localSum > topThreeSums[i]) {
                for (let j = i; j < 2; j++) {
                    topThreeSums[j + 1] = topThreeSums[j]
                }
                topThreeSums[i] = localSum
                break
            }
        }
    })

    console.log(topThreeSums[0] + topThreeSums[1] + topThreeSums[2])
});