pub struct Queue<T> {
    data: Vec<T>,
}
impl<T> Queue<T> {
    pub fn new(data: Vec<T>) -> Self {
        Queue { data: data }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_queue() {
        let mut queue = Queue::new(vec![1, 2, 3, 4, 5]);
        assert_eq!(queue.data.pop(), Some(5));
        assert_eq!(queue.data.len(), 4);
    }
}
