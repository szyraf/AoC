use std::{env, fs};
use std::collections::HashMap;

fn parse(file: &str) -> HashMap<&str, Vec<&str>> {
    let mut map = HashMap::new();
    
    for line in file.lines() {
        let parts = line.split(": ").collect::<Vec<&str>>();
        let key = parts[0];
        let values = parts[1].split(" ").collect::<Vec<&str>>();
        map.insert(key, values);
    }
    
    map
}

fn find_accessible_keys(map: &HashMap<&str, Vec<&str>>, start: &str) -> Vec<String> {
    let mut keys_accessible = vec![start.to_string()];
    
    let mut something_changed = true;
    while something_changed {
        something_changed = false;
        let mut new_keys = Vec::new();
        for key in keys_accessible.iter() {
            if map.contains_key(key.as_str()) {
                for value in &map[key.as_str()] {
                    let value_str = value.to_string();
                    if !keys_accessible.contains(&value_str) && !new_keys.contains(&value_str) && *value != "out" {
                        new_keys.push(value_str);
                        something_changed = true;                        
                    }
                }
            }
        }
        keys_accessible.extend(new_keys);
    }
    
    keys_accessible
}

fn remove_inaccessible_keys<'a>(map: &mut HashMap<&'a str, Vec<&'a str>>, accessible_keys: &[String]) {
    let keys_to_remove: Vec<_> = map.keys()
        .filter(|key| !accessible_keys.iter().any(|k| k == *key))
        .cloned()
        .collect();
    
    for key in keys_to_remove {
        map.remove(key);
    }
}

fn count_paths_to_devices<'a>(map: &HashMap<&'a str, Vec<&'a str>>, start: &'a str) -> HashMap<&'a str, i32> {
    let mut paths_to_devices = HashMap::new();
    paths_to_devices.insert(start, 1);
    for key in map.keys() {
        for value in &map[key] {
            if *value != "out" {
                if !paths_to_devices.contains_key(value) {
                    paths_to_devices.insert(value, 1);
                }
                else {
                    paths_to_devices.insert(value, *paths_to_devices.get(value).unwrap() + 1);
                }
            }
        }
    }
    
    paths_to_devices
}

fn part1(file: &str) -> i32 {
    let mut map = parse(file);
    let keys_accessible_from_you = find_accessible_keys(&map, "you");
    remove_inaccessible_keys(&mut map, &keys_accessible_from_you);
    let paths_to_devices = count_paths_to_devices(&map, "you");

    let mut points: HashMap<_, _> = paths_to_devices.keys().map(|k| (*k, 0)).collect();
    points.insert("you", 1);

    let mut locks = paths_to_devices.clone();
    locks.insert("you", 0);

    let mut out_sum = 0;
    let mut something_changed = true;
    while something_changed {
        something_changed = false;
        for key in locks.keys().cloned().collect::<Vec<_>>() {
            if *locks.get(key).unwrap() == 0 {
                something_changed = true;
                for value in &map[key] {
                    if *value == "out" {
                        out_sum += *points.get(key).unwrap();
                    }
                    else {
                        locks.insert(value, *locks.get(value).unwrap() - 1);
                        points.insert(value, *points.get(value).unwrap() + *points.get(key).unwrap());
                    }
                }
                locks.insert(key, -1);
            }
        }
    }

    out_sum
}

fn part2(file: &str) -> i64 {
    let mut map = parse(file);
    let keys_accessible_from_svr = find_accessible_keys(&map, "svr");
    remove_inaccessible_keys(&mut map, &keys_accessible_from_svr);
    let paths_to_devices = count_paths_to_devices(&map, "svr");

    let mut points: HashMap<_, _> = paths_to_devices.keys().map(|k| (*k, (0i64, 0i64, 0i64, 0i64))).collect();
    points.insert("svr", (1i64, 0i64, 0i64, 0i64));

    let mut locks = paths_to_devices.clone();
    locks.insert("svr", 0);

    let mut out_sum = 0i64;
    let mut something_changed = true;
    while something_changed {
        something_changed = false;
        for key in locks.keys().cloned().collect::<Vec<_>>() {
            if *locks.get(key).unwrap() == 0 {
                something_changed = true;
                for value in &map[key] {
                    if *value == "out" {
                        out_sum += points.get(key).unwrap().3;
                    }
                    else {
                        locks.insert(value, *locks.get(value).unwrap() - 1);

                        if *value == "fft" {
                            points.insert(
                                value,
                                (
                                    points.get(value).unwrap().0 + points.get(key).unwrap().0,
                                    points.get(value).unwrap().1 + points.get(key).unwrap().0,
                                    points.get(value).unwrap().2 + points.get(key).unwrap().2,
                                    std::cmp::min(points.get(value).unwrap().1 + points.get(key).unwrap().0, points.get(value).unwrap().2 + points.get(key).unwrap().2)
                                )
                            );
                        }
                        else if *value == "dac" {
                            points.insert(
                                value,
                                (
                                    points.get(value).unwrap().0 + points.get(key).unwrap().0,
                                    points.get(value).unwrap().1 + points.get(key).unwrap().1,
                                    points.get(value).unwrap().2 + points.get(key).unwrap().0,
                                    std::cmp::min(points.get(value).unwrap().1 + points.get(key).unwrap().1, points.get(value).unwrap().2 + points.get(key).unwrap().0)
                                )
                            );
                        }
                        else {
                            points.insert(
                                value,
                                (
                                    points.get(value).unwrap().0 + points.get(key).unwrap().0,
                                    points.get(value).unwrap().1 + points.get(key).unwrap().1,
                                    points.get(value).unwrap().2 + points.get(key).unwrap().2,
                                    points.get(value).unwrap().3 + points.get(key).unwrap().3
                                )
                            );
                        }
                    }
                }
                locks.insert(key, -1);
            }
        }
    }

    out_sum
}

fn main() {
    let filename = env::args().nth(1).unwrap_or_else(|| "../../AoC-input/2025/11/input.txt".to_string());
    let file = fs::read_to_string(&filename).expect("Failed to read file");
    
    println!("Part 1: {}", part1(&file));
    println!("Part 2: {}", part2(&file));
}
