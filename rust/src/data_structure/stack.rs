pub struct Stack<T> {
    items: Vec<T>,
}

impl<T> Stack<T> {
    pub fn new(items: Vec<T>) -> Self {
        Stack { items: items }
    }

    // pub fn pop(&self) -> Option<T> {}
    pub fn push(&mut self, item: T) {
        self.items.push(item);
    }

    pub fn pop(&mut self) -> Option<T> {
        self.items.pop()
    }

    pub fn len(&self) -> usize {
        self.items.len()
    }

    pub fn is_empty(&self) -> bool {
        self.items.is_empty()
    }
}

pub fn main() {
    let mut stack = Stack::new(vec![1, 2, 3]);
    stack.push(4);
    stack.pop();
    stack.len();
    stack.is_empty();
}
