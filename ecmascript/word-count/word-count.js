class Words{

  count(sentence){
    sentence = sentence.toLowerCase().trim();
    let counts = new Map();
    let words = sentence.split(/\s+/);

    [...words].map((word) => {
      if(counts.has(word)){
        counts.set(word, counts.get(word) +1);
      }else{
        counts.set(word, 1);
      }
    });
    return strMapToObj(counts);
  }


}
  function strMapToObj(strMap) {
    let obj = Object.create(null);
    for (let [k,v] of strMap) {
        // We don’t escape the key '__proto__'
        // which can cause problems on older engines
        obj[k] = v;
    }
    return obj;
  }

export default Words
