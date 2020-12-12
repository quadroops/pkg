# pkg

This repo used as a collection of golang libraries.  If you want to install a library from this collection, follow this format:

```
go get -v -u github.com/quadroops/pkg/<domain>
```

## Versioning

Please use commit hash id as version.  The reason behind this step is, `quadroops/pkg` is a _monorepo_ of library collections.  Each folder from `quadroops/pkg` represent a single pkg, the cons is we cannot tagging a version based on spesific folder.

An example using commit hash:

```
go get -v -u github.com/quadroops/pkg/log@c849d91569ada73bfb115c8c62d114025535c605
```

You can check our commit histories based on spesific library, like this -> [quadroops/pkg/log](https://github.com/quadroops/pkg/commits/main/log)