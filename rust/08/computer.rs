pub struct Computer {
    instructions: Vec<Instruction>,
    ip: usize,
    acc: i32,
    flipped_idx: Option<usize>,
}

enum Instruction {
    Acc(i32),
    Jmp(i32),
    Nop(i32),
}

impl Computer {
    /// Creates a new computer from a string containing intructions
    pub fn new(input: String) -> Computer {
        let instructions: Vec<Instruction> = input
            .split("\n")
            .map(|line| {
                let inst = &line[..3];
                let sign_value = match &line[4..5] {
                    "+" => 1,
                    "-" => -1,
                    _ => panic!("Expected + or -, got something else"),
                };
                let value_str = &line[5..];
                let value = sign_value
                    * value_str
                        .parse::<i32>()
                        .expect(format!("Found an invalid number: {}", value_str).as_str());

                match inst {
                    "acc" => Instruction::Acc(value),
                    "jmp" => Instruction::Jmp(value),
                    "nop" => Instruction::Nop(value),
                    _ => panic!("Found an invalid instruction: {}", inst),
                }
            })
            .collect();

        Computer {
            instructions,
            ip: 0,
            acc: 0,
            flipped_idx: None,
        }
    }

    /// Returns the number of instructions of this computer
    pub fn len(&self) -> usize {
        self.instructions.len()
    }

    /// Executes the instruction at the current ip
    pub fn step(&mut self) {
        match self.instructions.get(self.ip) {
            Some(Instruction::Acc(value)) => {
                self.acc += value;
                self.ip += 1;
            }
            Some(Instruction::Jmp(value)) => {
                let ip_i32: i32 = self.ip.try_into().expect("Found an invalid ip");
                let new_value = ip_i32 + value;

                self.ip = new_value.try_into().expect("Tried to assign an invalid ip");
            }
            Some(Instruction::Nop(_)) => {
                self.ip += 1;
            }
            None => {
                panic!("Tried to execute an out of bounds instruction.")
            }
        }
    }

    /// Flips a JMP into NOP, or NOP into JMP instruction at position `idx`.
    ///
    /// If this function is called more than once, it panics.
    pub fn flip(&mut self, idx: usize) {
        match self.flipped_idx {
            Some(_) => panic!("An instruction on this computer has already been flipped."),
            None => {
                self.flipped_idx = Some(idx);
                self.flip_instruction(idx);
            }
        }
    }

    /// Internal implementation of flip
    fn flip_instruction(&mut self, idx: usize) {
        match self.instructions.get(idx) {
            Some(Instruction::Jmp(v)) => {
                self.instructions[idx] = Instruction::Nop(*v);
            }
            Some(Instruction::Nop(v)) => {
                self.instructions[idx] = Instruction::Jmp(*v);
            }
            Some(_) => {}
            None => {
                panic!(
                    "Tried to flip an instruction on an out of bounds index. {}",
                    idx
                )
            }
        }
    }

    /// Returns whether the instruction at `idx` is JMP or NOP.
    pub fn jmp_or_nop_at(&mut self, idx: usize) -> bool {
        match self.instructions.get(idx) {
            Some(Instruction::Nop(_)) => true,
            Some(Instruction::Jmp(_)) => true,
            _ => false,
        }
    }

    /// Returns the current ip
    pub fn get_current_ip(&self) -> usize {
        self.ip
    }

    /// Returns the current acc
    pub fn get_acc(&self) -> i32 {
        self.acc
    }

    /// Resets the state of the computer, including any JMP/NOP flip.
    pub fn reset(&mut self) {
        self.ip = 0;
        self.acc = 0;

        if let Some(idx) = self.flipped_idx {
            self.flipped_idx = None;
            self.flip_instruction(idx);
        }
    }
}
