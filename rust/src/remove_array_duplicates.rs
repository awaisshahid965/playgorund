use std::collections::HashMap;

pub fn remove_array_duplicates(vector: &Vec<i32>) -> Vec<i32> {
    let mut values_map = HashMap::new();

    for value in vector {
        println!("test vec {}", value);
        values_map.insert(value, value);
    }

    let result_vector = values_map.keys().map(|v| **v).collect();
    return result_vector;
}
