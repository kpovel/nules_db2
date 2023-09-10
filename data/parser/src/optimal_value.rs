use std::fs;

use crate::create_file::create_file;

pub fn parse_optimal_value() {
    let input = fs::read_to_string("../input/Optimal_Value.csv").unwrap();
    let parsed_file = parse_content(&input);

    let path = "../output/optimal_value.csv";
    create_file(path);

    let res = fs::write("../output/optimal_value.csv", parsed_file);
    match res {
        Ok(_) => (),
        Err(err) => println!("{}", err),
    }
}

fn parse_content(file: &String) -> String {
    let mut parsed = String::new();

    for l in file.lines() {
        parsed.push_str(&l.replace("NULL", ""));
        parsed.push('\n');
    }

    return parsed;
}
