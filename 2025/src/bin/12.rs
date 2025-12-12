use std::{env, fs};

// Brute force solution, works for small inputs but is too slow for large inputs.
/*
fn rotate_shape(shape: &Vec<Vec<bool>>) -> Vec<Vec<bool>> {
    let mut rotated_shape = vec![vec![false; 3]; 3];
    for j in 0..3 {
        for k in 0..3 {
            rotated_shape[j][k] = shape[2 - k][j];
        }
    }
    rotated_shape
}

fn flip_shape_horizontal(shape: &Vec<Vec<bool>>) -> Vec<Vec<bool>> {
    let mut flipped_shape = vec![vec![false; 3]; 3];
    for j in 0..3 {
        for k in 0..3 {
            flipped_shape[j][k] = shape[j][2 - k];
        }
    }
    flipped_shape
}

fn is_the_same_shape(shape1: &Vec<Vec<bool>>, shape2: &Vec<Vec<bool>>) -> bool {
    for j in 0..3 {
        for k in 0..3 {
            if shape1[j][k] != shape2[j][k] {
                return false;
            }
        }
    }
    true
}

fn add_shape_if_unique(shapes: &mut Vec<Vec<Vec<Vec<bool>>>>, i: usize, new_shape: Vec<Vec<bool>>) {
    let mut is_the_same = false;
    for j in 0..shapes[i].len() {
        if is_the_same_shape(&shapes[i][j], &new_shape) {
            is_the_same = true;
            break;
        }
    }
    if !is_the_same {
        shapes[i].push(new_shape);
    }
}

fn solve_grid(shapes: &Vec<Vec<Vec<Vec<bool>>>>, mut grid: Vec<Vec<bool>>, mut data: Vec<i32>, x: usize, y: usize, shape_variation: usize) -> bool {
    let mut which_shape = usize::MAX;
    for i in 0..data.len() {
        if data[i] != 0 {
            which_shape = i;
            break;
        }
    }

    if which_shape == usize::MAX {
        return true;
    }

    if shape_variation < shapes[which_shape].len() - 1 {
        if solve_grid(shapes, grid.clone(), data.clone(), x, y, shape_variation + 1) {
            return true;
        }
    }

    let mut next_y = y;
    let mut next_x = x + 1;
    if next_x == grid[0].len() - 1 {
        next_x = 1;
        next_y += 1;
        if next_y == grid.len() - 1 {
            for i in 0..data.len() {
                if data[i] != 0 {
                    return false;
                }
            }
            return true;
        }
    }

    if solve_grid(shapes, grid.clone(), data.clone(), next_x, next_y, shape_variation) {
        return true;
    }
    
    let mut is_placable = true;
    'outer: for i in -1isize..=1 {
        for j in -1isize..=1 {
            let shape_i = (i + 1) as usize;
            let shape_j = (j + 1) as usize;
            let grid_y = (y as isize + j) as usize;
            let grid_x = (x as isize + i) as usize;
            
            if shapes[which_shape][shape_variation][shape_j][shape_i] {
                if grid[grid_y][grid_x] {
                    is_placable = false;
                    break 'outer;
                }
            }
        }
    }

    if is_placable {
        for i in -1isize..=1 {
            for j in -1isize..=1 {
                let shape_i = (i + 1) as usize;
                let shape_j = (j + 1) as usize;
                let grid_y = (y as isize + j) as usize;
                let grid_x = (x as isize + i) as usize;
                
                if shapes[which_shape][shape_variation][shape_j][shape_i] {
                    grid[grid_y][grid_x] = true;
                }
            }
        }
        data[which_shape] -= 1;

        let mut next_y = y;
        let mut next_x = x + 1;
        if next_x == grid[0].len() - 1 {
            next_x = 1;
            next_y += 1;
            if next_y == grid.len() - 1 {
                for i in 0..data.len() {
                    if data[i] != 0 {
                        return false;
                    }
                }
                return true;
            }
        }

        if solve_grid(shapes, grid.clone(), data.clone(), next_x, next_y, 0) {
            return true;
        }
    }

    return false;    
}

fn part1(file: &str) -> i32 {
    // assumes exactly 6 shapes, each 3x3. Only works for standard AoC inputs.
    let mut shapes = vec![vec![vec![vec![false; 3]; 3]; 1]; 6];
    let mut regions: Vec<Vec<Vec<i32>>> = vec![];

    for (i, line) in file.lines().enumerate() {
        if !line.contains("x") {
            if i % 5 >= 1 && i % 5 <= 3 {
                for (j, char) in line.trim().chars().enumerate() {
                    shapes[i / 5][0][i % 5 - 1][j] = char == '#';
                }
            }  
        }
        else {
            let (dimensions, data) = line.split_once(": ").unwrap();
            let (width, height) = dimensions.split_once('x').unwrap();
            let region_data: Vec<i32> = data.split(' ').map(|s| s.parse().unwrap()).collect();
            
            regions.push(vec![
                vec![width.parse().unwrap(), height.parse().unwrap()],
                region_data
            ]);
        }
    }

    for i in 0..shapes.len() {
        let rotation90 = rotate_shape(&shapes[i][0]);
        let rotation180 = rotate_shape(&rotation90);
        let rotation270 = rotate_shape(&rotation180);
        let flipped_horizontal_shape = flip_shape_horizontal(&shapes[i][0]);
        let flipped_horizontal_rotation90 = flip_shape_horizontal(&rotation90);
        let flipped_horizontal_rotation180 = flip_shape_horizontal(&rotation180);
        let flipped_horizontal_rotation270 = flip_shape_horizontal(&rotation270);

        add_shape_if_unique(&mut shapes, i, rotation90);
        add_shape_if_unique(&mut shapes, i, rotation180);
        add_shape_if_unique(&mut shapes, i, rotation270);
        add_shape_if_unique(&mut shapes, i, flipped_horizontal_shape);
        add_shape_if_unique(&mut shapes, i, flipped_horizontal_rotation90);
        add_shape_if_unique(&mut shapes, i, flipped_horizontal_rotation180);
        add_shape_if_unique(&mut shapes, i, flipped_horizontal_rotation270);
    }

    let mut possible_grids = 0;
    for (i, region) in regions.iter().enumerate() {
        let width = region[0][0];
        let height = region[0][1];
        let data = &region[1];

        let grid = vec![vec![false; width as usize]; height as usize];

        if solve_grid(&shapes, grid.clone(), data.clone(), 1, 1, 0) {
            possible_grids += 1;
            println!("Possible grid found! Region {}: {:?}", i, region);
        }
        else {
            println!("No possible grid found! Region {}: {:?}", i, region);
        }
    }
    
    possible_grids
}
*/

fn part1(file: &str) -> i32 {
    let mut shapes = [0; 6];
    let mut possible_grids = 0;
    for (i, line) in file.lines().enumerate() {
        if !line.contains("x") {
            if i % 5 >= 1 && i % 5 <= 3 {
                for char in line.trim().chars() {
                    shapes[i / 5] += (char == '#') as i32;
                }
            }  
        }
        else {
            let (dimensions, data) = line.split_once(": ").unwrap();
            let (width, height) = dimensions.split_once('x').map(|(w, h)| (w.parse::<i32>().unwrap(), h.parse::<i32>().unwrap())).unwrap();
            let region_data: Vec<i32> = data.split(' ').map(|s| s.parse().unwrap()).collect();
            
            let mut shape_space = 0;
            for j in 0..region_data.len() {
                shape_space += shapes[j] * region_data[j];
            }

            if shape_space <= width * height {
                possible_grids += 1;
            }
        }
    }

    possible_grids
}

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/12/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
}
