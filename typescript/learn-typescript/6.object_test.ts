import { assertEquals } from "@std/assert";
import { expect } from "jsr:@std/expect";
// globalThis.name = "window name";
const obj = {
  name: "John",
  age: 30,
  address: {
    city: "New York",
    state: "NY",
  },
  getName() {
    return this.name;
  },
  getNameByArrowFunc: () => {
    return name;
  },
};

Deno.test("6.对象类型", () => {
  // 访问对象属性
  assertEquals(obj.name, "John");
  assertEquals(obj.age, 30);
  obj.name = "chess";
  // 修改对象属性
  assertEquals(obj.name, "chess");
  // 访问嵌套对象属性
  assertEquals(obj.getName(), "chess");
  assertEquals(obj.getNameByArrowFunc(), globalThis.name);
  expect(obj).toMatchObject({
    name: "chess",
    age: 30,
    address: {
      city: "New York",
      state: "NY",
    },
  });
});
