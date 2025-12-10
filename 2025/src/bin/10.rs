use std::{env, fs};

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

// fn part2(file: &str) -> i32 {
//     let mut sum = 0;
//     for line in file.lines() {
//         let chunks = line.split(" ").collect::<Vec<&str>>();

//         let mut buttons = vec![];
//         for chunk in chunks[1..chunks.len() - 1].iter() {
//             let parsed_chunk = chunk[1..chunk.len()-1].split(",")
//                 .map(|x| x.parse::<i32>().unwrap())
//                 .collect::<Vec<i32>>();
//             buttons.push(parsed_chunk);
//         }

//         let mut solution = vec![];
//         solution = chunks[chunks.len() - 1][1..chunks[chunks.len() - 1].len()-1].split(",")
//             .map(|x| x.parse::<i32>().unwrap())
//             .collect::<Vec<i32>>();

//         println!("buttons: {:?}", buttons);
//         println!("solution: {:?}", solution);

//         sum += 1;
//         break;
//     }
//     sum
// }


fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/10/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    // println!("Part 2: {}", part2(&file));
}
