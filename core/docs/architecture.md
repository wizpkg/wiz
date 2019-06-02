# architecture

```
├── core
│   ├── config
│   │   └── config.go
│   ├── delete
│   ├── install
│   │   ├── install.go
│   │   └── install_test.go
│   └── store
│       └── store.go
```

Under the core directory, we have commands, which perform operations on the `store`, which keeps our packages

## Core concepts

#### Store

A place where packages are installed. Usually on-disk on one machine (can be shared across machines becuase of hash garuntees)

#### Registry

A place where packages are stored (usually in a compressed file `.tar.gz`)
Also implements the Wiz Registry API over HTTP so packages can be downloaded.

#### Registry Package Index

A file that describes all the available packages in a registry and where to find them.

#### Reference

A field in a manifest that describes where and how to get additional data (usually source code or datasets)

#### The Hash Garuntee

Garuntees that given inputs of a package (it's name, version, and resolved dependencies), Wiz should fetch the exact same code and data.

#### Runtime dependency resolution/dependency fluidity

Allows a package to specify that it's version does not have to be locked down, that the resolved version is just recommended, and that Wiz can fetch any version in the dependency range without problems. This can be useful for things like large data sets,where versions may be quite minor, but we don't want Wiz to fetch the entire thing.

## Problems and Solutions

1. Dependency Hell - We make resolve the dependencies of each package at build time, and then bake those dependencies into that package. That means each package only refers to one specific version of each package, ensuring repeatability
2. Fast downloads - We cache every package, and don't delete packages/clear the cache unless asked to do so.
   - Another idea is to use something like git diffs where we store one version and only fetch the difference in versions. This would obviously only work for text/source code, but we would have binary caching for binaries. The performance impact of this would likely be negligible because it would only affect the first download of that version, as all others would be cached.
3. Easy updating of packages. You can just run wiz update, which will update your lock file to the highest available version in the acceptable range. You can also run `wiz update --full` or `-f` to modify the manifest with a range of the newest available versions
4. CI or your machine can't install packages when offline
   1. We keep a cache of the registry package index on your machine so we always know what packages there are and how many there are. (for Ubuntu 20k pkgs = 15mb)
   2. As described earlier, we cache packages, so if it's already installed, we're good
   3. Idea- use ML to determine packages to pre-install
