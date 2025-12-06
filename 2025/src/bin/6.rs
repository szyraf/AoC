use std::{env, fs};

fn part1(file: &str) -> i64 {
    let mut matrix: Vec<Vec<i64>> = vec![];
    let mut is_addition = vec![];
    for line in file.lines() {
        if line.contains("+") {
            for c in line.split_whitespace() {
                is_addition.push(c == "+");
            }
        }
        else {
            let mut row = vec![];
            for number in line.split_whitespace() {
                row.push(number.parse::<i64>().unwrap());
            }
            matrix.push(row);
        }
    }

    let mut sum = 0;
    for i in 0..is_addition.len() {
        if is_addition[i] {
            for j in 0..matrix.len() {
                sum += matrix[j][i];
            }
        }
        else {
            let mut local_sum = 1;
            for j in 0..matrix.len() {
                local_sum *= matrix[j][i];
            }
            sum += local_sum;
        }
    }
    sum
}

// fn part2(file: &str) -> i64 {

// }

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/6/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    // println!("Part 2: {}", part2(&file));
}
