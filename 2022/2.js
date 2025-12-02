const fs = require("fs")

fs.readFile("../../AoC-input/2022/2/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let lines = data.split('\n')
    let sum = 0

    let picks = {
        'A': 1,
        'B': 2,
        'C': 3,
        'X': 1,
        'Y': 2,
        'Z': 3
    }

    lines.forEach(line => {
        let enemyPick = picks[line.split(' ')[0]]
        let yourPick = picks[line.split(' ')[1].replace('\r', '')]

        if (yourPick === enemyPick) {
            sum += 3
        }
        else if (yourPick - 1 === enemyPick % 3) {
            sum += 6
        }

        sum += yourPick
    })

    console.log(sum);
});

fs.readFile("../../AoC-input/2022/2/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()

    let lines = data.split('\n')
    let sum = 0

    let picks = {
        'A': 1,
        'B': 2,
        'C': 3,
        'X': 1,
        'Y': 2,
        'Z': 3
    }

    lines.forEach(line => {
        let enemyPick = picks[line.split(' ')[0]]
        let yourPick = picks[line.split(' ')[1].replace('\r', '')]

        if (yourPick === 1) {
            yourPick = enemyPick === 1 ? 3 : enemyPick - 1
        }
        else if (yourPick === 2) {
            yourPick = enemyPick
        }
        else {
            yourPick = enemyPick === 3 ? 1 : enemyPick + 1
        }

        if (yourPick === enemyPick) {
            sum += 3
        }
        else if (yourPick - 1 === enemyPick % 3) {
            sum += 6
        }

        sum += yourPick
    })

    console.log(sum);
});
