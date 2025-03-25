pub const RUST_NAME: &str = "rust";

pub fn main() {
    println!("Hello, {}!", RUST_NAME);
    let items = vec![1, 2, 3];
    println!("stack len: {}", items.len());
}
