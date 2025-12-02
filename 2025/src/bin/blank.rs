use std::{env, fs};

fn part1(file: &str) -> i32 {
    for line in file.lines() {

    }
    1
}

// fn part2(file: &str) -> i32 {
//     for line in file.lines() {

//     }
//     1
// }

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/blank/testinput.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    // println!("Part 2: {}", part2(&file));
}
