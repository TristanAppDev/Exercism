// faster runtime, more code, more time to code
function getDigits(number) {
  let digits = [];
  while (number >= 1) {
    const lastDigit = number % 10;
    number = (number / 10) | 0;
    digits.push(lastDigit);
  }
  return digits;
}

// less time coding, less code
function getDigitsFunctional(number) {
  return [...number.toString(10)].map(Number);
}

export const isArmstrongNumber = number => {
  const digits = getDigits(number);
  const reducer = (acc, val) => acc + Math.pow(val, digits.length);
  const result = digits.reduce(reducer, 0);
  return result === number;
};