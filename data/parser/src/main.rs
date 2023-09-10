use optimal_value::parse_optimal_value;
use results::parse_results;
use stations::parse_stations;

use crate::mqtt_unit::parse_mqtt_unit;

mod create_file;
mod optimal_value;
mod results;
mod stations;
mod mqtt_unit;

fn main() {
    parse_optimal_value();
    parse_stations();
    parse_results();
    parse_mqtt_unit();

    println!("Files are successfully parsed!");
}
