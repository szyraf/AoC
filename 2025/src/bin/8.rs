use std::{env, fs};

fn parse_input(file: &str) -> Vec<(i64, i64, i64)> {
    let mut junction_boxes_coords: Vec<(i64, i64, i64)> = Vec::new();
    for line in file.lines() {
        let numbers = line.split(',').map(|x| x.parse::<i64>().unwrap()).collect::<Vec<i64>>();
        junction_boxes_coords.push((numbers[0], numbers[1], numbers[2]));
    }
    junction_boxes_coords
}

fn compute_distance_ranking(junction_boxes_coords: &[(i64, i64, i64)]) -> Vec<(usize, usize, i64)> {
    let mut distance_ranking: Vec<(usize, usize, i64)> = Vec::new();
    for i in 0..junction_boxes_coords.len() {
        for j in i + 1..junction_boxes_coords.len() {
            let dist = (junction_boxes_coords[i].0 - junction_boxes_coords[j].0).pow(2)
                     + (junction_boxes_coords[i].1 - junction_boxes_coords[j].1).pow(2)
                     + (junction_boxes_coords[i].2 - junction_boxes_coords[j].2).pow(2);
            distance_ranking.push((i, j, dist));
        }
    }
    distance_ranking.sort_by_key(|x| x.2);
    distance_ranking
}

fn merge_junction_boxes_into_circuits(circuits: &mut Vec<Vec<usize>>, i: usize, j: usize) {
    let mut first_circuit_index: Option<usize> = None;
    let mut second_circuit_index: Option<usize> = None;
    for (circuit_idx, circuit) in circuits.iter().enumerate() {
        if circuit.contains(&i) {
            first_circuit_index = Some(circuit_idx);
        }

        if circuit.contains(&j) {
            second_circuit_index = Some(circuit_idx);
        }

        if first_circuit_index.is_some() && second_circuit_index.is_some() {
            break;
        }
    }

    match (first_circuit_index, second_circuit_index) {
        (Some(first), Some(second)) if first != second => {
            let second_circuit = circuits.remove(second);
            let first = if second < first { first - 1 } else { first };
            circuits[first].extend(second_circuit);
        }
        (Some(_first), Some(_)) => {
            // Both junction boxes are already in the same circuit
        }
        (Some(first), None) => {
            circuits[first].push(j);
        }
        (None, Some(second)) => {
            circuits[second].push(i);
        }
        (None, None) => {
            circuits.push(vec![i, j]);
        }
    }
}

fn part1(file: &str, shortest_connections_count: usize) -> i64 {
    let junction_boxes_coords = parse_input(file);
    let distance_ranking = compute_distance_ranking(&junction_boxes_coords);

    let mut circuits: Vec<Vec<usize>> = Vec::new();

    for idx in 0..shortest_connections_count {
        let (i, j, _) = distance_ranking[idx];
        merge_junction_boxes_into_circuits(&mut circuits, i, j);
    }

    circuits.sort_by_key(|x| std::cmp::Reverse(x.len()));

    let top_three_circuits_multiplier = circuits[0].len() * circuits[1].len() * circuits[2].len();
    top_three_circuits_multiplier as i64
}

fn part2(file: &str) -> i64 {
    let junction_boxes_coords = parse_input(file);
    let distance_ranking = compute_distance_ranking(&junction_boxes_coords);

    let mut circuits: Vec<Vec<usize>> = Vec::new();

    for idx in 0..distance_ranking.len() {
        let (i, j, _) = distance_ranking[idx];
        merge_junction_boxes_into_circuits(&mut circuits, i, j);

        if circuits[0].len() == junction_boxes_coords.len() {
            return junction_boxes_coords[i].0 * junction_boxes_coords[j].0;
        }
    }

    0
}

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/8/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    let shortest_connections_count = 1000;
    println!("Part 1: {}", part1(&file, shortest_connections_count));
    println!("Part 2: {}", part2(&file));
}
