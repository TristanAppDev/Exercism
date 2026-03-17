// @ts-check

/**
 * Respond with the correct character, given the blurb, if this were said at
 * the front door.
 *
 * @param {string} blurb
 * @returns {string}
 */
export const frontDoorResponse = (blurb) => blurb[0];

/**
 * Respond with the correct character, given the blurb, if this were said at
 * the back door.
 *
 * @param {string} blurb
 * @returns {string}
 */
export function backDoorResponse(blurb) {
  const trimmedBlurb = blurb.trim();
  return trimmedBlurb[trimmedBlurb.length - 1];
}

/**
 * Capitalizes a word, meaning only the first character is a capital, and the
 * remaining letters are lower case.
 *
 * @param {string} word
 * @returns {string}
 */
const capitalize = (word) => word[0].toUpperCase() + word.slice(1).toLowerCase();

/**
 * Give the password for the front-door, given the responses.
 *
 * @param {string} responses the responses
 * @returns {string} the password
 */
export const frontDoorPassword = (responses) => capitalize(responses);

/**
 * Give the password for the back-door, given the responses.
 *
 * @param {string} responses the responses
 * @returns {string} the password
 */
export const backDoorPassword = (responses) => `${capitalize(responses)}, please`;
