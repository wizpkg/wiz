import { client } from './grpc';

const DEFAULT_REGISTRY = '/wiz/registry';

export const publish = (cwd: string, pkgName: string) => {
  client.publish(pkgName, DEFAULT_REGISTRY);
};
