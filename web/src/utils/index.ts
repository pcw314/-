/**
 * 把对象转为formData
 */
export function objToFormData(obj: Recordable) {
  const formData = new FormData();
  Object.keys(obj).forEach((key) => {
    formData.append(key, obj[key]);
  });
  return formData;
}
const toString = Object.prototype.toString;

export const is = (val: unknown, type: string) => {
  return toString.call(val) === `[object ${type}]`;
};
export const isString = (val: unknown): val is string => {
  return is(val, "String");
};
