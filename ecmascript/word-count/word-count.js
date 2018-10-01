class Words{

  count(sentence){
    return sentence
      .toLowerCase()
      .trim()
      .split(/\s+/)
      .reduce(
        (acc, val) => {
          acc[val]++ || (acc[val] = 1);
          return acc;
        } , {}
      );
  }
}

export default Words
