use optimal_value::parse_optimal_value;
use stations::parse_stations;

mod create_file;
mod optimal_value;
mod stations;

fn main() {
    parse_optimal_value();
    parse_stations();

    println!("Files are successfully parsed!");
}
