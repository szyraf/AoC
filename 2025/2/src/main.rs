use std::{env, fs};
use std::collections::HashMap;

fn part1(file: &str) -> i64 {
    let mut sum = 0;
    
    for line in file.lines() {
        let ranges = line.split(',');
        for range in ranges {
            let mut numbers = range.split('-').map(|x| x.parse::<i64>().unwrap());
            let start = numbers.next().unwrap();
            let end = numbers.next().unwrap();

            for number in start..=end {
                let number_string = number.to_string();
                
                if number_string.len() % 2 == 1 {
                    continue;
                }

                let first_half = number_string.chars().take(number_string.len() / 2).collect::<String>();
                let second_half = number_string.chars().skip(number_string.len() / 2).collect::<String>();

                if first_half == second_half {
                    sum += number;
                }
            }
        }
    }
    sum
}


fn part2(file: &str) -> i64 {
    let mut sum = 0;
    
    for line in file.lines() {
        let ranges = line.split(',');
        for range in ranges {
            let mut numbers = range.split('-').map(|x| x.parse::<i64>().unwrap());
            let start = numbers.next().unwrap();
            let end = numbers.next().unwrap();

            let mut dict: HashMap<usize, Vec<usize>> = HashMap::new();

            for number in start..=end {
                let number_string = number.to_string();

                if number_string.len() < 2 {
                    continue;
                }

                let mut divisors = vec![1];
                if dict.contains_key(&number_string.len()) {
                    divisors = dict.get(&number_string.len()).unwrap().clone();
                }
                else {
                    for i in 2..=number_string.len() / 2 {
                        if number_string.len() % i == 0 {
                            divisors.push(i);
                        }
                    }
                    dict.insert(number_string.len(), divisors.clone());
                }

                for &divisor in divisors.iter() {
                    let mut is_valid = true;
                    for i in 0..divisor {
                        for j in 1..(number_string.len() / divisor) {
                            if number_string.chars().nth(i).unwrap() != number_string.chars().nth(j * divisor + i).unwrap() {
                                is_valid = false;
                                break;
                            }
                        }
                    }
                    if is_valid {
                        sum += number;
                        break;
                    }
                }
            }
        }
    }
    sum
}
fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "testinput.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    println!("Part 2: {}", part2(&file));
}
