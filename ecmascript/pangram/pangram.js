class Pangram {

  constructor(sentence){
    this.sentence = sentence;
  }

  isPangram(){
    return new Set([...this.sentence.toLowerCase().replace(/[\W^_\d]/g,'')]).size === 26
  }

}

export default Pangram
