pub struct Queue<T> {
    data: Vec<T>,
}
impl<T> Queue<T> {
    pub fn new(data: Vec<T>) -> Self {
        Queue { data: data }
    }
}

pub fn main() {
    let q = Queue::new(vec![1, 2, 3]);
    println!("{}", q.data.len());
}
