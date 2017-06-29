class Beer {

  static verse(num){
    let v;

    if(num === 0){
      v = `No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
`
    }else if(num === 1){
      v =`1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.
`;
    }else if(num === 2){
      v = `${num} bottles of beer on the wall, ${num} bottles of beer.\nTake one down and pass it around, ${num -1} bottle of beer on the wall.\n`
    }else{
      v = `${num} bottles of beer on the wall, ${num} bottles of beer.\nTake one down and pass it around, ${num -1} bottles of beer on the wall.\n`
    }
    return v;
  }

  static sing(start = 99, end = 0){
    let s = "";

    for(; start >= end; start--){
      s = s + Beer.verse(start);
      if(start !== end){
        s = s + "\n";
      }
    }

    return s;
  }
}

export default Beer
