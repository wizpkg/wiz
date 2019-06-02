import { basename, resolve } from 'path';
import { manifestFileName, packageTypes } from './config';
import mkdirp, { Made } from 'mkdirp';
import { readFileSync, rename, rmdir } from 'fs';

import { c as compress } from 'tar';
import { tmpdir } from 'os';

import crypto = require('crypto');

// TODO make async
export const publish = (cwd: string, pkgName: string) => {
  process.chdir(cwd);

  const data = readFileSync(resolve(cwd, manifestFileName), 'utf-8');
  const pkgMetadata = JSON.parse(data);
  console.log('Publishing package: ', pkgMetadata.name);
  console.log('Meta: ', pkgMetadata);

  if (pkgMetadata.type === packageTypes.PROGRAM) {
    console.log('prog');
    const id = crypto.randomBytes(16).toString('hex');
    const fileName = `${pkgMetadata.name}-${pkgMetadata.version}.tgz`;
    const fileLocation = `${tmpdir()}/wiz/publish/${id}/`;
    // console.log(fileLocation);
    // let dir = basename(fileLocation);
    // console.log(dir);
    console.log(fileLocation);
    mkdirp(fileLocation, (err: NodeJS.ErrnoException, m: Made) => {
      console.log('hi');
      console.log(m, err);
      compress(
        {
          gzip: true,
          file: `${fileLocation}${fileName}`,
        },
        [cwd]
      ).then(_ => {
        console.log('tarball created');
        rename(`${fileLocation}${fileName}`, resolve(cwd, fileName), err => {
          console.log(err);
          rmdir(fileLocation, err => {
            console.log(err);
          });
        });
      });
    });
  }
};
