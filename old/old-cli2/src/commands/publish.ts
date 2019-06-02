export const command = ['publish'];
export const describe = 'Publish a package';
export const builder = {};
export const handler = (args: any) => {
  console.log('Publishing packages:', args.packages);
};
