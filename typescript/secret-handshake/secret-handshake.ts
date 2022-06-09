export function commands(n: number): Array<string> {
  let cmds = [];

  if (n & 1) cmds.push("wink");
  if (n & 2) cmds.push("double blink");
  if (n & 4) cmds.push("close your eyes");
  if (n & 8) cmds.push("jump");
  if (n & 16) cmds.reverse();

  return cmds;
}
