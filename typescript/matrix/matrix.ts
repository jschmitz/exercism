export class Matrix {
  matrix: number[][] = [[]];

  constructor(m: string) {
    const rows = m.split("\n");

    rows.forEach((row: string, index) => {
      const cells = row.split(" ");
      this.matrix[index] = cells.map((cell) => Number(cell));
    });
  }

  get rows(): number[][] {
    return this.matrix;
  }

  get columns(): number[][] {
    return this.matrix[0].map((_, colIndex) =>
      this.matrix.map((row) => row[colIndex])
    );
  }
}
