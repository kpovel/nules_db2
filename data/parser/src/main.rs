use optimal_value::parse_optimal_value;
use results::parse_results;
use stations::parse_stations;

mod create_file;
mod optimal_value;
mod results;
mod stations;

fn main() {
    parse_optimal_value();
    parse_stations();
    parse_results();

    println!("Files are successfully parsed!");
}
