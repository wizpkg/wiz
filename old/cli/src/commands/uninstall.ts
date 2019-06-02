export const command = ['uninstall [packages..]', 'remove', 'rm'];
export const describe = 'Uninstall packages';
export const builder = {
  packages: {
    type: 'string',
    describe: 'a list of package to uninstall',
  },
};
export const handler = (args: any) => {
  console.log('Uninstalling packages:', args.packages);
};
