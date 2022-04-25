export class DnDCharacter {
  strength: number;
  dexterity: number;
  constitution: number;
  intelligence: number;
  wisdom: number;
  charisma: number;
  hitpoints: number;

  constructor() {
    this.strength = DnDCharacter.generateAbilityScore();
    this.dexterity = DnDCharacter.generateAbilityScore();
    this.constitution = DnDCharacter.generateAbilityScore();
    this.intelligence = DnDCharacter.generateAbilityScore();
    this.wisdom = DnDCharacter.generateAbilityScore();
    this.charisma = DnDCharacter.generateAbilityScore();
    this.hitpoints = 10 + DnDCharacter.getModifierFor(this.constitution);
  }

  public static generateAbilityScore(): number {
    let scores: Array<number> = [
      DnDCharacter.throwDice(),
      DnDCharacter.throwDice(),
      DnDCharacter.throwDice(),
      DnDCharacter.throwDice(),
    ];

    const min = Math.min(...scores);

    return scores
      .filter((e) => e !== min)
      .reduce((total, current) => total + current, 0);
  }

  public static getModifierFor(abilityValue: number): number {
    return Math.floor((abilityValue - 10) / 2);
  }

  private static throwDice(): number {
    return Math.floor(Math.random() * 6) + 1;
  }
}
