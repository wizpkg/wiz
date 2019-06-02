interface Command {
  // The title of the command
  command: string | string[];
  // A description of the command
  describe: string;
  // An object containing the options that the command accepts
  builder: object;
  // A function that will be passed the arguments
  handler: (argv: object) => void;
}

export function validateCommand(cmd: Command): Command {
  return cmd;
}
