const isTypeArray= (v) => {
  let gettype = Object.prototype.toString;
  return gettype.call(v) === "[object Array]";
};

const isTypeObject = (v) => {
  let gettype = Object.prototype.toString;
  return gettype.call(v) ==="[object Object]";
};

export const objectDeepCopy = (item) => {
  if (isTypeArray(item) || isTypeObject(item)) {
    return JSON.parse(JSON.stringify(item));
  } else {
    return item;
  }
};
