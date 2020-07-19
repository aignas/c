type R<T> = Result<T, String>;

pub fn score(input: &str) -> R<usize> {
    Ok(parse(input)?.iter().map(|f| f.sum()).sum())
}

#[cfg(test)]
mod score_tests {
    use super::score;

    #[test]
    fn bowling_test() {
        assert_eq!(0, score("--------------------").unwrap());
        assert_eq!(30, score("3-3-3-3-3-3-3-3-3-3-").unwrap());
        assert_eq!(50, score("5-5-5-5-5-5-5-5-5-5-").unwrap());
        assert_eq!(100, score("-/-/-/-/-/-/-/-/-/-/-").unwrap());
        assert_eq!(150, score("5/5/5/5/5/5/5/5/5/5/5").unwrap());
        assert_eq!(70, score("43434343434343434343").unwrap());
        assert_eq!(300, score("XXXXXXXXXXXX").unwrap());
        assert_eq!(8 * 30 + 23 + 20, score("XXXXXXXXXX3/").unwrap());
    }
}

fn parse(score: &str) -> R<Vec<Frame>> {
    let mut throws = score
        .chars()
        .map(to_digit)
        .map(|d| d.expect("unexpected input"))
        .collect::<Vec<usize>>();
    match throws.last() {
        Some(10) => {}
        _ => throws.push(0),
    };

    let mut skip = false;
    Ok(throws
        .windows(3)
        .map(|c| {
            if skip {
                skip = false;
                return None;
            }

            let frame = match c {
                &[10, 10, _] => Frame::Strike(c[1], c[2]),
                &[10, _, 10] => Frame::Strike(c[1], c[2] - c[1]),
                &[_, 10, _] => Frame::Spare(c[2]),
                &[_, _, _] => Frame::Some(c[0], c[1]),
                _ => todo!(),
            };

            if let Frame::Strike(_, _) = &frame {
                //
            } else {
                skip = true;
            }

            Some(frame)
        })
        .flatten()
        .collect())
}

fn to_digit(c: char) -> Option<usize> {
    match c {
        '-' => Some(0),
        '/' => Some(10),
        'X' => Some(10),
        _ => c.to_digit(10).map(|c| c as usize),
    }
}

#[cfg(test)]
mod parse_tests {
    use super::parse;

    #[test]
    fn parse_test() {
        assert_eq!(10, parse("-/-/-/-/-/-/-/-/-/-/-").unwrap().len());
        assert_eq!(10, parse("5-5-5-5-5-5-5-5-5-5-").unwrap().len());
        assert_eq!(10, parse("--------------------").unwrap().len());
    }
}

enum Frame {
    Strike(usize, usize),
    Spare(usize),
    Some(usize, usize),
}

impl Frame {
    fn sum(&self) -> usize {
        match self {
            Frame::Strike(a, b) => 10 + a + b,
            Frame::Spare(a) => 10 + a,
            Frame::Some(a, b) => a + b,
        }
    }
}
