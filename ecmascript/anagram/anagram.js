class Anagram{
  constructor(word){
    this.subject = word;
  }

  matches(candidates){
    if (candidates.constructor !== Array) {
      candidates = Array.prototype.slice.call(arguments);
    }

    let results = [];

    for (let candidate of candidates){
      let clone = this.subject.toLowerCase();
      let tCandidate = candidate.toLowerCase();

      if(clone !== tCandidate){
        let check1 = [...tCandidate].sort().join('')
        let check2 = [...clone].sort().join('')

        if(check1 === check2){
          results.push(candidate);
        }
      }
    }
    return results;
  }
}

export default Anagram
