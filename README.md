# Go Orchestra Agent

> This thing can help you with computing your expressions with multi threading.
> After deploying Gorchestra deploy GoOrchestraAgents(https://github.com/J3olchara/GoOrchestraAgent)

# Installation guide
- Clone this repository
```bash
git clone https://github.com/J3olchara/GoOrchestra && cd GoOrchestra
```


# Linux
- Firstly you need installed docker-compose with minimum version 2.24
> to check version use ```docker-compose --version```
```bash
sudo apt-get install docker-compose
```
- Then create environment directory
```bash
mkdir ./env
```
- Then you need to fill your application environment in ./env directory as written in env-example directory or make defaults by
```bash
cp ./env_example/dev.app.env-example ./env/dev.app.env && cp ./env_example/dev.db.env-example ./env/dev.db.env
```
- Start Application
```bash
docker-compose -f docker-compose.dev.yml up
```


# MacOS
- Firstly you need installed docker-compose with minimum version 2.24
> to check version use ```docker-compose --version```
```bash
brew install docker-compose
```
- Then create environment directory
```bash
mkdir ./env
```
- Then you need to fill your application environment in ./env directory as written in env-example directory or make defaults by
```bash
cp ./env_example/dev.app.env-example ./env/dev.app.env && cp ./env_example/dev.db.env-example ./env/dev.db.env
```
- Start Application
```bash
docker-compose -f docker-compose.dev.yml up
```

# Windows
- Firstly you need installed docker-compose with minimum version 2.24
> to check version use ```docker-compose --version```
```bash
winget install docker-compose
```
- Then create environment directory
```bash
mkdir ./env
```
- Then you need to fill your application environment in ./env directory as written in env-example directory or make defaults by
```bash
cp ./env_example/dev.app.env-example ./env/dev.app.env && cp ./env_example/dev.db.env-example ./env/dev.db.env
```
- Start Application
```bash
docker-compose -f docker-compose.dev.yml up
```

# Have fun