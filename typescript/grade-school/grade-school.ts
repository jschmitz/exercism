type Dictionary = { [index: number]: string[] };

export class GradeSchool {
  private school: Dictionary = {};
  roster(): object {
    const clone = JSON.parse(JSON.stringify(this.school));
    return clone;
  }

  add(name: string, level: number): void {
    if (this.school[level] === undefined) {
      this.school[level] = [];
    }

    for (let l in this.school) {
      this.school[l] = this.school[l].filter((existingName) => {
        return existingName !== name;
      });
    }

    this.school[level].push(name);
    this.school[level].sort();
  }

  grade(level: number): Array<string> {
    if (this.school[level] === undefined) {
      return [];
    }

    return JSON.parse(JSON.stringify(this.school[level].sort()));
  }
}
