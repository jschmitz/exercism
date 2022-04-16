const resistors: { [key: string]: number } = {
  black: 0,
  brown: 1,
  red: 2,
  orange: 3,
  yellow: 4,
  green: 5,
  blue: 6,
  violet: 7,
  grey: 8,
  white: 9,
};

export function decodedResistorValue(colors: Array<string>) {
  const zeros = "0".repeat(resistors[colors[2]]);
  let ohms = parseInt(`${resistors[colors[0]]}${resistors[colors[1]]}${zeros}`);
  let label = "ohms";

  if (ohms > 999) {
    ohms = ohms / 1000;
    label = "kiloohms";
  }

  return `${ohms} ${label}`;
}
