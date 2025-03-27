#[cfg(test)]
mod tests {
    #[test]
    fn test_string() {
        let s = String::from("hello");
        let s2 = s + " world";
        assert_eq!(s2, "hello world");
    }
}
