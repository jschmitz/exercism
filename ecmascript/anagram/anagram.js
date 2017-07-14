class Anagram{
  constructor(word){
    this.subject = word;
  }

  matches(candidates){
    if (candidates.constructor !== Array) {
      candidates = Array.prototype.slice.call(arguments);
    }

    return candidates
      .filter(c => this.subject.toLowerCase() !== c.toLowerCase())
      .filter(c => [...c.toLowerCase()].sort().join() === [...this.subject.toLowerCase()].sort().join())

  }
}
export default Anagram
