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

[//]: <> (#### Requirements)
[comment]: <> (- Docker & docker-compose installations)
[comment]: <> (#### Run commands)
[comment]: <> (go to the docker/ folder and execute:)
[comment]: <> (`docker-compose up`)


## Remaining work (TO-DO)

- [ ] Improve coverage tests
- [ ] Change the way we compute who wins and losses a game [useful link](https://upload.wikimedia.org/wikipedia/commons/thumb/f/fe/Rock_Paper_Scissors_Lizard_Spock_en.svg/400px-Rock_Paper_Scissors_Lizard_Spock_en.svg.png)
- [ ] Add one extension to the game