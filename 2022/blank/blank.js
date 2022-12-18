/*
Description:


*/

const fs = require("fs")

fs.readFile("input.txt", (err, input) => {
    if (err) throw err
    let data = input.toString()
    let lines = data.replaceAll('\r', '').split('\n')

    lines.forEach(line => {
        
    })

    console.log()
    // 
})

/*

*/