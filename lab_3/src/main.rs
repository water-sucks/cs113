use std::io;
use std::io::Write;

#[derive(Debug, PartialEq, Eq, Hash)]
struct InputPair {
    letter: &'static str,
    number: u8,
}

#[derive(Debug)]
struct VendingMachineItem {
    name: &'static str,
    count: u8,
    price: f64,
}

// Named tuple struct for truth table row
#[derive(Debug)]
struct TruthTableInput(u8, u8, u8);

const CHARACTER_INPUTS: [&str; 6] = ["A", "B", "C", "D", "E", "F"];
const VALID_INPUTS: [InputPair; 4] = [
    InputPair {
        letter: "A",
        number: 5,
    },
    InputPair {
        letter: "B",
        number: 2,
    },
    InputPair {
        letter: "D",
        number: 9,
    },
    InputPair {
        letter: "F",
        number: 7,
    },
];

const VALID_INPUT_TABLE_INDICES: [usize; 4] = [1, 3, 5, 6];
const TRUTH_TABLE_INPUTS: [TruthTableInput; 8] = [
    TruthTableInput(0, 0, 0),
    TruthTableInput(0, 0, 1),
    TruthTableInput(0, 1, 0),
    TruthTableInput(0, 1, 1),
    TruthTableInput(1, 0, 0),
    TruthTableInput(1, 0, 1),
    TruthTableInput(1, 1, 0),
    TruthTableInput(1, 1, 1),
];

fn main() {
    let mut vending_machine: [VendingMachineItem; 8] = [
        VendingMachineItem {
            name: "",
            count: 0,
            price: 0.00,
        },
        VendingMachineItem {
            name: "Coke",
            count: 5,
            price: 1.50,
        },
        VendingMachineItem {
            name: "",
            count: 0,
            price: 0.00,
        },
        VendingMachineItem {
            name: "Pepsi",
            count: 3,
            price: 1.75,
        },
        VendingMachineItem {
            name: "",
            count: 0,
            price: 0.00,
        },
        VendingMachineItem {
            name: "Water",
            count: 10,
            price: 1.00,
        },
        VendingMachineItem {
            name: "7UP",
            count: 3,
            price: 1.25,
        },
        VendingMachineItem {
            name: "",
            count: 0,
            price: 0.00,
        },
    ];

    'main: loop {
        print_keypad();
        let letter: &str;
        let number: u8;

        loop {
            println!("Enter a character (A, B, C, D, E, or F; CLR to clear):");
            let raw = get_raw_input();
            match raw {
                Some(input) => match parse_letter(input) {
                    Some(parsed) => {
                        letter = parsed;
                        break;
                    }
                    None => {
                        println!("Invalid input, must be a letter between A-F or CLR to restart.");
                        continue;
                    }
                },
                None => {
                    println!("Cleared input");
                    continue 'main;
                }
            }
        }

        loop {
            println!("Enter a number (0-9; CLR to clear):");
            let raw = get_raw_input();
            match raw {
                Some(input) => match parse_number(input) {
                    Some(parsed) => {
                        number = parsed;
                        break;
                    }
                    None => {
                        println!("Invalid input, must be a number between 0-9 or CLR to restart.");
                        continue;
                    }
                },
                None => {
                    println!("Cleared input");
                    continue 'main;
                }
            };
        }

        let input = InputPair { letter, number };

        match VALID_INPUTS.iter().position(|e| e == &input) {
            Some(index) => {
                let truth_table_index = VALID_INPUT_TABLE_INDICES[index];
                if TRUTH_TABLE_INPUTS[truth_table_index].eval() {
                    let item = &mut vending_machine[truth_table_index];
                    match item.transaction() {
                        Ok(change) => {
                            println!("Your change is ${:.2}.", change);
                            println!("Thanks for using this vending machine!");
                            continue;
                        }
                        Err(e) => {
                            println!("Oh no! Something went wrong. {}", e);
                            continue;
                        }
                    }
                }
            }
            None => println!("Not available."),
        }
    }
}

fn print_keypad() {
    println!("-------------");
    println!("| A | 1 | 2 |");
    println!("-------------");
    println!("| B | 3 | 4 |");
    println!("-------------");
    println!("| C | 5 | 6 |");
    println!("-------------");
    println!("| D | 7 | 8 |");
    println!("-------------");
    println!("| E | 9 | 0 |");
    println!("---------------");
    println!("| F | * | CLR |");
    println!("---------------");
}

fn get_raw_input() -> Option<String> {
    print!(">>> ");
    io::stdout().flush().unwrap();

    let mut input = String::new();

    io::stdin().read_line(&mut input).unwrap();

    let input = input.trim().to_uppercase();

    if input == "CLR" {
        None
    } else {
        Some(input.to_string())
    }
}

fn parse_letter(input: String) -> Option<&'static str> {
    match CHARACTER_INPUTS.iter().position(|&e| e == input.as_str()) {
        Some(r) => Some(CHARACTER_INPUTS[r]),
        None => None,
    }
}

fn parse_number(input: String) -> Option<u8> {
    match input.parse::<u8>() {
        Ok(num) => {
            if (0..10).contains(&num) {
                Some(num)
            } else {
                None
            }
        }
        Err(_) => None,
    }
}

fn parse_money(input: String) -> Option<f64> {
    match input.parse::<f64>() {
        Ok(num) => {
            if num > 0.0 {
                Some(num)
            } else {
                None
            }
        }
        Err(_) => None,
    }
}

// boolean expression: (PQ~R) v (~PR) v (~QR)
impl TruthTableInput {
    fn eval(&self) -> bool {
        let TruthTableInput(c1, c2, c3) = self;
        if (c1 & c2 & !c3) | (!c1 & c3) | (!c2 & c3) == 1 {
            true
        } else {
            false
        }
    }
}

impl VendingMachineItem {
    pub fn transaction(&mut self) -> Result<f64, &'static str> {
        loop {
            if self.count < 1 {
                return Err("The selected item is not in stock.");
            }
            println!(
                "{} is in stock! There are {} left. The price is ${:.2}",
                self.name, self.count, self.price
            );
            println!("You may now insert money into the machine.");
            println!("How much are you inserting? (Enter without the dollar sign, i.e. 1.50):");

            let raw = get_raw_input();
            match raw {
                Some(input) => match parse_money(input) {
                    Some(parsed) => {
                        if parsed < self.price {
                            return Err("Not enough money was inserted.");
                        }
                        self.count -= 1;
                        return Ok(parsed - self.price);
                    }
                    None => {
                        println!("Invalid input, can't insert no money.");
                        continue;
                    }
                },
                None => {
                    println!("Unable to receive money.");
                }
            }
        }
    }
}
