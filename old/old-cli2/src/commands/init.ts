import { basename } from 'path';
import { writeFile } from 'fs';
import readlineSync = require('readline-sync');

export const command = 'init';

export const describe = 'Initialize a package';

export const builder = {};

export const handler = function() {
  const pkg: any = {};

  function q(prompt: string, d?: string, optional: boolean = true) {
    let p: string;
    if (d) {
      p = prompt + ': (' + d + '): ';
    } else {
      p = prompt + ': ';
    }
    const res = readlineSync.question(p, {
      defaultInput: d,
    });
    if (optional) {
      if (!res) {
        return undefined;
      }
    }
    return res;
  }

  const helpText = `This utility will walk you through creating a wiz.json file.
It only covers the most common items, and tries to guess sensible defaults.

See \`wiz help json\` for definitive documentation on these fields
and exactly what they do.

Use \`wiz install <pkg>\` afterwards to install a package and
save it as a dependency in the wiz.json file.

Press ^C at any time to quit.
    `;

  console.log(helpText);
  pkg.name = q('package name', basename(process.cwd()));
  pkg.type = q('package type', 'program');
  pkg.version = q('version', '1.0.0');
  pkg.description = q('description', undefined, false);

  pkg.dependencies = {};
  pkg.devDependencies = {};

  pkg.repository = q('git repository');
  pkg.keywords = q('keywords', 'ai, ml, wiz');
  pkg.author = q('author', undefined, false);
  pkg.license = q('license', 'MIT');
  const output = JSON.stringify(pkg, null, '  ') + '\n';
  const outputFile = process.cwd() + '/wiz.json';
  console.log('About to write to ' + outputFile);
  console.log();
  console.log(output);
  q('Is this OK?', 'yes');
  writeFile(outputFile, output, (err: any) => {
    if (err) throw err;
  });
};
