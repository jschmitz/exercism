class Transcriptor {
  toRna(strand){
    let newStrand = "";

    [...strand].map(c => {
      console.log(c)
      switch(c){
        case 'C':
          newStrand += 'G'
          break;
        case 'G':
          newStrand += 'C'
          break;
        case 'A':
          newStrand += 'U'
          break;
        case 'T':
          newStrand += 'A'
          break;
        default:
          throw new Error('Invalid input DNA.')
      }
    })

    return newStrand;
  }
}

export default Transcriptor
