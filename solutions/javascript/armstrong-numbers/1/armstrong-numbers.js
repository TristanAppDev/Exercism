// my try to NOT convert to string
function getDigits(number) {
  let digits = [];
  while (number >= 1) {
    const lastDigit = number % 10;
    number = (number / 10) | 0;
    digits.unshift(lastDigit);
  }
  return digits;
}

export const isArmstrongNumber = number => {
  const digits = getDigits(number);
  const reducer = (acc, val) => acc + Math.pow(val, digits.length);
  const result = digits.reduce(reducer, 0);
  return result === number;
};