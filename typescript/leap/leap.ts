export function isLeap(year: number) {
  let leap = false;

  if (year % 4 === 0) leap = true;
  if (year % 100 === 0 && year % 400 !== 0) leap = false;

  return leap;
}
