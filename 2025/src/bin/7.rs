use std::{env, fs};

fn move_down(position: (usize, usize), splitter_matrix: &mut Vec<Vec<i32>>) {
    let new_position = (position.0 + 1, position.1);
    
    if new_position.0 >= splitter_matrix.len() {
        return;
    }

    if splitter_matrix[new_position.0][new_position.1] == 1 {
        move_down((new_position.0, new_position.1 - 1), splitter_matrix);
        move_down((new_position.0, new_position.1 + 1), splitter_matrix);
        splitter_matrix[new_position.0][new_position.1] = 2;
    }
    else if splitter_matrix[new_position.0][new_position.1] == 0 {
        move_down(new_position, splitter_matrix)    
    }
}

fn part1(file: &str) -> i32 {
    let mut start_position = (0, 0);
    let mut splitter_matrix: Vec<Vec<i32>> = vec![vec![0; file.lines().next().unwrap().len()]; file.lines().count()];
    for (i, line) in file.lines().enumerate() {
        for (j, c) in line.chars().enumerate() {
            if c == 'S' {
                start_position = (i, j);
            } else if c == '^' {
                splitter_matrix[i][j] = 1;
            }
        }
    }

    move_down(start_position, &mut splitter_matrix);

    let mut splitter_count = 0;
    for i in 0..splitter_matrix.len() {
        for j in 0..splitter_matrix[i].len() {
            if splitter_matrix[i][j] == 2 {
                splitter_count += 1;
            }
        }
    }

    splitter_count
}

fn move_down_part2(position: (usize, usize), splitter_matrix: &mut Vec<Vec<i64>>, timelines_count: i64, global_timelines_count: &mut i64) {
    let new_position = (position.0 + 1, position.1);
    
    if new_position.0 >= splitter_matrix.len() {
        *global_timelines_count += timelines_count;
        return;
    }

    if splitter_matrix[new_position.0][new_position.1] == -1 {
        move_down_part2(new_position, splitter_matrix, timelines_count, global_timelines_count);
    }
    else {
        splitter_matrix[new_position.0][new_position.1] += timelines_count;
    }
}

fn part2(file: &str) -> i64 {
    let mut start_position = (0, 0);
    let mut splitter_matrix: Vec<Vec<i64>> = vec![vec![-1; file.lines().next().unwrap().len()]; file.lines().count()];
    for (i, line) in file.lines().enumerate() {
        for (j, c) in line.chars().enumerate() {
            if c == 'S' {
                start_position = (i, j);
            } else if c == '^' {
                splitter_matrix[i][j] = 0;
            }
        }
    }

    let mut global_timelines_count: i64 = 0;
    move_down_part2(start_position, &mut splitter_matrix, 1, &mut global_timelines_count);
    for i in 0..splitter_matrix.len() {
        for j in 0..splitter_matrix[i].len() {
            if splitter_matrix[i][j] != -1 {
                let count = splitter_matrix[i][j];
                move_down_part2((i, j - 1), &mut splitter_matrix, count, &mut global_timelines_count);
                move_down_part2((i, j + 1), &mut splitter_matrix, count, &mut global_timelines_count);
            }
        }
    }

    global_timelines_count
}

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/7/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    println!("Part 2: {}", part2(&file));
}
