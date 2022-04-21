const transforms: { [key: string]: string } = {
  G: "C",
  C: "G",
  T: "A",
  A: "U",
};

export function toRna(strand: string): string {
  const transformedArr = strand.split("").map((nucleotide) => {
    let rna = transforms[nucleotide];
    if (rna === undefined) throw "Invalid input DNA.";
    return rna;
  });

  return transformedArr.join("");
}
