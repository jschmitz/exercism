class PhoneNumber{

  constructor(phoneNumber){
    this.phoneNumber = phoneNumber;
  }

  number(){
    let num = "";

    if (this.phoneNumber.match(/[a-z]/i)) {
      num = null;
    }else{
      num = this.phoneNumber.replace(/[^0-9]/g, '');
      if(this.validCountryCode(num)){
        num = num.substring(1);
      }else if(num.length !== 10){
        num = null;
      }
    }

    return num;
  }

  validCountryCode(n){
    return (n.length === 11 && n.charAt(0) === "1");
  }
}

export default PhoneNumber;
