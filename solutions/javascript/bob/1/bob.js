//
// This is only a SKELETON file for the 'Bob' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const hey = (message) => {
  if (isSilence(message)) {
    return 'Fine. Be that way!';
  } else if (isShoutingQuestion(message)) {
    return 'Calm down, I know what I\'m doing!';
  } else if (isShouting(message)) {
    return 'Whoa, chill out!';
  } else if (isQuestion(message)) {
    return 'Sure.';
  } else {
    return 'Whatever.';
  }
};

const isSilence = (message) => {
  return message.trim() === '';
}

const isShouting = (message) => {
  return message.toUpperCase() === message && message.toLowerCase() !== message;
}

const isQuestion = (message) => {
  return message.trim().endsWith('?');
}

const isShoutingQuestion = (message) => {
  return isShouting(message) && isQuestion(message);
}