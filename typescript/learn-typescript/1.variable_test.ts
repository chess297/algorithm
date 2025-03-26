import { assertEquals, assertNotEquals } from "@std/assert";
import { a, b } from "./1.variable.ts";

Deno.test("variable", () => {
  assertEquals(a, 2);
  assertNotEquals(b, 2, "b is not 2");
});
