/* https://www.hackerrank.com/challenges/crush/problem

SPDX-FileCopyrightText: 2022 Vladimir Rusinov
SPDX-License-Identifier: Apache-2.0

Starting with a 1-indexed array of zeros and a list of operations, for each
operation add a value to each the array element between two given indices,
inclusive. Once all operations have been performed, return the maximum value in
the array.
*/

use std::env;
use std::fs::File;
use std::io::{self, BufRead, Write};

/*
 * Complete the 'array_manipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

fn array_manipulation_slow(n: i32, queries: &[Vec<i32>]) -> i64 {
    // Given constrains is 3 <= n <= 10**7.
    // That's around 80 MiB of memory for array of 10**7 i64s.
    // Having 80 MiB array may not be the best, but it should still work.

    let mut array: Vec<i64> = vec![0; n as usize];
    for query in queries {
        let mut a = query[0]-1;
        let b = query[1]-1;
        let k = query[2];
        // However this loop makes in n^2, which times out.
        while a <= b {
            array[a as usize] += k as i64;
            a += 1;
        }
    }

    return *array.iter().max().unwrap();
}

//
fn array_manipulation(n: i32, queries: &[Vec<i32>]) -> i64 {
    // Given constrains is 3 <= n <= 10**7.
    // That's around 80 MiB of memory for array of 10**7 i64s.
    // Having 80 MiB array may not be the best, but it should still work.
    //
    // This time, use https://en.wikipedia.org/wiki/Prefix_sum to avoid inner
    // loop.

    let mut array: Vec<i64> = vec![0; (n+1) as usize];
    for query in queries {
        let a = query[0]-1;
        let b = query[1]-1;
        let k = query[2];
        array[a as usize] += k as i64;
        array[(b+1) as usize] -= k as i64;
    }
    let mut current_sum: i64 = 0;
    let mut max_sum: i64 = 0;
    for a in array {
        current_sum += a;
        if current_sum > max_sum {
            max_sum = current_sum
        }
    }
    return max_sum;
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let mut fptr = File::create(env::var("OUTPUT_PATH").unwrap()).unwrap();

    let first_multiple_input: Vec<String> = stdin_iterator.next().unwrap().unwrap()
        .split(' ')
        .map(|s| s.to_string())
        .collect();

    let n = first_multiple_input[0].trim().parse::<i32>().unwrap();

    let m = first_multiple_input[1].trim().parse::<i32>().unwrap();

    let mut queries: Vec<Vec<i32>> = Vec::with_capacity(m as usize);

    for i in 0..m as usize {
        queries.push(Vec::with_capacity(3_usize));

        queries[i] = stdin_iterator.next().unwrap().unwrap()
            .trim_end()
            .split(' ')
            .map(|s| s.to_string().parse::<i32>().unwrap())
            .collect();
    }

    let result = array_manipulation(n, &queries);

    writeln!(&mut fptr, "{}", result).ok();
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_example() {
        let queries = [vec!(1, 5, 3), vec!(4, 8, 7), vec!(6, 9, 1)];
        assert_eq!(crate::array_manipulation(10, &queries), 10);
    }

    #[test]
    fn test_sample() {
        let queries = [vec!(1, 2, 100), vec!(2, 5, 100), vec!(3, 4, 100)];
        assert_eq!(crate::array_manipulation(5, &queries), 200);
    }
}
