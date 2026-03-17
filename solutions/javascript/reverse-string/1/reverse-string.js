export const reverseString = (text) => {
  let revText = '';
  for (let i = text.length - 1; i >= 0; i--) {
    revText += text[i];
  }
  return revText;
};