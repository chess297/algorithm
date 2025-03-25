import { assertEquals } from "@std/assert";
import { linearSearch } from "./linear_search.ts";

Deno.test(function addTest() {
  assertEquals(linearSearch([3, 4, 5, 6], 4), 1);
});
