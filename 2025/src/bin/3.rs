use std::{env, fs};

fn part1(file: &str) -> u32 {
    let mut sum = 0;
    for line in file.lines() {
        let mut maxnumber_1 = 1;
        let mut maxnumber_1_index = 0;

        for (index, number) in line.chars().take(line.len() - 1).filter_map(|c| c.to_digit(10)).enumerate() {
            if number > maxnumber_1 {
                maxnumber_1 = number;
                maxnumber_1_index = index;
            }
        }

        let mut maxnumber_2 = 1;
        for number in line.chars().skip(maxnumber_1_index + 1).filter_map(|c| c.to_digit(10)) {
            if number > maxnumber_2 {
                maxnumber_2 = number;
            }
        }

        sum += maxnumber_1 * 10 + maxnumber_2;
    }
    sum
}

fn part2(file: &str) -> u64 {
    let mut sum = 0_u64;
    for line in file.lines() {
        let mut maxnumber_number_arr = vec![1_u64; 12];
        let mut maxnumber_index_arr = vec![0_u64; 12];

        for i in 0..12 {
            let mut last_index: usize = 0;
            if i > 0 {
                maxnumber_index_arr[i] = maxnumber_index_arr[i - 1] + 1;
                last_index = maxnumber_index_arr[i - 1] as usize + 1;
            }

            for (index, number) in line.chars().skip(last_index).take(line.len() - 11 + i - last_index).filter_map(|c| c.to_digit(10)).enumerate() {
                if u64::from(number) > maxnumber_number_arr[i] {
                    maxnumber_number_arr[i] = u64::from(number);
                    maxnumber_index_arr[i] = (index + last_index) as u64;
                }
            }
            sum += maxnumber_number_arr[i] * 10_u64.pow((11 - i) as u32);
        }
    }
    sum
}

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/3/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    println!("Part 2: {}", part2(&file));
}
