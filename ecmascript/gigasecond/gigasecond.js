class Gigasecond {
  constructor(initialTime){
    this.startTime = initialTime
  }

  date(){
    return new Date(this.startTime.getTime() + Math.pow(10,12));
  }
}

export default Gigasecond
