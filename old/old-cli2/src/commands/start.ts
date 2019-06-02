export const command = ['start [packages..]'];
export const describe = 'Run start script for current package or list of packages';
export const builder = {
  packages: {
    type: 'string',
    describe: 'a list of package to start',
  },
};
export const handler = (args: any) => {
  console.log('Starting packages:', args.packages);
};
