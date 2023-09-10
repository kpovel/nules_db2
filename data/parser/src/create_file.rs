use std::{path::Path, fs};

pub fn create_file(path: &str) {
    let path = Path::new(path);
    let prefix = path.parent().unwrap();
    fs::create_dir_all(prefix).unwrap();
}
