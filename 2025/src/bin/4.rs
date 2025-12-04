use std::{env, fs};

fn part1(file: &str) -> i32 {
    let mut mines = vec![vec![false; file.lines().next().unwrap().len()]; file.lines().count()];
    let mut mines_around = vec![vec![0; file.lines().next().unwrap().len()]; file.lines().count()];
    for (y, line) in file.lines().enumerate() {
        for (x, char) in line.chars().enumerate() {
            if char == '@' {
                mines[y][x] = true;
                for dy in -1isize..=1 {
                    for dx in -1isize..=1 {
                        if dy == 0 && dx == 0 {
                            continue;
                        }

                        let ny = y as isize + dy;
                        let nx = x as isize + dx;
                        if ny >= 0 && ny < mines.len() as isize && nx >= 0 && nx < mines[0].len() as isize {
                            mines_around[ny as usize][nx as usize] += 1;
                        }
                    }
                }
            }
        }
    }

    let mut count = 0;
    for y in 0..mines.len() {
        for x in 0..mines[0].len() { 
            if mines[y][x] {
                if mines_around[y][x] < 4 {
                    count += 1;
                }
            }
        }
    }
    count
}

fn part2(file: &str) -> i32 {
    let mut mines = vec![vec![false; file.lines().next().unwrap().len()]; file.lines().count()];
    for (y, line) in file.lines().enumerate() {
        for (x, char) in line.chars().enumerate() {
            if char == '@' {
                mines[y][x] = true;
            }
        }
    }

    let mut count = 0;
    loop {
        let mut mines_around = vec![vec![0; file.lines().next().unwrap().len()]; file.lines().count()];
        for (y, _) in mines.iter().enumerate() {
            for (x, _) in mines[y].iter().enumerate() {
                if mines[y][x] {
                    for dy in -1isize..=1 {
                        for dx in -1isize..=1 {
                            if dy == 0 && dx == 0 {
                                continue;
                            }

                            let ny = y as isize + dy;
                            let nx = x as isize + dx;
                            if ny >= 0 && ny < mines.len() as isize && nx >= 0 && nx < mines[0].len() as isize {
                                mines_around[ny as usize][nx as usize] += 1;
                            }
                        }
                    }
                }
            }
        }

        let mut interation_count = 0;
        for y in 0..mines.len() {
            for x in 0..mines[0].len() { 
                if mines[y][x] {
                    if mines_around[y][x] < 4 {
                        interation_count += 1;
                        mines[y][x] = false;
                    }
                }
            }
        }

        if interation_count == 0 {
            break;
        }
        count += interation_count;
    }
    count
}

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/4/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    println!("Part 2: {}", part2(&file));
}
