pub fn main() {
    let items = vec![1, 2, 3];
    // items.pop(); // 不可变实例，不可调用可变借用方法
    println!("stack len: {}", items.len()); // 不可变实例，不可调用可变借用方法
    let mut mut_item = vec![1, 2, 3];
    mut_item.pop(); // 可变实例，可调用可变借用方法
}
