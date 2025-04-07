import { assertEquals } from "@std/assert";
import { expect } from "jsr:@std/expect";

// 普通函数
function add(a: number, b: number) {
  return a + b;
}
// 箭头函数
const sub = (a: number, b: number) => {
  return a - b;
};

// Promise
const ps = new Promise((resolve) => {
  resolve(1);
});
// 迭代器
const generator = function* () {
  yield 1;
};

// deno-lint-ignore require-await
const asyncFunc = async () => {
  return 1;
};

Deno.test("5.函数", async () => {
  expect(add(1, 2)).toBe(3);
  expect(sub(1, 2)).toBe(-1);
  assertEquals(add.prototype, {}); // 普通函数的 prototype 是一个空对象
  expect(sub.prototype).toBeUndefined(); // 箭头函数的 prototype 是 undefined
  expect(
    // 立即执行函数 匿名函数
    (function () {
      return 1;
    })()
  ).toBe(1);
  let res = await ps;
  assertEquals(res, 1);
  for (const item of generator()) {
    assertEquals(item, 1);
  }
  res = await asyncFunc();
  assertEquals(res, 1);
});
