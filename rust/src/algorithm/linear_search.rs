pub fn linear_search<T: PartialEq>(arr: &[T], target: &T) -> Option<usize> {
    for i in 0..arr.len() {
        if arr[i] == *target {
            return Some(i);
        }
    }
    None
}
#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_linear_search() {
        let arr = [1, 2, 3, 4, 5];
        assert_eq!(linear_search(&arr, &3), Some(2));
        assert_eq!(linear_search(&arr, &6), None);
        let arr2 = ["1", "2", "3", "4", "5"];

        assert_eq!(linear_search(&arr2, &"1"), Some(0));
    }
}
