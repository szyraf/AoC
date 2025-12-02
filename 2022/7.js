const fs = require("fs")

let sum = 0
let drive = {
    folders: {},
    value: 0
}

function recurention(dir) {
    if (dir['value'] <= 100000) {
        sum += dir['value']
    }
    if (dir['folders'] === {}) return
    for (let key in dir['folders']) {
        recurention(dir['folders'][key])
    }
}

fs.readFile("../../AoC-input/2022/7/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let currentLocation = []

    lines.forEach(rawLine => {
        line = rawLine.replace('\r', '')
        if (line[0] === '$') {
            if (line === '$ cd /' || line === '$ ls') {
                // nothing
            }
            else if (line === '$ cd ..') {
                currentLocation.pop()
            }
            else {
                // cd folder
                let folderName = line.substring(5)
                currentLocation.push(folderName)
            }
        }
        else if (line.substring(0, 3) === 'dir') {
            objname = line.substring(4)
            command = "drive"
            currentLocation.forEach(folder => {
                command += "['folders']['" + folder + "']"
            })
            command += "['folders']['" + objname + "'] = {folders: {}, value: 0}"
            eval(command)
        }
        else {
            for (let i = -1; i < currentLocation.length; i++) {
                command = "drive"
                for (let j = 0; j <= i; j++) {
                    command += "['folders']['" + currentLocation[j] + "']"
                }
                command += "['value'] += parseInt(line.split(' ')[0])"
                eval(command)
            }
        }
    })

    recurention(drive)

    console.log(sum)
})

let minDeletion = 70000000
let freeSpace = 70000000
let drive2 = {
    folders: {},
    value: 0
}

function recurention2(dir) {
    if (dir['value'] + freeSpace >= 30000000) {
        if (dir['value'] < minDeletion) {
            minDeletion = dir['value']
        }
    }
    if (dir['folders'] === {}) return
    for (let key in dir['folders']) {
        recurention2(dir['folders'][key])
    }
}

fs.readFile("../../AoC-input/2022/7/input.txt", (err, input) => {
    if (err) throw err;
    let data = input.toString()
    let lines = data.split('\n')

    let currentLocation = []

    lines.forEach(rawLine => {
        line = rawLine.replace('\r', '')
        if (line[0] === '$') {
            if (line === '$ cd /' || line === '$ ls') {
                // nothing
            }
            else if (line === '$ cd ..') {
                currentLocation.pop()
            }
            else {
                // cd folder
                let folderName = line.substring(5)
                currentLocation.push(folderName)
            }
        }
        else if (line.substring(0, 3) === 'dir') {
            objname = line.substring(4)
            command = "drive2"
            currentLocation.forEach(folder => {
                command += "['folders']['" + folder + "']"
            })
            command += "['folders']['" + objname + "'] = {folders: {}, value: 0}"
            eval(command)
        }
        else {
            for (let i = -1; i < currentLocation.length; i++) {
                command = "drive2"
                for (let j = 0; j <= i; j++) {
                    command += "['folders']['" + currentLocation[j] + "']"
                }
                command += "['value'] += parseInt(line.split(' ')[0])"
                eval(command)
            }
        }
    })

    freeSpace -= drive2['value']
    if (freeSpace >= 30000000) {
        minDeletion = 0
    }
    recurention2(drive2)

    console.log(minDeletion)
})
