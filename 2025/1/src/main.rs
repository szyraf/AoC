use std::{env, fs};

fn part1(file: &str) -> i32 {
    let mut number = 50;
    let mut count = 0;

    for line in file.lines() {
        let firstletter = line.chars().next().unwrap();
        let distance = line.chars().skip(1).collect::<String>().parse::<i32>().unwrap();

        if firstletter == 'L' {
            number = (number - distance) % 100;
        } else {
            number = (number + distance) % 100;
        }

        if number < 0 {
            number += 100;
        }

        if number == 0 {
            count += 1;
        }
    }
    count
}

fn part2(file: &str) -> i32 {
    let mut number = 50;
    let mut count = 0;

    for line in file.lines() {
        let firstletter = line.chars().next().unwrap();
        let distance = line.chars().skip(1).collect::<String>().parse::<i32>().unwrap();
        let mut was_zero = false;

        if number == 0 {
            was_zero = true;
        }

        if firstletter == 'L' {
            number -= distance;
        } else {
            number += distance;
        }

        if number == 0 {
            count += 1;    
        }
        else {
            while number < 0 {
                number += 100;
                if was_zero {
                    was_zero = false;
                }
                else {
                    count += 1;
                }

                if number == 0 {
                    count += 1;    
                }
            }

            while number > 99 {
                number -= 100;
                count += 1;    
            }
        }
    }
    count
}

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "testinput.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    println!("Part 2: {}", part2(&file));
}
