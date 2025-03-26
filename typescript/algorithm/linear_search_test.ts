import { assertEquals } from "@std/assert";
import { linearSearch } from "./linear_search.ts";

Deno.test(function addTest() {
  assertEquals(linearSearch([3, 4, 5, 6], 4), 1);
  assertEquals(linearSearch([3, 4, 5, 6], 1), -1);
  assertEquals(linearSearch(["a", "b", "c", 6], "d"), -1);
  assertEquals(linearSearch(["a", "b", "c", 6], "c"), 2);
});
