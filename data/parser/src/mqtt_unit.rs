use std::fs;

use crate::create_file::create_file;

pub fn parse_mqtt_unit() {
    let mqtt_unit_file = fs::read_to_string("../input/MQTT_Message_Unit.csv").unwrap();
    let mut parsed_file = vec![String::from("ID_Station,ID_Measured_Unit,Message,Order")];

    for l in mqtt_unit_file.lines().skip(1) {
        let split_line: Vec<_> = l.split(",").collect();
        let formatted_string = format!(
            "{},{},{},{}",
            "2",
            split_line[1],
            split_line[2],
            split_line[3]
        );
        parsed_file.push(formatted_string);
    }

    let output_path = "../output/mqtt_unit.csv";
    create_file(output_path);

    let write_res = fs::write(output_path, parsed_file.join("\n"));
    match write_res {
        Ok(_) => (),
        Err(e) => println!("{}", e),
    }
}
