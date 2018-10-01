class Transcriptor {
  constructor(){
    this.complements = {
      'G': 'C',
      'C': 'G',
      'T': 'A',
      'A': 'U'
    };
  }

  toRna(strand){
    return [...strand].map(c => {
      if(this.complements[c] === undefined){
        throw new Error('Invalid input DNA.')
      }
      return this.complements[c];
    }).join('')
  }
}

export default Transcriptor
