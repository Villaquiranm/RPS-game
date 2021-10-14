# RPS_GAME

## Architecture

This application is composed of two services, the front-ent component developed on vueJS and the backend service running on Golang.

## Running on host machine

#### Requirements

- NodeJS with npm installed
- Golang (1.15)

#### Run commands

All these commands should be run from the root folder:

- Run front-end: `npm run front`
- Run backend: `npm run api`
- Run unit tests: `npm run test`

  ```
  ?       RPS_gaming      [no test files]
  ok      RPS_gaming/game 0.569s  coverage: 17.1% of statements
  ?       RPS_gaming/game/constants       [no test files]
  ?       RPS_gaming/game/logic   [no test files]
  ?       RPS_gaming/model        [no test files]
  ok      RPS_gaming/moves        1.168s  coverage: 80.0% of statements
  ```

## Running on Docker

Comming soon...

#### Requirements

- Docker & docker-compose installations

#### Run commands

Go to the docker/ folder and execute `docker-compose up` or `docker-compose up` in case your user is not a member of `docker` group.

## Remaining work (TO-DO)

- [ ] Improve coverage tests
- [ ] Change the way we compute who wins and losses a game [useful link]()
- [ ] Add one extension to the game
