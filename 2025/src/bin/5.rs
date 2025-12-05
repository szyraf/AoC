use std::{env, fs};

fn part1(file: &str) -> i32 {
    let mut ranges_mode = true;

    let mut ranges_start = Vec::new();
    let mut ranges_end = Vec::new();

    let mut count = 0;

    for line in file.lines() {
        if line.is_empty() {
            ranges_mode = false;
            continue;
        }

        if ranges_mode {
            let range = line.split('-').map(|x| x.parse::<i64>().unwrap()).collect::<Vec<i64>>();
            ranges_start.push(range[0]);
            ranges_end.push(range[1]);
        }
        else {
            let number = line.parse::<i64>().unwrap();
            for i in 0..ranges_start.len() {
                if number >= ranges_start[i] && number <= ranges_end[i] {
                    count += 1;
                    break;
                }
            }
        }
    }
    count
}

fn part2(file: &str) -> i64 {
    let mut ranges_start = Vec::new();
    let mut ranges_end = Vec::new();

    let mut count = 0;

    let mut range_start_end = Vec::new();

    for line in file.lines() {
        if line.is_empty() {
            break;
        }

        let range = line.split('-').map(|x| x.parse::<i64>().unwrap()).collect::<Vec<i64>>();
        range_start_end.push([range[0], range[1]]);
    }

    range_start_end.sort_by_key(|x| x[0]);

    for range in range_start_end {
        let range_start = range[0];
        let range_end = range[1];

        let mut range_0_index: Option<usize> = None;
        for i in 0..ranges_start.len() {
            if range_start >= ranges_start[i] && range_start <= ranges_end[i] {
                range_0_index = Some(i);
                break;
            }
        }

        let mut range_1_index: Option<usize> = None;
        for i in 0..ranges_start.len() {
            if range_end >= ranges_start[i] && range_end <= ranges_end[i] {
                range_1_index = Some(i);
                break;
            }
        }

        match (range_0_index, range_1_index) {
            (None, None) => {
                ranges_start.push(range_start);
                ranges_end.push(range_end);
            }
            (None, Some(idx1)) => {
                ranges_start[idx1] = range_start;
            }
            (Some(idx0), None) => {
                ranges_end[idx0] = range_end;
            }
            (Some(idx0), Some(idx1)) => {
                if idx0 == idx1 {
                    continue;
                }

                if ranges_start[idx0] > ranges_start[idx1] {
                    ranges_start[idx0] = ranges_start[idx1];
                }
                else {
                    ranges_start[idx1] = ranges_start[idx0];
                }

                if ranges_end[idx0] < ranges_end[idx1] {
                    ranges_end[idx0] = ranges_end[idx1];
                }
                else {
                    ranges_end[idx1] = ranges_end[idx0];
                }

                ranges_start.remove(idx1);
                ranges_end.remove(idx1);
            }
        }
    }

    for i in 0..ranges_start.len() {
        count += ranges_end[i] - ranges_start[i] + 1;
    }

    count
}

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/5/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    println!("Part 2: {}", part2(&file));
}
