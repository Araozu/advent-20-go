use std::{collections::HashMap, fs::read_to_string};

mod computer;

fn main() {
    // current time
    let start = std::time::Instant::now();

    let input = read_to_string("input.txt").expect("Input file not found");
    let mut computer = computer::Computer::new(input);

    let computer_len = computer.len();

    'outer: for i in 0..(computer_len) {
        // If the instruction at position `i` can't be flipped
        if !computer.jmp_or_nop_at(i) {
            continue;
        }

        // Stores the indexes already executed.
        let mut map: HashMap<usize, bool> = HashMap::new();
        computer.flip(i);

        loop {
            let current_ip = computer.get_current_ip();

            // Success condition
            if current_ip == computer_len {
                println!("Result: {}", computer.get_acc());
                break 'outer;
            }

            // Loop detected
            if map.contains_key(&current_ip) {
                break;
            }

            // Store current ip to keep track of what instructions have been executed
            map.insert(current_ip, true);

            // Run the current instruction
            computer.step();
        }

        // Reset the computer state
        computer.reset();
    }

    // Print the time it took to run the program
    println!("Time: {} micros", start.elapsed().as_micros());

    println!("No result was found");
}
