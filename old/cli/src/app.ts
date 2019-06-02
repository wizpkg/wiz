#!/usr/bin/env node
const yargs = require('yargs');

import { readFileSync } from 'fs';
import { resolve } from 'path';
// require('./commands');

const data = readFileSync(resolve(__dirname, '../package.json'), 'utf-8');
const version = JSON.parse(data).version;

yargs
  .version(version)
  .commandDir('./commands')
  .command('version', 'Show verison number', {}, () => {
    console.log(version);
  })
  .recommendCommands() // Recommend commands with similar spelling
  .demandCommand(1, '') // Require a command to be run
  .showHelpOnFail(true) // When it fails because no command is run (just `wiz`), thenshow help
  .parse();
