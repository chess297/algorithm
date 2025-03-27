// 选择排序
export function selectSort<T>(arr: T[]): T[] {
  for (let i = 0; i < arr.length; i++) {
    for (let j = i + 1; j < arr.length; j++) {
      if (arr[j] < arr[i]) {
        const temp = arr[j];
        arr[j] = arr[i];
        arr[i] = temp;
      }
    }
  }
  return arr;
}

// 插入排序
export function insertSort<T>(arr: T[]): T[] {
  for (let i = 1; i < arr.length; i++) {
    for (let j = i - 1, temp = arr[i]; j > 0; j--) {
      if (arr[j] > temp) {
        arr[i] = arr[j];
      } else {
        arr[j] = temp;
        break;
      }
    }
  }
  return arr;
}
