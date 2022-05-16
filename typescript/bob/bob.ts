const isQuestion = (str: string): boolean => str.endsWith("?");
const isShouting = (str: string): boolean =>
  /[A-Z]+/.test(str) && str === str.toUpperCase();

export const hey = (message: string): string => {
  let response = "Whatever.";
  const trimmedMsg = message.trim();

  if (trimmedMsg === "") response = "Fine. Be that way!";
  else if (isQuestion(trimmedMsg) && isShouting(trimmedMsg))
    response = "Calm down, I know what I'm doing!";
  else if (isShouting(trimmedMsg)) response = "Whoa, chill out!";
  else if (isQuestion(trimmedMsg)) response = "Sure.";

  return response;
};
