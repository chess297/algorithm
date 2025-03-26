#[cfg(test)]
mod tests {
    struct Queue<T> {
        data: Vec<T>,
    }
    impl<T> Queue<T> {
        fn new(data: Vec<T>) -> Self {
            Queue { data: data }
        }
    }

    #[test]
    fn test_queue() {
        let mut queue = Queue::new(vec![1, 2, 3, 4, 5]);
        assert_eq!(queue.data.pop(), Some(5));
        assert_eq!(queue.data.len(), 4);
    }
}
