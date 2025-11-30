pub fn pow(base: i32, exponent: i32) -> f32 {
    let mut result: f32 = 1.0;
    let loop_range: i32 = if exponent < 0 { -exponent } else { exponent };
    
    for _ in 1..=loop_range {
        result *= base as f32;
    }
    
    if exponent < 0 {
        return 1.0/result;
    }
    result
}