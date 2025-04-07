import { assertEquals } from "@std/assert";

Deno.test("3.基本类型", () => {
  const a = 1; // number
  assertEquals(typeof a, "number");
  const b = 1.1; // number
  assertEquals(typeof b, "number");
  const c = "hello"; // string
  assertEquals(typeof c, "string");
  const d = true; // boolean
  assertEquals(typeof d, "boolean");
  const e = null; // null
  assertEquals(typeof e, "object");
  const f = undefined; // undefined
  assertEquals(typeof f, "undefined");
  const g = Symbol("hello"); // symbol
  assertEquals(typeof g, "symbol");
  const h = BigInt(1); // bigint
  assertEquals(typeof h, "bigint");
});
