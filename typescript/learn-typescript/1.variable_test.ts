import { assertEquals, assertNotEquals } from "@std/assert";
let a = 1;
let b = 1;
a = 2;
b = 3;

Deno.test("1.变量", () => {
  assertEquals(a, 2);
  assertNotEquals(b, 2, "b is not 2");
});
