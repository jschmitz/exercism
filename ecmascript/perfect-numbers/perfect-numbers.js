class perfectNumbers{

  classify(num){
    let arr = [1]
    let result = ""

    if(num <= 0){
      throw new Error('Classification is only possible for natural numbers.')
    }

    for(let i=2; i <= num/2; i++){
      if(num % i === 0){
        arr.push(i);
      }
    }

    let sum = arr.reduce((acc, cur) => acc + cur);

    if(sum < num || num === 1){
      result = 'deficient';
    }else if(sum === num){
      result = 'perfect';
    }else{
      result = 'abundant';
    }

    return result;
  }

}

export default perfectNumbers
