import { join } from 'path';
import { publish } from '../core/publish2';

export const command = ['publish'];
export const describe = 'Publish a package';
export const builder = {};

export const handler = (args: any) => {
  console.log(args.packages);
  if (!args.packages) {
    console.log('Publishing current package');
    args.packages = [null];
  } else {
    console.log('Publishing packages:', args.packages);
  }
  for (let pkg of args.packages) {
    let cwd = './';
    if (pkg !== null) {
      cwd = join(process.cwd(), pkg);
    }
    console.log(cwd, pkg);
    publish(cwd, pkg);
  }
};
