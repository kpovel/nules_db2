use std::fs;

use crate::{create_file::create_file, stations::parse_csv_line};

pub fn parse_results() {
    let rusults_file = fs::read_to_string("../input/Results.csv").unwrap();
    let parsed_file = parse_content(&rusults_file);

    let output_path = "../output/results.csv";
    create_file(output_path);

    let write_res = fs::write(output_path, parsed_file);
    match write_res {
        Ok(_) => (),
        Err(e) => println!("{}", e),
    }
}

struct Measurement {
    id_measurement: String,
    time: String,
    value: String,
    id_station: String,
    id_measured_unit: String,
}

fn parse_content(file: &str) -> String {
    let mut parsed: Vec<_> = vec![String::from(
        "ID_Measurement,Time,Value,ID_Station,ID_Measured_Unit",
    )];

    for l in file.lines().skip(1) {
        let parsed_line = parse_csv_line(l);

        let measurement = Measurement {
            id_measurement: parsed_line[2].to_string(),
            time: parsed_line[0].to_string(),
            value: parsed_line[1].to_string().replace(",", "."),
            id_station: parsed_line[3].to_string(),
            id_measured_unit: parsed_line[4].to_string(),
        };

        let formatted_line = format!(
            "{},{},{},{},{}",
            measurement.id_measurement,
            measurement.time,
            measurement.value,
            measurement.id_station,
            measurement.id_measured_unit
        );

        parsed.push(formatted_line);
    }

    return parsed.join("\n");
}
