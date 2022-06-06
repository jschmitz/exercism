export class Clock {
  readonly minutesInDay = 1440;
  readonly minutesInHour = 60;
  readonly hoursInDay = 24;

  minutes = 0;

  constructor(hour: number, minute?: number) {
    minute = minute || 0;
    this.minutes = hour * this.minutesInHour + minute;
  }

  public toString(): string {
    this.minutes = this.adjustNegative(this.minutes);
    return `${this.hoursString()}:${this.minutesString()}`;
  }

  public plus(minutes: number): Clock {
    this.minutes = this.minutes + minutes;
    return this;
  }

  public minus(minutes: number): Clock {
    this.minutes = this.minutes - minutes;
    return this;
  }

  public equals(other: Clock): boolean {
    return other.toString() === this.toString();
  }

  private hoursString(): string {
    return `${Math.floor(
      (this.minutes / this.minutesInHour) % this.hoursInDay
    )}`.padStart(2, "0");
  }

  private minutesString(): string {
    return `${this.minutes % this.minutesInHour}`.padStart(2, "0");
  }

  private adjustNegative(mins: number): number {
    return mins < 0
      ? (this.minutes % this.minutesInDay) + this.minutesInDay
      : mins;
  }
}
