export const command = ['install [packages..]', 'i', 'add'];
export const describe = 'Install packages';
export const builder = {
  packages: {
    type: 'string',
    describe: 'a list of package to install',
  },
};
export const handler = (args: any) => {
  console.log('Installing packages:', args.packages);
};
