use std::{env, fs};

fn part1(file: &str) -> i64 {
    let mut min_x: i64 = i64::MAX;
    let mut max_x: i64 = 0;
    let mut min_y: i64 = i64::MAX;
    let mut max_y: i64 = 0;
    
    let mut coordinates = vec![];

    let mut closest_to_corners = vec![vec![i64::MAX; 3]; 4]; // [distance, x, y] x4 (right top, right bottom, left bottom, left top)
    
    for line in file.lines() {
        let parts: Vec<i64> = line.split(',').map(|s| s.parse::<i64>().unwrap()).collect();
        let (x, y) = (parts[0], parts[1]);
        coordinates.push((x, y));

        if x < min_x {
            min_x = x;
        }
        else if x > max_x {
            max_x = x;
        }

        if y < min_y {
            min_y = y;
        }
        else if y > max_y {
            max_y = y;
        }
    }

    for coordinate in coordinates {
        let (x, y) = coordinate;
        let distance_to_right_top = (max_x - x) + (max_y - y);
        if distance_to_right_top < closest_to_corners[0][0] {
            closest_to_corners[0][0] = distance_to_right_top;
            closest_to_corners[0][1] = x;
            closest_to_corners[0][2] = y;
        }

        let distance_to_right_bottom = (max_x - x) + (y - min_y);
        if distance_to_right_bottom < closest_to_corners[1][0] {
            closest_to_corners[1][0] = distance_to_right_bottom;
            closest_to_corners[1][1] = x;
            closest_to_corners[1][2] = y;
        }

        let distance_to_left_bottom = (x - min_x) + (y - min_y);
        if distance_to_left_bottom < closest_to_corners[2][0] {
            closest_to_corners[2][0] = distance_to_left_bottom;
            closest_to_corners[2][1] = x;
            closest_to_corners[2][2] = y;
        }
        
        let distance_to_left_top = (x - min_x) + (max_y - y);
        if distance_to_left_top < closest_to_corners[3][0] {
            closest_to_corners[3][0] = distance_to_left_top;
            closest_to_corners[3][1] = x;
            closest_to_corners[3][2] = y;
        }
    }

    let surface_area = (closest_to_corners[0][1] - closest_to_corners[2][1] + 1) * (closest_to_corners[0][2] - closest_to_corners[2][2] + 1);
    let surface_area_2 = (closest_to_corners[1][1] - closest_to_corners[3][1] + 1) * (closest_to_corners[3][2] - closest_to_corners[1][2] + 1);

    if surface_area > surface_area_2 {
        return surface_area
    }
    surface_area_2
}

// fn part2(file: &str) -> i32 {
//     for line in file.lines() {

//     }
//     1
// }

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/9/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    // println!("Part 2: {}", part2(&file));
}
