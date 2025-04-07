import { assertEquals } from "@std/assert";
// js 中万物皆对象
Deno.test("4.复杂类型", () => {
  // 数组
  const arr = [1, 2, 3];
  assertEquals(typeof arr, "object");
  // 对象
  const obj = {
    name: "John",
    age: 30,
  };
  assertEquals(typeof obj, "object");
  // 函数
  const func = () => {
    return 1;
  };
  assertEquals(typeof func, "function");
  // 日期时间类
  const date = new Date();
  assertEquals(typeof date, "object");
  // 正则表达式
  const regex = /hello/;
  assertEquals(typeof regex, "object");
  // Map 映射
  const map = new Map();
  assertEquals(typeof map, "object");
  // Set 集合
  const set = new Set();
  assertEquals(typeof set, "object");
  // Promise 异步
  const promise = new Promise((resolve) => {
    resolve(1);
  });
  assertEquals(typeof promise, "object");
  // 迭代器
  const generator = function* () {
    yield 1;
  };
  assertEquals(typeof generator, "function");
  // 异步函数
  // deno-lint-ignore require-await
  const asyncFunc = async () => {
    return 1;
  };
  assertEquals(typeof asyncFunc, "function");
});
