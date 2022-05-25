export class Robot {
  private iName: string = "";
  constructor() {}

  public get name(): string {
    if (this.iName === "") {
      const s1 = NamesSingleton.getInstance();
      this.iName = s1.newName();
    }
    return this.iName;
  }

  public resetName(): void {
    const s1 = NamesSingleton.getInstance();
    this.iName = s1.newName();
  }

  public static releaseNames(): void {
    const s1 = NamesSingleton.getInstance();
    s1.deleteNames();
  }
}

class NamesSingleton {
  private static instance: NamesSingleton;
  private names: Set<string> = new Set();

  private constructor() {}

  public static getInstance(): NamesSingleton {
    if (!NamesSingleton.instance) {
      NamesSingleton.instance = new NamesSingleton();
    }

    return NamesSingleton.instance;
  }

  public deleteNames() {
    this.names.clear();
  }

  public newName() {
    let newName = this.generateName();
    while (this.names.has(newName)) {
      newName = this.generateName();
    }

    this.names.add(newName);
    return newName;
  }

  private generateName(): string {
    return `${this.randomLetter()}${this.randomLetter()}${this.randomThreeDigitNumber()}`;
  }

  private randomThreeDigitNumber(): string {
    return Math.floor(Math.random() * 1000)
      .toString()
      .padStart(3, "0");
  }

  private randomLetter(): string {
    const upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    return upperChars[Math.floor(Math.random() * upperChars.length)];
  }
}
