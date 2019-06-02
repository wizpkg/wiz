#!/usr/bin/env node
const readlineSync = require('readline-sync');
const path = require('path');
const fs = require('fs');
const yargs = require('yargs')

yargs
  .version('0.0.1')
  .command('init', 'Initialize a package', {}, () => {
    let pkg: any = {};

    function q(prompt: string, d?: string, optional: boolean = true) {
      if(d) {
        prompt = prompt + ': (' + d + '): '
      } else {
        prompt = prompt + ': '
      }
      let res = readlineSync.question(prompt, {
        defaultInput: d
      });
      if (optional) {
        if (!res) {
          return undefined;
        }
      }
      return res;
    }

    var helpText = `This utility will walk you through creating a wiz.json file.
It only covers the most common items, and tries to guess sensible defaults.

See \`wiz help json\` for definitive documentation on these fields
and exactly what they do.

Use \`wiz install <pkg>\` afterwards to install a package and
save it as a dependency in the wiz.json file.

Press ^C at any time to quit.
    `
    
    console.log(helpText);
    pkg.name = q('package name', path.basename(process.cwd()))
    pkg.type = q('package type', 'program')
    pkg.version = q('version', '1.0.0')
    pkg.description = q('description', null, false)

    pkg.dependencies = {};
    pkg.devDependencies = {};

    pkg.repository = q('git repository')
    pkg.keywords = q('keywords', 'ai, ml, wiz')
    pkg.author = q('author', null, false)
    pkg.license = q('license', 'MIT')
    let output = JSON.stringify(pkg, null, '  ') + '\n';
    let outputFile = process.cwd() + '/wiz.json';
    console.log('About to write to ' + outputFile)
    console.log()
    console.log(output)
    let ok = q('Is this OK?', 'yes')
    fs.writeFile(outputFile, output, (err) => {
      if (err) throw err;
    })
  })
  .command(['install [packages..]', 'i', 'add'], 'Install packages', (args) => {
    args
      .positional('packages', {
        type: 'string',
        describe: 'a list of package to install',
        // coerce: function(arg) {
        //   return {
        //     name: arg.
        //   }
        // }
      })
  }, (args) => {
    console.log('Installing packages:', args.packages)
  })
  .recommendCommands()  // Recommend commands with similar spelling
  .demandCommand(1, '') // Require a command to be run
  .showHelpOnFail(true) // When it fails because no command is run (just `wiz`), thenshow help
  .parse()