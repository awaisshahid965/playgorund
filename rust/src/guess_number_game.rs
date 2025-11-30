use std::io;
use std::cmp::Ordering;
use rand::{thread_rng, Rng};
use colored::*;

pub fn guess_number_game() {
    println!("Welcome to Game");
    println!("Where you guess Number to Win...\n\n");

    let mut rng = thread_rng();
    let secret_number: u32 = rng.gen_range(1..100);
    println!("Secret Number is: {}", secret_number.to_string().dimmed());

    loop {
        println!("Enter your Guess:");
        let mut user_input: String = String::new();
        io::stdin().read_line(&mut user_input).expect("Invalid Input!");

        let user_input_guess: u32 = match user_input.trim().parse() {
            Ok(num) => num,
            Err(val) => {
                println!("{}", val.to_string().red());
                println!("{}", "Retry entering valid guess number...".blue());
                continue;
            },
        };

        match user_input_guess.cmp(&secret_number) {
            Ordering::Greater => println!("{}", "guessed value greater that generated number!".red()),
            Ordering::Less => println!("{}", "guessed value less that generated number!".red()),
            Ordering::Equal => {
                println!("{}", "You Win!".green());
                break;
            },
        }
    }
}
