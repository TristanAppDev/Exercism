// @ts-check

const birdsPerDay = [2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1];

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export const totalBirdCount = birdsPerDay =>
  birdsPerDay.reduce((totalVal, currentVal) => totalVal + currentVal, 0);

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export const birdsInWeek = (birdsPerDay, week) =>
  birdsPerDay
    .slice(weekStartIndex(week), weekEndIndex(week))
    .reduce((totalVal, currentVal) => totalVal + currentVal, 0);

const weekStartIndex = (week) => (week - 1) * 7;
const weekEndIndex = (week) => weekStartIndex(week) + 7;

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {number[]} corrected bird count data
 */
export function fixBirdCountLog(birdsPerDay) {
  for (let i = 0; i < birdsPerDay.length; i++) {
    if (i % 2 === 0) birdsPerDay[i]++;
  }
  return birdsPerDay;
}
