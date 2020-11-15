// Jumping on the Clouds
//
// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem

fn jump(current_position: usize, c: &Vec<i32>, n: i32, jumps: i32) -> Option<i32> {
    // Overjumped?
    if current_position >= n as usize {
        return None;
    }

    // Bad cloud?
    if c[current_position] == 1 {
        return None;
    }

    // Are we at the end?
    if current_position == (n-1) as usize {
        return Some(0);
    }

    // Try jumping 2
    match jump(current_position+2, c, n, jumps) {
        Some(val) => {
            return Some(jumps+val+1);
        },
        None => {}
    }
    // Try jumping 1
    match jump(current_position+1, c, n, jumps) {
        Some(val) => {
            return Some(jumps+val+1);
        },
        None => {}
    }

    return None;
}

fn jumping_on_clouds(n: i32, c: Vec<i32>) -> i32 {
    let current_position: usize = 0;
    return jump(current_position, &c, n, 0).unwrap();
}

fn main() -> std::io::Result<()> {
    use std::io::{self, BufRead, Write};
    use std::fs::File;

    let stdin = io::stdin();
    let mut cin_iterator = stdin.lock().lines();

    let buffer = cin_iterator.next().unwrap().unwrap();
    let n: i32 = buffer.trim().parse().expect("input was not an integer");

    let buffer = cin_iterator.next().unwrap().unwrap();
    let c = buffer
        .split_whitespace()
        .map(|x| x.parse::<i32>())
        .collect::<Result<Vec<i32>, _>>()
        .unwrap();

    let ret = jumping_on_clouds(n, c);

    let mut file: Box<dyn Write> = match std::env::var("OUTPUT_PATH") {
         Ok(val) => Box::new(File::create(val)?),
         Err(_e) => Box::new(io::stdout()),
    };
    file.write_all(format!("{}\n", ret).as_bytes())?;

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test00() {
        assert_eq!(jumping_on_clouds(7, vec![0, 0, 1, 0, 0, 1, 0]), 4);
    }

    #[test]
    fn test01() {
        assert_eq!(jumping_on_clouds(6, vec![0, 0, 0, 1, 0, 0]), 3);
    }
}
