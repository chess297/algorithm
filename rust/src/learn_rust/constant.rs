#[cfg(test)]
mod tests {
    const RUST_NAME: &str = "rust";
    #[test]
    fn test_constant() {
        assert_eq!(RUST_NAME, "rust");
    }
}
