use std::cmp::max;

fn minimum_bribes(_t: i32, queue: Vec<i32>) -> Result<i32, &'static str> {
    // Just loop over the vector and check for out of place people.

    let mut result: i32 = 0;

    // https://doc.rust-lang.org/stable/std/iter/struct.Enumerate.html
    for (i, p) in queue.iter().enumerate() {
        if p > &(i as i32 + 1) {
            let bribes = p - (i as i32 + 1);
            if bribes > 2 {
                return Err("Too chaotic");
            }
        }

        // Calculate how many people `p` received bribes from.
        // Need to count everyone whose ticket is higher than p ahead of us.
        // Only need to start from p original position.
        // max in case p is 1.
        println!("{}, {}", p-2, i);
        for j in max(p-2, 0)..i as i32 {
            if queue[j as usize] > *p {
                result += 1;
            }
        }
    }
    return Ok(result);
}

fn main() -> std::io::Result<()> {
    use std::io::{self, BufRead, Write};
    use std::fs::File;

    let stdin = io::stdin();
    let mut cin_iterator = stdin.lock().lines();

    let mut file: Box<dyn Write> = match std::env::var("OUTPUT_PATH") {
        Ok(val) => Box::new(File::create(val)?),
        Err(_e) => Box::new(io::stdout()),
   };

    // Input format:
    // The first line contains an integer `t`, the number of test cases.
    // Each of the next `t` pairs of lines are as follows:
    // - The first line contains an integer `t`, the number of people in the queue
    // - The second line has `n` space-separated integers describing the final state of the queue.

    let buffer = cin_iterator.next().unwrap().unwrap();
    let num_testcases: i32 = buffer.trim().parse().expect("input was not an integer");

    for _ in 0..num_testcases {
        let buffer = cin_iterator.next().unwrap().unwrap();
        let t: i32 = buffer.trim().parse().expect("input was not an integer");
        let buffer = cin_iterator.next().unwrap().unwrap();
        let queue = buffer
            .split_whitespace()
            .map(|x| x.parse::<i32>())
            .collect::<Result<Vec<i32>, _>>()
            .unwrap();
        let ret = minimum_bribes(t, queue);
        match ret {
            Ok(v) => file.write_all(format!("{}\n", v).as_bytes())?,
            Err(e) => file.write_all(format!("{}\n", e).as_bytes())?,
        }
    }
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test00() -> Result<(), String> {
        assert_eq!(minimum_bribes(8, vec![1, 2, 3, 5, 4, 6, 7, 8])?, 1);
        Ok(())
    }

    #[test]
    fn test1() -> Result<(), String> {
        assert_eq!(minimum_bribes(5, vec![2, 1, 5, 3, 4])?, 3);
        Ok(())
    }

    #[test]
    fn test2() -> Result<(), String> {
        assert!(minimum_bribes(5, vec![2, 5, 1, 3, 4]).is_err());
        Ok(())
    }

    #[test]
    fn test3() -> Result<(), String> {
        assert_eq!(minimum_bribes(8, vec![1, 2, 5, 3, 7, 8, 6, 4])?, 7);
        Ok(())
    }

    #[test]
    fn test_no_bribes() -> Result<(), String> {
        assert_eq!(minimum_bribes(8, vec![1, 2, 3, 4, 5, 6, 7, 8])?, 0);
        Ok(())
    }
}
