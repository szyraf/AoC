class Monkey {
    items = []
    operation = ''
    testDivisibleBy = 1
    ifTrue = 0
    ifFalse = 0
    inspectedItems = 0
}

const fs = require("fs")

fs.readFile("../../AoC-input/2022/11/input.txt", (err, input) => {
    if (err) throw err
    let data = input.toString()
    let lines = data.replace('\r', '').split('\n').map(line => line.trim())

    let monkeys = []
    let monkeyId = 0

    lines.forEach(line => {
        if (line === '') {
            monkeyId++
        }
        else if (line[0] === 'M') {
            let monkey = new Monkey()
            monkeys.push(monkey)
        }
        else if (line[0] === 'S') {
            monkeys[monkeyId].items = line.substring(16).replace(' ', '').split(',').map(number => parseInt(number))
        }
        else if (line[0] === 'O') {
            monkeys[monkeyId].operation = line.substring(17)
        }
        else if (line[0] === 'T') {
            monkeys[monkeyId].testDivisibleBy = parseInt(line.substring(19))
        }
        else if (line[3] === 't') {
            monkeys[monkeyId].ifTrue = parseInt(line.substring(25))
        }
        else if (line[3] === 'f') {
            monkeys[monkeyId].ifFalse = parseInt(line.substring(26))
        }
    })

    for (let i = 0; i < 20; i++) {
        monkeys.forEach(monkey => {
            while (monkey.items.length != 0) {
                let worryLevel = monkey.items.shift()
                worryLevel = Math.floor(eval(monkey.operation.replaceAll('old', worryLevel.toString())) / 3)
                if (worryLevel % monkey.testDivisibleBy === 0) {
                    monkeys[monkey.ifTrue].items.push(worryLevel)
                }
                else {
                    monkeys[monkey.ifFalse].items.push(worryLevel)
                }
                monkey.inspectedItems++
            }
        })    
    } 

    let topInspectedItems = []
    monkeys.forEach(monkey => {
        topInspectedItems.push(monkey.inspectedItems)
    })
    topInspectedItems.sort((a, b) => b - a)

    console.log(topInspectedItems[0] * topInspectedItems[1])
})

class Monkey2 {
    items = []
    operation = ''
    testDivisibleBy = 1
    ifTrue = 0
    ifFalse = 0
    inspectedItems = 0
}

fs.readFile("../../AoC-input/2022/11/input.txt", (err, input) => {
    if (err) throw err
    let data = input.toString()
    let lines = data.replace('\r', '').split('\n').map(line => line.trim())

    let monkeys = []
    let monkeyId = 0

    let superModulo = 1n

    lines.forEach(line => {
        if (line === '') {
            monkeyId++
        }
        else if (line[0] === 'M') {
            let monkey = new Monkey2()
            monkeys.push(monkey)
        }
        else if (line[0] === 'S') {
            monkeys[monkeyId].items = line.substring(16).replace(' ', '').split(',').map(number => BigInt(parseInt(number)))         
        }
        else if (line[0] === 'O') {
            monkeys[monkeyId].operation = line.substring(17)
        }
        else if (line[0] === 'T') {
            monkeys[monkeyId].testDivisibleBy = BigInt(parseInt(line.substring(19)))
            superModulo *= monkeys[monkeyId].testDivisibleBy
        }
        else if (line[3] === 't') {
            monkeys[monkeyId].ifTrue = parseInt(line.substring(25))
        }
        else if (line[3] === 'f') {
            monkeys[monkeyId].ifFalse = parseInt(line.substring(26))
        }
    })

    for (let i = 0; i < 10000; i++) {
        monkeys.forEach(monkey => {
            while (monkey.items.length != 0) {
                let worryLevel = BigInt(monkey.items.shift())
                let operations = monkey.operation.replaceAll('old', worryLevel.toString()).split(' ')
                worryLevel = eval(`${operations[0]}n ${operations[1]} ${operations[2]}n`)
                worryLevel = worryLevel % superModulo

                if (worryLevel % monkey.testDivisibleBy == 0) {
                    monkeys[monkey.ifTrue].items.push(worryLevel)
                }
                else {
                    monkeys[monkey.ifFalse].items.push(worryLevel)
                }
                monkey.inspectedItems++
            }
        })
    } 

    let topInspectedItems = []
    monkeys.forEach(monkey => {
        topInspectedItems.push(monkey.inspectedItems)
    })
    topInspectedItems.sort((a, b) => b - a)

    console.log(topInspectedItems[0] * topInspectedItems[1])
})
