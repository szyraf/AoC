use std::{env, fs};
use good_lp::{constraint, variable, Solution, SolverModel, ProblemVariables, Expression};

fn part1(file: &str) -> i32 {
    let mut sum = 0;
    for line in file.lines() {
        let chunks = line.split(" ").collect::<Vec<&str>>();
        let mut solution = vec![false; chunks[0].len() - 2];
        for (i, c) in chunks[0][1..chunks[0].len()-1].chars().enumerate() {
            if c == '#' {
                solution[i] = true;
            }
        }

        let mut buttons = vec![];
        for chunk in chunks[1..chunks.len() - 1].iter() {
            let parsed_chunk = chunk[1..chunk.len()-1].split(",")
                .map(|x| x.parse::<i32>().unwrap())
                .collect::<Vec<i32>>();
            buttons.push(parsed_chunk);
        }

        let mut stack: Vec<(Vec<bool>, i32, usize)> = vec![]; // (indicator_lights, iteration_count, index)
        let mut length = 0;
        for i in 0..buttons.len() {
            let mut indicator_lights = vec![false; solution.len()];
            for j in 0..buttons[i].len() {
                indicator_lights[buttons[i][j] as usize] = true;
            }

            if indicator_lights == solution {
                length = 1;
                break;
            }

            for j in i + 1..buttons.len() {
                stack.push((indicator_lights.clone(), 1, j));
            }
        }

        while length == 0 {
            let (mut indicator_lights, iteration_count, index) = stack.remove(0);

            for i in 0..buttons[index].len() {
                indicator_lights[buttons[index][i] as usize] = !indicator_lights[buttons[index][i] as usize];
            }

            if indicator_lights == solution {
                length = iteration_count + 1;
                break;
            }

            for j in index + 1..buttons.len() {
                stack.push((indicator_lights.clone(), iteration_count + 1, j));
            }
        }

        sum += length;
    }
    sum
}

// part 2

// Brute force approach – extremely slow, takes many hours to complete.
/*
fn check(ligths: Vec<i32>, buttons: Vec<Vec<i32>>, iteration_count: i32, last_min_number_index: i32, next_index: i32, last_possible_buttons: Vec<Vec<i32>>) -> i32 {
    let mut min_number_index = i32::MAX;
    let mut possible_buttons = vec![];

    if last_possible_buttons.len() > 0 {
        min_number_index = last_min_number_index;
        possible_buttons = last_possible_buttons.clone();
    }
    else {
        for i in 0..ligths.len() {
            if ligths[i] > 0 && ligths[i] < min_number_index {
                min_number_index = i as i32;
            }
        }

        if min_number_index == i32::MAX {   
            return iteration_count;
        }
        
        for i in 0..buttons.len() {
            if buttons[i].contains(&min_number_index) {
                possible_buttons.push(buttons[i].clone());
            }
        }

        let mut local_sum = 0;
        for i in 0..possible_buttons.len() {
            let mut local_min = i32::MAX;
            for j in 0..possible_buttons[i].len() {
                if ligths[possible_buttons[i][j] as usize] < local_min {
                    local_min = ligths[possible_buttons[i][j] as usize];
                }
            }
            local_sum += local_min;
        }
        if local_sum < ligths[min_number_index as usize] {
            return i32::MAX;
        }
    }

    let mut min_count = i32::MAX;
    
    let mut index: usize = 0;
    if last_min_number_index == min_number_index { // if the last min_number_index is the same as the current min_number_index, then we can start from the same button or next to prevent repeating the same button
        index = next_index as usize;
    }

    let mut i = index;
    while i < possible_buttons.len() {
        let mut ligths_copy = ligths.clone();
        let mut buttons_copy = buttons.clone();
        let mut possible_buttons_copy = possible_buttons.clone();

        let mut remove_index = vec![];
        let mut is_valid = true;
        
        for j in 0..possible_buttons_copy[i].len() {
            if ligths_copy[possible_buttons_copy[i][j] as usize] == 0 {
                possible_buttons_copy.remove(i);
                is_valid = false;
                break;
            }
            ligths_copy[possible_buttons_copy[i][j] as usize] -= 1;
            if ligths_copy[possible_buttons_copy[i][j] as usize] == 0 {
                remove_index.push(possible_buttons_copy[i][j]);
            }
        }

        if !is_valid {
            continue;
        }

        let mut new_last_possible_buttons = vec![];
        if ligths_copy[min_number_index as usize] != 0 {
            new_last_possible_buttons = possible_buttons_copy.clone();
        }

        for j in remove_index {
            let mut k = 0;
            while k < buttons_copy.len() {
                if buttons_copy[k].contains(&j) {
                    buttons_copy.remove(k);
                } else {
                    k += 1;
                }
            }

            let mut k = 0;
            while k < new_last_possible_buttons.len() {
                if new_last_possible_buttons[k].contains(&j) {
                    new_last_possible_buttons.remove(k);
                    if k < i {
                        i -= 1;
                    }
                } else {
                    k += 1;
                }
            }
        }

        let count = check(ligths_copy, buttons_copy, iteration_count + 1, min_number_index, i as i32, new_last_possible_buttons);
        if count < min_count {
            min_count = count;
            return min_count;
        }
        
        i += 1;
    }

    if index + 1 < possible_buttons.len() {
        let count = check(ligths.clone(), buttons.clone(), iteration_count, min_number_index, index as i32 + 1, possible_buttons.clone());
        if count < min_count {
            min_count = count;
            return min_count;
        }
    }

    min_count
}
*/

