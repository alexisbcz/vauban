# Vauban

[AGPL-3.0 license](./LICENSE)

[![Join Discord](https://img.shields.io/badge/Join%20Discord-gray?style=flat&logo=discord&logoColor=white&link=https://discord.gg/eMUC7ejHja)](https://discord.gg/eMUC7ejHja)

> A simple web interface to manage your Docker installation.

### Development environment

```bash
git clone https://github.com/alexisbcz/vauban.git # clone the Git repository
cd vauban # change directory to Git repository

make setup # install deps

make db-up # spin up postgres container

make migration-up # migrate the database

make seed # seed the database
```

### Production environment

_WIP_
