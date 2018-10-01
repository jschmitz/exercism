class Bob{
  hey(message){
    let response =''
    const punc = [...message][message.length -1];
    const allUpper = [...message].every( c => c === c.toUpperCase());
    const hasChar = [...message].some( c => c.toUpperCase() !== c.toLowerCase())

    if(message.trim().length === 0){
      response = 'Fine. Be that way!';
    }else if(allUpper && hasChar){
      response = 'Whoa, chill out!';
    }else if(punc === '?'){
      response = 'Sure.';
    }else{
      response = 'Whatever.';
    }
    return response;
  }

}

export default Bob