fn linear_solution(lights: Vec<i32>, buttons: Vec<Vec<i32>>) -> i32 {
    let mut vars = ProblemVariables::new();
    let button_vars: Vec<_> = (0..buttons.len())
        .map(|_| vars.add(variable().integer().min(0)))
        .collect();
    
    let objective: Expression = button_vars.iter().map(|&v| v).sum();
    
    let mut problem = vars.minimise(objective).using(good_lp::default_solver);
    
    for (light_idx, &target) in lights.iter().enumerate() {
        let affecting: Expression = buttons.iter()
            .enumerate()
            .filter(|(_, btn)| btn.contains(&(light_idx as i32)))
            .map(|(i, _)| button_vars[i])
            .sum();
        
        problem = problem.with(constraint!(affecting == target as f64));
    }
    
    match problem.solve() {
        Ok(solution) => {
            let result: f64 = button_vars.iter()
                .map(|&var| solution.value(var).round())
                .sum();
            result as i32
        }
        Err(_) => i32::MAX,
    }
}

fn part2(file: &str) -> i64 {
    let mut sum: i64 = 0;
    
    for line in file.lines() {
        let chunks = line.split(" ").collect::<Vec<&str>>();

        let mut buttons = vec![];
        for chunk in chunks[1..chunks.len() - 1].iter() {
            let parsed_chunk = chunk[1..chunk.len() - 1].split(",")
                .map(|x| x.parse::<i32>().unwrap())
                .collect::<Vec<i32>>();
            buttons.push(parsed_chunk);
        }

        let mut seen = Vec::new();
        let mut unique_buttons = Vec::new();
        for mut button in buttons {
            button.sort();
            if !seen.contains(&button) {
                seen.push(button.clone());
                unique_buttons.push(button);
            }
        }
        buttons = unique_buttons;

        buttons.sort_by_key(|x| x.len());
        buttons.reverse();

        let mut solution: Vec<i32> = vec![];
        for i in chunks[chunks.len() - 1][1..chunks[chunks.len() - 1].len()-1].split(",") {
            solution.push(i.parse::<i32>().unwrap());
        }

        // let count = check(solution, buttons, 0, -1, -1, vec![]);
        // println!("count: {:?}", count);

        let count = linear_solution(solution, buttons);
        sum += count as i64;
    }
    sum
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let filename = args.get(1)
        .map(|s| s.clone())
        .unwrap_or_else(|| "../../AoC-input/2025/10/input.txt".to_string());
    
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    println!("Part 2: {}", part2(&file));
}
