use crate::create_file::create_file;
use std::fs;

pub fn parse_stations() {
    let stations_file = fs::read_to_string("../input/Station.csv").unwrap();
    let coordinates_file = fs::read_to_string("../input/Coordinates.csv").unwrap();

    let parsed_file = parse_files(&stations_file, &coordinates_file);
    let path = "../output/station.csv";
    create_file(path);

    let res = fs::write(path, parsed_file);

    match res {
        Ok(_) => (),
        Err(e) => println!("{}", e),
    }
}

struct Station {
    id_station: String,
    city: String,
    name: String,
    status: String,
    id_saveecobot: String,
    id_server: String,
}

struct Coordinate {
    id_station: String,
    coordinate: (String, String),
}

fn parse_files(stations_content: &str, coordinates_content: &str) -> String {
    let mut stations: Vec<Station> = vec![];
    for line in stations_content.lines().skip(1) {
        let mut tokens = parse_csv_line(line);
        if tokens[4] == "NULL" {
            tokens[4] = "1".to_string();
        }
        let station = Station {
            id_station: tokens[2].to_string(),
            city: tokens[0].to_string(),
            name: tokens[1].to_string(),
            status: tokens[3].to_string(),
            id_server: tokens[4].to_string(),
            id_saveecobot: tokens[5].to_string(),
        };
        stations.push(station);
    }

    let mut coordinates = vec![];
    for line in coordinates_content.lines().skip(1) {
        let tokens: Vec<&str> = line.split(',').collect();
        let coordinate = Coordinate {
            id_station: tokens[2].to_string(),
            coordinate: (tokens[0].to_string(), tokens[1].to_string()),
        };
        coordinates.push(coordinate);
    }

    return combine_data(&stations, &coordinates).join("\n");
}

fn combine_data(stations: &Vec<Station>, coordinates: &Vec<Coordinate>) -> Vec<String> {
    let mut combined_data =
        vec!["ID_Station,City,Name,Status,ID_SaveEcoBot,ID_Server,Coordinates".to_string()];
    for station in stations {
        let coordinate = coordinates
            .iter()
            .find(|c| c.id_station == station.id_station)
            .unwrap();
        let combined_record = format!(
            "{},{},\"{}\",{},{},{},\"{},{}\"",
            station.id_station,
            station.city,
            station.name,
            station.status,
            station.id_saveecobot,
            station.id_server,
            coordinate.coordinate.0,
            coordinate.coordinate.1
        );
        combined_data.push(combined_record);
    }

    return combined_data;
}

pub fn parse_csv_line(line: &str) -> Vec<String> {
    let mut tokens = vec![];
    let mut current_token = String::new();
    let mut in_quotes = false;
    for ch in line.chars() {
        match ch {
            '"' => in_quotes = !in_quotes,
            ',' if !in_quotes => {
                tokens.push(current_token.clone());
                current_token.clear();
            }
            _ => current_token.push(ch),
        }
    }
    if !current_token.is_empty() {
        tokens.push(current_token);
    }

    return tokens;
}
