import { assertEquals } from "@std/assert";

export const TS_NAME = "TypeScript";

Deno.test("2.常量", () => {
  assertEquals(TS_NAME, "TypeScript");
});
