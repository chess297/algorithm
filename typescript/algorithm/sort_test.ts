import { assertEquals } from "@std/assert";
import { selectSort } from "./sort.ts";

Deno.test("SelectSortTest", () => {
  const arr = [5, 3, 2, 4, 1];
  selectSort(arr);
  assertEquals(arr[0], 1, "arr[0] should be 1");
  assertEquals(arr[1], 2, "arr[1] should be 2");
  assertEquals(arr[2], 3, "arr[2] should be 3");
  assertEquals(arr[3], 4, "arr[3] should be 4");
  assertEquals(arr[4], 5, "arr[4] should be 5");
});
