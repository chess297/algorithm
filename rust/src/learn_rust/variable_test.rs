#[cfg(test)]
mod tests {
    #[test]
    fn test_variable() {
        let mut _x: i32 = 1;
        _x = 2;
        let y = 3;
        let z = _x + y;
        assert_eq!(_x, 2);
        assert_eq!(z, 5);
    }
}
